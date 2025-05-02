variable "aws_region" {
  type    = string
  default = "ap-south-1"
}

variable "s3_file_upload" {}


variable "dynamo_file_meta_data" {}

variable "dynamo_file_meta_data_hash_key" {}

variable "dynamo_file_meta_data_range_key" {}

variable "dynamo_file_meta_data_read_capacity" {
  type = int
}

variable "dynamo_file_meta_data_write_capacity" {
  type = int
}

variable "dynamo_file_meta_data_attributes" {
  type = list(object({
    name = string
    type = string
  }))
}

variable "dynamo_user_meta_data" {}

variable "dynamo_user_meta_data_hash_key" {}

variable "dynamo_user_meta_data_read_capacity" {
  type = int
}

variable "dynamo_user_meta_data_write_capacity" {
  type = int
}

variable "dynamo_user_meta_data_attributes" {
  type = list(object({
    name = string
    type = string
  }))
}

variable "env_name" {
  type    = string
  default = "prod"
}

variable "secret_manager_aws_region" {
  type    = string
  default = "ap-south_1"
}
variable "upload_signed_url_env" {
  type = map(string)
}



variable "download_signed_url_env" {
  type = map(string)
}

variable "google_client_id" {}

variable "google_client_secret" {}

variable "google_redirect_url" {}



variable "function_timeout" {
  type    = number
  default = 60
}

variable "api_gateway_name" {}
variable "endpoint_type" {}
variable "api_stage_name" {}
variable "allow_headers" {
  type = list(string)
}
variable "allow_methods" {
  type = list(string)
}

