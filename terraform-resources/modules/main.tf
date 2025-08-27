# This is still a test. Not a real thing in fact. Just wanted different ways to test the go app

locals{
  resources_map = {for r in var.resources_list : "${r.application_name}-${r.resource_count}-${r.resource_type}" => {
    application_name = r.application_name
    resource_type    = r.resource_type
    resource_count   = r.resource_count
  }}
}
resource "null_resource" "tf_cli_null" {
  for_each = { for k, v in local.resources_map : k => v if v.resource_type == "null" }
  provisioner "local-exec" {
    command = "echo '###### This is ${each.value.application_name}-${each.key}-${each.value.resource_type} file' > ./${each.value.application_name}-${each.key}-${each.value.resource_type}.txt"
  }
}

resource "null_resource" "tf_cli_rds" {
  for_each = { for k, v in local.resources_map : k => v if v.resource_type == "rds" }
  provisioner "local-exec" {
    command = "echo '###### This is ${each.value.application_name}-${each.key}-${each.value.resource_type} file' > ./${each.value.application_name}-${each.key}-${each.value.resource_type}.txt"
  }
}

resource "null_resource" "tf_cli_ec2" {
  for_each = { for k, v in local.resources_map : k => v if v.resource_type == "ec2" }
  provisioner "local-exec" {
    command = "echo '###### This is ${each.value.application_name}-${each.key}-${each.value.resource_type} file' > ./${each.value.application_name}-${each.key}-${each.value.resource_type}.txt"
  }
}