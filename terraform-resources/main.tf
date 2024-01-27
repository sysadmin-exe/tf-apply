resource "null_resource" "tf_cli" {
  count = var.resource_count
  provisioner "local-exec" {
    command = "echo '###### This is ${count.index}-${var.resource_name} file' > ./${count.index}-${var.resource_name}"
  }
}
