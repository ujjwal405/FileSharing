variable "rest_api_id" {
  description = "API Gateway REST API ID"
  type        = string
}

variable "lambda_function_name" {
  description = "Name of the Lambda function"
  type        = string
}

variable "http_method" {
  description = "HTTP method (e.g. GET, POST)"
  type        = string
}

variable "resource_path" {
  description = "Resource path (e.g. /signup, /downloadSignedUrl)"
  type        = string
}

variable "region" {
  description = "AWS region"
  type        = string
}

variable "account_id" {
  description = "AWS account ID"
  type        = string
}

variable "stage" {
  description = "API Gateway stage (e.g., prod, dev)"
  type        = string
}
