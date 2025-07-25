variable "rest_api_id" {}

# variable "resource_ids" {
#   type = map(string)

# }

# variable "methods" {
#   type = list(object({
#     resource_name  = string
#     http_method    = string
#     lambda_arn     = string
#     use_authorizer = optional(bool, false)
#   }))
# }
variable "resource_id" {}

variable "http_method" {}

variable "authorizer_id" {
  type    = string
  default = null
}
variable "use_authorizer" {
  type    = bool
  default = false
}
