# modules/network/outputs.tf

# 作成した VPC の ID
output "vpc_id" {
  description = "The ID of the VPC"
  value       = aws_vpc.main.id
}

# 作成したパブリックサブネットの ID のリスト
output "public_subnet_ids" {
  description = "List of public subnet IDs"
  value       = aws_subnet.public.*.id # *.id は、複数のリソースのIDリストを取得する記法
}

# 作成したプライベートサブネットの ID のリスト
output "private_subnet_ids" {
  description = "List of private subnet IDs"
  value       = aws_subnet.private.*.id
}

# (オプション) 作成したパブリックサブネットの CIDR ブロックのリスト
# output "public_subnet_cidrs" {
#   description = "List of public subnet CIDR blocks"
#   value       = aws_subnet.public.*.cidr_block
# }

# (オプション) 作成したプライベートサブネットの CIDR ブロックのリスト
# output "private_subnet_cidrs" {
#   description = "List of private subnet CIDR blocks"
#   value       = aws_subnet.private.*.cidr_block
# }

# 作成したパブリックルートテーブルの ID
output "public_route_table_id" {
  description = "The ID of the public route table"
  value       = aws_route_table.public.id
}

# 作成したプライベート・ルートテーブルの ID のリスト
output "private_route_table_ids" {
  description = "List of private route table IDs (one per AZ)"
  value       = aws_route_table.private.*.id
}

