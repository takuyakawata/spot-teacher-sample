# modules/network/main.tf

# VPC の作成
resource "aws_vpc" "main" {
  cidr_block           = var.vpc_cidr
  enable_dns_hostnames = true # DNSホスト名を有効化
  enable_dns_support   = true # DNSサポートを有効化

  tags = {
    Name = "${var.name}-${var.env}-vpc"
    Env  = var.env
  }
}

# インターネットゲートウェイ (IGW) の作成
resource "aws_internet_gateway" "main" {
  vpc_id = aws_vpc.main.id

  tags = {
    Name = "${var.name}-${var.env}-igw"
    Env  = var.env
  }
}

# パブリックサブネットの作成 (各AZに1つずつ)
resource "aws_subnet" "public" {
  count = length(var.public_subnet_cidrs) # public_subnet_cidrs の数だけ作成

  vpc_id            = aws_vpc.main.id
  cidr_block        = var.public_subnet_cidrs[count.index]
  availability_zone = var.azs[count.index] # 対応するAZに関連付け
  # パブリックサブネットでは、ここに配置されるインスタンスに自動でパブリックIPを割り当てる設定をすることが多い
  # ALBなどは自動でパブリックIPが割り当てられるため必須ではないですが、NAT GWなどには必要
  map_public_ip_on_launch = true

  tags = {
    Name = "${var.name}-${var.env}-public-subnet-${var.azs[count.index]}"
    Env  = var.env
  }
}

# プライベートサブネットの作成 (各AZに1つずつ)
resource "aws_subnet" "private" {
  count = length(var.private_subnet_cidrs) # private_subnet_cidrs の数だけ作成

  vpc_id            = aws_vpc.main.id
  cidr_block        = var.private_subnet_cidrs[count.index]
  availability_zone = var.azs[count.index] # 対応するAZに関連付け
  map_public_ip_on_launch = false # プライベートサブネットでは自動パブリックIP割り当てはしない

  tags = {
    Name = "${var.name}-${var.env}-private-subnet-${var.azs[count.index]}"
    Env  = var.env
  }
}

# EIP の作成 (NAT Gateway 用に各AZで必要)
resource "aws_eip" "nat_gateway" {
  count = length(var.azs) # AZの数だけ作成
  # VPC内のEIPは、インスタンスではなくNAT GWなどに関連付けられるため VPC = true が必要
  domain = "vpc"

  tags = {
    Name = "${var.name}-${var.env}-nat-eip-${var.azs[count.index]}"
    Env  = var.env
  }
}

# NAT Gateway の作成 (各AZに1つずつ、対応するパブリックサブネットに配置)
resource "aws_nat_gateway" "main" {
  count = length(var.azs) # AZの数だけ作成

  allocation_id = aws_eip.nat_gateway[count.index].id
  # NAT GWは、対応するAZのパブリックサブネットに配置する
  subnet_id     = aws_subnet.public[count.index].id

  tags = {
    Name = "${var.name}-${var.env}-nat-gw-${var.azs[count.index]}"
    Env  = var.env
  }
  # IGWが作成されてからNAT GWを作成する必要があるため依存関係を指定
  depends_on = [aws_internet_gateway.main]
}

# ルートテーブルの作成 (パブリックサブネット用)
resource "aws_route_table" "public" {
  vpc_id = aws_vpc.main.id

  # インターネット (0.0.0.0/0) へのルートは IGW を向ける
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.main.id
  }

  tags = {
    Name = "${var.name}-${var.env}-public-rt"
    Env  = var.env
  }
}

# ルートテーブルの関連付け (パブリックサブネットとパブリックRTを関連付け)
resource "aws_route_table_association" "public" {
  count = length(var.public_subnet_cidrs) # パブリックサブネットの数だけ関連付け

  subnet_id      = aws_subnet.public[count.index].id
  route_table_id = aws_route_table.public.id
}

# ルートテーブルの作成 (プライベートサブネット用)
# 各AZのプライベートサブネットは、それぞれのAZにあるNAT GW を向けるルートを持つ必要があるため、AZの数だけ作成
resource "aws_route_table" "private" {
  count = length(var.azs) # AZの数だけ作成

  vpc_id = aws_vpc.main.id

  # インターネット (0.0.0.0/0) へのルートは、対応するAZのNAT GW を向ける
  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.main[count.index].id
  }

  tags = {
    Name = "${var.name}-${var.env}-private-rt-${var.azs[count.index]}"
    Env  = var.env
  }
}

# ルートテーブルの関連付け (プライベートサブネットとプライベートRTを関連付け)
# 各プライベートサブネットを、対応するAZのプライベートRTに関連付ける
resource "aws_route_table_association" "private" {
  count = length(var.private_subnet_cidrs) # プライベートサブネットの数だけ関連付け

  subnet_id      = aws_subnet.private[count.index].id
  route_table_id = aws_route_table.private[count.index].id
}
