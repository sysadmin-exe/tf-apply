variable "resource_count" {
  type = number
  description = "Number of resources to create"
  default = 0
  # validation {
  #   condition     = length(var.image_id) > 4 && substr(var.image_id, 0, 4) == "ami-"
  #   error_message = "The image_id value must be a valid AMI id, starting with \"ami-\"."
  # }
}

variable "resource_name" {
  type = string
  description = "File name"
  validation {
    condition     = can(regex("null|ec2|rds", var.resource_name))
    error_message = "The only resource that can be created are null | ec2 | rds "
  }
}