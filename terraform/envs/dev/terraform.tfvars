# envs/dev/terraform.tfvars
# envs/dev/variables.tf で定義した変数の値を記述します
# terraform apply 実行時に自動的に読み込まれます

# 環境固有の基本設定
region       = "ap-northeast-1" # 使用するAWSリージョン
app_name     = "spot-teacher-dev"  # アプリケーション名 (開発環境用)

# デプロイするDockerイメージ
image_url    = "xxxxxxxxxxxx.dkr.ecr.ap-northeast-1.amazonaws.com/my-go-app:latest" # 実際にプッシュしたイメージURLに置き換える

# ECSタスク設定
desired_count  = 1 # 実行タスク数
fargate_cpu    = 256
fargate_memory = 512
container_port = 8080 # アプリがListenするポート (Goアプリのコードに合わせる)

# ネットワーク設定 (Networkモジュールに渡す値)
# variables.tf でデフォルト値を指定しない場合はここで必須
vpc_cidr            = "10.0.0.0/16"
public_subnet_cidrs = ["10.0.1.0/24", "10.0.2.0/24"] # 使用するAZの数だけ定義
private_subnet_cidrs = ["10.0.10.0/24", "10.0.11.0/24"] # 使用するAZの数だけ定義
azs                 = ["ap-northeast-1a", "ap-northeast-1c"] # 使用するAZのリスト (public/private subnet の数と一致させる)

# RDS設定 (RDSモジュールに渡す値)
db_instance_type     = "db.t3.micro"
db_allocated_storage = 20
db_engine            = "postgres" # 使用するDBエンジン (例: postgres, mysql)
db_engine_version    = "15.5" # 使用するDBエンジンのバージョン (db_engine に合わせる)
db_name              = "spot-teacher-db"
db_username          = "user"
db_password_secret_id = "your/db/password/secret/id" # ★ Secrets Managerに事前に作成したSecret IDに置き換える ★

db_multi_az          = false # 開発環境なので Multi-AZ は false
db_deletion_protection = false # 開発環境なので削除保護は false

# Security Groups 設定 (Security Groupsモジュールに渡す値)
# variables.tf でデフォルト値を指定しない場合はここで必須
# rds_port = 5432 # RDS がListenするポート (DB エンジンに合わせて変更)
# ssh_port = 22 # Bastion 用 SSH ポート
# bastion_ingress_cidrs = [] # BastionへのSSHを許可するIPリスト (例: ["203.0.113.0/24"]) 空リストならBastion関連リソースは作成されない

# ALB 設定 (ALBモジュールに渡す値)
# variables.tf でデフォルト値を指定しない場合はここで必須
# alb_listener_port     = 80 # ALB が待ち受けるポート (HTTP:80, HTTPS:443 など)
# alb_listener_protocol = "HTTP" # ALB が待ち受けるプロトコル (HTTP or HTTPS)
# ssl_certificate_arn   = null # HTTPS の場合、ACM 証明書の ARN を記述 (HTTP なら null)

# IAM 関連変数 (IAMモジュールに渡す値)
# envs/dev/variables.tf で定義している場合、ここで値を設定
# github_access_token = "..." # アプリが必要なシークレットなら Secrets Manager を推奨
# vercel_api_token = "..."    # アプリが必要なシークレットなら Secrets Manager を推奨
# team_id = "..." # 用途が明確なら記述
