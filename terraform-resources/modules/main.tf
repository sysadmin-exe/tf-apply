# This is still a test. Not a real thing in fact. Just wanted different ways to test the go app

locals{
  resources_map = {for r in var.resources_list : "${r.application_name}-${r.env}-${r.kind}" => {
    application_name = r.application_name
    kind             = r.kind
    env              = r.env
  }}
}
resource "null_resource" "this" {
  for_each = local.resources_map
  provisioner "local-exec" {
    command = "echo '###### This is ${each.value.application_name}-${each.key}-${each.value.kind} file' > ${path.root}/modules/${each.value.application_name}-${each.key}-${each.value.kind}.txt"
  }
}
