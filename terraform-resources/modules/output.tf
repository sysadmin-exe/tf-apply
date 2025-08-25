output "null_resource" {
  value = null_resource.tf_cli_null.*.id
}

output "rds_resource" {
  value = null_resource.tf_cli_rds.*.id
}

output "ec2_resource" {
  value = null_resource.tf_cli_ec2.*.id
}