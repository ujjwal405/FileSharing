variable "aws_region" {
  type        = string
  description = "The AWS region to deploy dynamo"
}


variable "table_name" {
  type        = string
  description = "The name for dynamo table"
}

variable "hash_key" {
  type        = string
  description = "The name of hash key"
}


variable "range_key" {
  type    = string
  default = null
}


# variable "read_capacity" {
#   type        = number
#   description = "The read capacity of table"
# }

# variable "write_capacity" {
#   type        = number
#   description = "The write capacity of table"
# }

variable "attributes" {
  type = list(object({
    name = string
    type = string
  }))
  description = "List of attributes for the DynamoDB table"
}


