# modules/ecs/outputs.tf

# 作成したECS クラスター名
output "cluster_name" {
  description = "The name of the ECS cluster"
  value       = aws_ecs_cluster.main.name
}

# 作成した ECS サービス名
output "service_name" {
  description = "The name of the ECS service"
  value       = aws_ecs_service.app_service.name
}

# 作成した ECS クラスター ID
output "ecs_cluster_id" {
  description = "The ID of the ECS cluster"
  value       = aws_ecs_cluster.main.id
}

