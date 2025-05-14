variable "rest_api_id" {}

variable "resources" {
  type = list(object({
    name      = string
    parent_id = string
  }))
}



