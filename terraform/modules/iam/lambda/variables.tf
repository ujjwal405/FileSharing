variable "lambda_name" {
  description = "Lambda function name"
  type        = string
}

variable "env_name" {
  description = "The environment (e.g., dev, prod)"
  type        = string
}

variable "lambda_policy_arns" {
  type = list(string)
}
