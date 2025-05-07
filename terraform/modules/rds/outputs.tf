output "master_user_secret_arn" {
  value = try(aws_db_instance.main.master_user_secret[0].secret_arn, "") # The actual value to be outputted
}
