# modules/alb/outputs.tf

# 作成した ALB の DNS 名 (外部からのアクセスに利用)
output "alb_dns_name" {
  description = "The DNS name of the ALB"
  value       = aws_lb.main.dns_name
}

# 作成した ALB の ARN
output "alb_arn" {
  description = "The ARN of the ALB"
  value       = aws_lb.main.arn
}

# 作成したターゲットグループの ARN (ECS サービスとの関連付けに利用)
output "alb_target_group_arn" {
  description = "The ARN of the ALB target group"
  value       = aws_lb_target_group.main.arn
}
