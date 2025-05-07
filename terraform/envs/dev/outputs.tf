# ==============================================================================
# 環境レベルでの出力値 (デプロイ後に確認したい情報など)
# 例: ALB の DNS 名, RDS のエンドポイントなど
# ==============================================================================

# ALB の DNS 名を出力 (ALB モジュールの出力を参照)
output "alb_dns_name" {
  description = "The DNS name of the application load balancer"
  value       = module.alb.alb_dns_name
}

# (オプション) RDS の DB 名とユーザー名
# output "db_name" {
#   description = "The name of the database"
#   value       = var.db_name # envs/dev の変数を直接参照
# }
# output "db_username" {
#   description = "The username for the database"
#   value       = var.db_username # envs/dev の変数を直接参照
# }

# (オプション) Secrets Manager の Secret ID (DBパスワード用)
# output "db_password_secret_id" {
#   description = "The Secrets Manager Secret ID for the database password"
#   value       = "your/db/password/secret/id" # 例: 直書きまたは envs/dev の変数から取得
# }

# (オプション) Bastion Host の Public IP など (Bastion を作成する場合)
# output "bastion_public_ip" {
#   description = "The public IP address of the Bastion host"
#   value       = module.bastion.public_ip # Bastion モジュールを作成した場合の出力を参照
# }
