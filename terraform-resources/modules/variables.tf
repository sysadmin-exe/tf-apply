variable "resources_list" {
  type = list(object({
    application_name = string
    env              = string
    kind             = string
  }))
  description = "List of resources to create"
  validation {
    condition     = length(var.resources_list) > 0
    error_message = "At least one resource must be specified"
  }
}