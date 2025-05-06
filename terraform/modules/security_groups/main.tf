# modules/security_groups/main.tf

# ==============================================================================
# ALB (Application Load Balancer) 用セキュリティグループ
# インターネットからの HTTP/HTTPS アクセスを許可
# ==============================================================================
resource "aws_security_group" "alb" {
  name        = "${var.name}-${var.env}-alb-sg"
  description = "Allow HTTP/HTTPS access to the ALB"
  vpc_id      = var.vpc_id # network モジュールから受け取った VPC ID を使用

  tags = {
    Name = "${var.name}-${var.env}-alb-sg"
    Env  = var.env
  }
}

# ALB SG - インバウンド HTTP (80番ポート)
resource "aws_security_group_rule" "alb_ingress_http" {
  type              = "ingress" # インバウンドルール
  from_port         = 80
  to_port           = 80
  protocol          = "tcp"
  cidr_blocks       = ["0.0.0.0/0"] # 全世界のIPv4アドレスからのアクセスを許可
  # ipv6_cidr_blocks = ["::/0"] # IPv6も許可する場合
  security_group_id = aws_security_group.alb.id # ALB SG自体にルールを適用
}

# ALB SG - インバウンド HTTPS (443番ポート) - HTTPSを使う場合に必要
# resource "aws_security_group_rule" "alb_ingress_https" {
#   type              = "ingress"
#   from_port         = 443
#   to_port           = 443
#   protocol          = "tcp"
#   cidr_blocks       = ["0.0.0.0/0"]
#   # ipv6_cidr_blocks = ["::/0"]
#   security_group_id = aws_security_group.alb.id
# }

# ALB SG - アウトバウンド (任意 - 必要に応じて制限)
# デフォルトではすべてのアウトバウンドが許可されていますが、必要に応じて明示的に許可ルールを追加・制限します。


# ==============================================================================
# ECS タスク (Go アプリ) 用セキュリティグループ
# ALB からのアクセス、RDS へのアウトバウンドなどを許可
# ==============================================================================
resource "aws_security_group" "ecs" {
  name        = "${var.name}-${var.env}-ecs-sg"
  description = "Allow ALB access to ECS tasks and outbound to RDS/Internet"
  vpc_id      = var.vpc_id

  tags = {
    Name = "${var.name}-${var.env}-ecs-sg"
    Env  = var.env
  }
}

# ECS SG - インバウンド (ALB からのアクセス許可)
# 許可元として ALB SG の ID を指定するのがベストプラクティスです。
# ALBからのトラフィックは、ALB SGを経由してこのECS SGに到達します。
resource "aws_security_group_rule" "ecs_ingress_from_alb" {
  type              = "ingress"
  from_port         = var.app_container_port # Go アプリがListenするポート
  to_port           = var.app_container_port
  protocol          = "tcp"
  # 許可元として ALB 用セキュリティグループの ID を指定！
  source_security_group_id = aws_security_group.alb.id
  security_group_id        = aws_security_group.ecs.id # ECS SG自体にルールを適用
}

# ECS SG - アウトバウンド (RDS へのアクセス許可)
# 許可先として RDS SG の ID を指定するのがベストプラクティスです。
resource "aws_security_group_rule" "ecs_egress_to_rds" {
  type              = "egress" # アウトバウンドルール
  from_port         = var.rds_port
  to_port           = var.rds_port
  protocol          = "tcp"
  # 許可先として RDS 用セキュリティグループの ID を指定！ (まだ定義してないですが、後で定義します)
  # ここではプレースホルダーとして aws_security_group.rds.id を使います。
  destination_security_group_id = aws_security_group.rds.id
  security_group_id             = aws_security_group.ecs.id
}

# ECS SG - アウトバウンド (インターネットへのアクセス許可 - 必要に応じて)
# NAT Gateway経由でインターネットに出るために、アウトバウンドを許可します。
# デフォルトでは全アウトバウンド許可されていることが多いですが、明示的に許可します。
resource "aws_security_group_rule" "ecs_egress_to_internet" {
  type              = "egress"
  from_port         = 0 # すべてのポート
  to_port           = 65535
  protocol          = "all" # TCP/UDPなどすべて
  cidr_blocks       = ["0.0.0.0/0"] # 全世界のIPv4アドレスへ
  security_group_id = aws_security_group.ecs.id
}


# ==============================================================================
# RDS (データベース) 用セキュリティグループ
# ECS タスクや Bastion ホストからのアクセスを許可
# ==============================================================================
resource "aws_security_group" "rds" {
  name        = "${var.name}-${var.env}-rds-sg"
  description = "Allow access to RDS database from ECS tasks/Bastion"
  vpc_id      = var.vpc_id

  tags = {
    Name = "${var.name}-${var.env}-rds-sg"
    Env  = var.env
  }
}

# RDS SG - インバウンド (ECS SG からのアクセス許可)
# 許可元として ECS SG の ID を指定するのがベストプラクティスです。
resource "aws_security_group_rule" "rds_ingress_from_ecs" {
  type              = "ingress"
  from_port         = var.rds_port # RDS がListenするポート
  to_port           = var.rds_port
  protocol          = "tcp"
  # 許可元として ECS 用セキュリティグループの ID を指定！
  source_security_group_id = aws_security_group.ecs.id
  security_group_id        = aws_security_group.rds.id # RDS SG自体にルールを適用
}

# RDS SG - インバウンド (Bastion SG からのアクセス許可 - 必要に応じて)
# Bastion ホスト経由でDBに接続する場合に必要です。
# 許可元として Bastion SG の ID を指定します。
# count を使うことで、bastion_ingress_cidrs が空の場合はこのルールは作成されません。
resource "aws_security_group_rule" "rds_ingress_from_bastion" {
  count = length(var.bastion_ingress_cidrs) > 0 ? 1 : 0 # bastion_ingress_cidrs が1つ以上ある場合のみルールを作成

  type              = "ingress"
  from_port         = var.rds_port
  to_port           = var.rds_port
  protocol          = "tcp"
  # 許可元として Bastion 用セキュリティグループの ID を指定！ (まだ定義してないですが、後で定義します)
  source_security_group_id = aws_security_group.bastion.id
  security_group_id        = aws_security_group.rds.id
}


# ==============================================================================
# Bastion Host 用セキュリティグループ (オプション)
# 特定の IP アドレスから SSH アクセスを許可
# ==============================================================================
resource "aws_security_group" "bastion" {
  # count を使うことで、bastion_ingress_cidrs が空の場合はこの SG 自体を作成しません。
  count = length(var.bastion_ingress_cidrs) > 0 ? 1 : 0

  name        = "${var.name}-${var.env}-bastion-sg"
  description = "Allow SSH access to Bastion host from specific IPs"
  vpc_id      = var.vpc_id

  tags = {
    Name = "${var.name}-${var.env}-bastion-sg"
    Env  = var.env
  }
}

# Bastion SG - インバウンド (指定された IP からの SSH 許可)
resource "aws_security_group_rule" "bastion_ingress_ssh" {
  count = length(var.bastion_ingress_cidrs) # 許可元 IP の数だけルールを作成

  type              = "ingress"
  from_port         = var.ssh_port
  to_port           = var.ssh_port
  protocol          = "tcp"
  # 許可元として bastion_ingress_cidrs 変数で指定された IP リストを使用
  cidr_blocks       = [var.bastion_ingress_cidrs[count.index]]
  security_group_id = aws_security_group.bastion[0].id # count を使っているため、[0] で SG を参照
}

# Bastion SG - アウトバウンド (任意 - 必要に応じて制限)
# Bastion からインターネットや社内ネットワークへのアクセスを許可する場合など。
# デフォルトでは全アウトバウンド許可されています。
