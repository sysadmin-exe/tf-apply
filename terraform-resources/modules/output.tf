output "null_resource" {
  value = { for k, v in null_resource.tf_cli_null : k => v.id }
}

output "rds_resource" {
  value = { for k, v in null_resource.tf_cli_rds : k => v.id }
}

output "ec2_resource" {
  value = { for k, v in null_resource.tf_cli_ec2 : k => v.id }
}