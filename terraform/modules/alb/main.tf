# modules/alb/main.tf

# ==============================================================================
# Application Load Balancer (ALB) の作成
# ==============================================================================
resource "aws_lb" "main" {
  name               = "${var.name}-${var.env}-alb"
  internal           = false # インターネット向けALB (内部向けの場合は true)
  load_balancer_type = "application" # ALBとして作成
  # ALBは複数のパブリックサブネットにまたがるように配置する
  subnets            = var.public_subnet_ids
  security_groups    = var.security_group_ids # ALB用セキュリティグループを適用

  tags = {
    Name = "${var.name}-${var.env}-alb"
  }
}

# ==============================================================================
# ターゲットグループの作成
# ECSタスクが登録されるグループ
# ==============================================================================
resource "aws_lb_target_group" "main" {
  name        = "${var.name}-${var.env}-tg"
  port        = var.app_container_port # ターゲット（ECSタスク）がListenするポート
  protocol    = "HTTP" # ALBとターゲット間のプロトコル (多くの場合TCP)
  vpc_id      = var.vpc_id
  target_type = "ip" # Fargate の場合、ターゲットタイプは 'ip'

  # ヘルスチェックの設定
  health_check {
    enabled           = true
    interval          = 30 # チェック間隔 (秒)
    path              = "/" # ヘルスチェックのパス (Goアプリに実装が必要)
    protocol          = "HTTP" # ヘルスチェックに使用するプロトコル (HTTP/HTTPS/TCPなど)
    timeout           = 5  # タイムアウト (秒)
    healthy_threshold = 2  # ヘルシーと判断するための連続成功回数
    matcher           = "200" # 正常と判断するHTTPステータスコード
    port              = "traffic-port" # ターゲットグループのポート、または特定のポート番号
  }

  tags = {
    Name = "${var.name}-${var.env}-tg"
  }
}

# ==============================================================================
# リスナーの作成
# ALB が外部からの接続を待ち受ける設定
# ==============================================================================
resource "aws_lb_listener" "main" {
  load_balancer_arn = aws_lb.main.arn # 作成したALBのARN
  port              = var.alb_listener_port # 待ち受けポート
  protocol          = var.alb_listener_protocol # 待ち受けプロトコル (HTTP or HTTPS)
  # certificate_arn   = var.ssl_certificate_arn # HTTPS の場合、ACM 証明書を指定

  # デフォルトアクション: 受信したトラフィックを上記のターゲットグループに転送
  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.main.arn
  }

  # dynamic "certificate" {
  #   for_each = var.ssl_certificate_arn != null ? [var.ssl_certificate_arn] : []
  #   content {
  #     certificate_arn = certificate.value
  #   }
  # }

  tags = {
    Name = "${var.name}-${var.env}-listener-${var.alb_listener_port}"
  }
}
