variable "resources_list" {
  type = list(object({
    resource_type    = string
    resource_count   = number
    application_name = string
  }))
  description = "List of resources to create"
}