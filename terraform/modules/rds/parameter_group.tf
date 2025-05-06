resource "aws_db_parameter_group" "main" {
  name        = "${var.name}-${var.env}-db-pram"
  description = "DB parameter group for ${var.name}-${var.env}-db"
  family      = "mysql8.0"
  tags = {
    Name = "${var.name}-${var.env}-db-pram"
  }
  parameter {
    apply_method = "immediate"
    name         = "character_set_client"
    value        = "utf8mb4"
  }
  parameter {
    apply_method = "immediate"
    name         = "character_set_connection"
    value        = "utf8mb4"
  }
  parameter {
    apply_method = "immediate"
    name         = "character_set_database"
    value        = "utf8mb4"
  }
  parameter {
    apply_method = "immediate"
    name         = "character_set_filesystem"
    value        = "utf8mb4"
  }
  parameter {
    apply_method = "immediate"
    name         = "character_set_results"
    value        = "utf8mb4"
  }
  parameter {
    apply_method = "immediate"
    name         = "character_set_server"
    value        = "utf8mb4"
  }
  parameter {
    apply_method = "immediate"
    name         = "collation_connection"
    value        = "utf8mb4_general_ci"
  }
  parameter {
    apply_method = "immediate"
    name         = "collation_server"
    value        = "utf8mb4_general_ci"
  }
  parameter {
    apply_method = "immediate"
    name         = "slow_query_log"
    value        = "0"
  }
  parameter {
    apply_method = "immediate"
    name         = "time_zone"
    value        = "Asia/Tokyo"
  }
  parameter {
    apply_method = "pending-reboot"
    name         = "performance_schema"
    value        = "1"
  }
  parameter {
    apply_method = "pending-reboot"
    name         = "skip-character-set-client-handshake"
    value        = "1"
  }
  parameter {
    apply_method = "pending-reboot"
    name         = "skip_name_resolve"
    value        = "1"
  }
}
