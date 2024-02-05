# This is still a test. Not a real thing in fact. Just wanted different ways to test the go app
resource "null_resource" "tf_cli_null" {
  count = var.resource_name == "null" ? var.resource_count : 0
  provisioner "local-exec" {
    command = "echo '###### This is ${count.index}-${var.resource_name} file' > ./${count.index}-${var.resource_name}.txt"
  }
}

resource "null_resource" "tf_cli_rds" {
  count = var.resource_name == "rds" ? var.resource_count : 0
  provisioner "local-exec" {
    command = "echo '###### This is ${count.index}-${var.resource_name} file' > ./${count.index}-${var.resource_name}.txt"
  }
}

resource "null_resource" "tf_cli_ec2" {
  count = var.resource_name == "ecr" ? var.resource_count : 0
  provisioner "local-exec" {
    command = "echo '###### This is ${count.index}-${var.resource_name} file' > ./${count.index}-${var.resource_name}.txt"
  }
}