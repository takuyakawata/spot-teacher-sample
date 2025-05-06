# modules/security_groups/outputs.tf

# ALB 用セキュリティグループの ID
output "alb_security_group_id" {
  description = "The ID of the ALB security group"
  value       = aws_security_group.alb.id
}

# ECS タスク用セキュリティグループの ID
output "ecs_security_group_id" {
  description = "The ID of the ECS task security group"
  value       = aws_security_group.ecs.id
}

# RDS データベース用セキュリティグループの ID
output "rds_security_group_id" {
  description = "The ID of the RDS database security group"
  value       = aws_security_group.rds.id
}

# Bastion Host 用セキュリティグループの ID (Bastion を作成しない場合は null になる)
output "bastion_security_group_id" {
  description = "The ID of the Bastion host security group"
  # count が 0 の場合はリソースが作成されないため、リストの最初の要素を参照する際に
  # ? を使うことで、存在しない場合は null を返すようにします。
  value = length(aws_security_group.bastion) > 0 ? aws_security_group.bastion[0].id : null
}
