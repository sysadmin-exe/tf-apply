output "resources_created" {
  value = { for k, v in null_resource.this : k => v.id }
}