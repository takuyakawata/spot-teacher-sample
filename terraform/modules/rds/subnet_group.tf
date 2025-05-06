resource "aws_db_subnet_group" "main" {
  name        = "${var.name}-${var.env}-subnet-group"
  description = "DB subnet group for ${var.name}-${var.env}-db"
  subnet_ids  = var.subnet_ids
  tags = {
    Name = "${var.name}-${var.env}-subnet-group"
  }
}
