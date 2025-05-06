# # ECS 用セキュリティグループの作成
# resource "aws_security_group" "ecs" {
#   name        = "${var.name}-${var.env}-ecs-sg"
#   description = "Security Group for ECS"
#   vpc_id      = aws_vpc.main.id
#   tags = {
#     Name = "${var.name}-${var.env}-ecs-sg"
#   }
#   egress {
#     from_port   = 0
#     to_port     = 0
#     protocol    = "-1"
#     cidr_blocks = ["0.0.0.0/0"]
#   }
# }
#
# # Bastionのセキュリティグループ
# resource "aws_security_group" "bastion" {
#   name        = "${var.name}-${var.env}-bastion-sg"
#   description = "Security Group for Bastion"
#   vpc_id      = aws_vpc.main.id
#   tags = {
#     Name = "${var.name}-${var.env}-bastion-sg"
#   }
#   egress {
#     from_port   = 0
#     to_port     = 0
#     protocol    = "-1"
#     cidr_blocks = ["0.0.0.0/0"]
#   }
# }
# resource "aws_security_group_rule" "bastion_same_sg" {
#   security_group_id        = aws_security_group.bastion.id
#   from_port                = 0
#   protocol                 = "-1"
#   source_security_group_id = aws_security_group.bastion.id
#   to_port                  = 0
#   type                     = "ingress"
#   description              = "Allow same security group"
# }
# resource "aws_security_group_rule" "ssh" {
#   security_group_id = aws_security_group.bastion.id
#   from_port         = 22
#   protocol          = "tcp"
#   cidr_blocks       = ["0.0.0.0/0"]
#   to_port           = 22
#   type              = "ingress"
#   description       = "Allow SSH inbound traffic"
# }
#
# #rds
# resource "aws_security_group" "rds" {
#   name        = "${var.name}-${var.env}-rds-sg"
#   description = "Security Group for RDS"
#   vpc_id      = aws_vpc.main.id
#   tags = {
#     Name = "${var.name}-${var.env}-rds-sg"
#   }
#   egress {
#     from_port   = 0
#     to_port     = 0
#     protocol    = "-1"
#     cidr_blocks = ["0.0.0.0/0"]
#   }
# }
# resource "aws_security_group_rule" "rds_same_sg" {
#   security_group_id        = aws_security_group.rds.id
#   from_port                = 0
#   protocol                 = "-1"
#   source_security_group_id = aws_security_group.rds.id
#   to_port                  = 0
#   type                     = "ingress"
#   description              = "Allow same security group"
# }
# resource "aws_security_group_rule" "ecs_access" {
#   security_group_id        = aws_security_group.rds.id
#   from_port                = 3306
#   protocol                 = "tcp"
#   source_security_group_id = aws_security_group.ecs.id
#   to_port                  = 3306
#   type                     = "ingress"
#   description              = "Allow ECS to access RDS"
# }
#
# resource "aws_security_group_rule" "bastion_access" {
#   security_group_id        = aws_security_group.rds.id
#   from_port                = 3306
#   protocol                 = "tcp"
#   source_security_group_id = aws_security_group.bastion.id
#   to_port                  = 3306
#   type                     = "ingress"
#   description              = "Allow Bastion to access RDS"
# }
#
# resource "aws_security_group_rule" "ecs_private_access" {
#   security_group_id        = aws_security_group.rds.id
#   from_port                = 3306
#   protocol                 = "tcp"
#   source_security_group_id = aws_security_group.ecs_private.id
#   to_port                  = 3306
#   type                     = "ingress"
#   description              = "Allow ECS Private Subnet to access RDS"
# }
