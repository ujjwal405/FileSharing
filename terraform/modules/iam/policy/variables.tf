variable "policy_name" {
  description = "Name of the Lambda function"
  type        = string
}

variable "env_name" {
  description = "Environment (e.g., dev, prod)"
  type        = string
}

variable "policy_statements" {
  description = "List of policy statements for IAM role"
  type = list(object({
    Effect   = string
    Action   = list(string)
    Resource = list(string)
  }))
}
