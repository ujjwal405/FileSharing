variable "lambda_name" {
  description = "Lambda function name"
  type        = string
}

variable "env_name" {
  description = "The environment (e.g., dev, prod)"
  type        = string
}

variable "function_runtime" {
  description = "Lambda function runtime"
  type        = string
  default     = "go1.x"
}

variable "function_timeout" {
  description = "Lambda function timeout"
  type        = number
}


variable "lambda_role_arn" {
  description = "ARN of the IAM role for Lambda"
  type        = string
}


variable "environment_variables" {
  type    = map(string)
  default = {} # Empty map
}
