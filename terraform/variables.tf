variable "access_token_validity" {
  type = number
}

variable "allow_headers" {
  type = list(string)
}

variable "allow_methods" {
  type = list(string)
}

variable "api_gateway_name" {}

variable "api_stage_name" {}

variable "aws_region" {
  type    = string
  default = "ap-south-1"
}
variable "cloudflare_api_token" {
  type = string
}

variable "cloudflare_zone_id" {}
variable "cloudflare_sub_domain_name" {}
variable "download_signed_url_env" {
  type = map(string)
}


variable "dynamo_file_meta_data" {}

variable "dynamo_file_meta_data_attributes" {
  type = list(object({
    name = string
    type = string
  }))
}

variable "dynamo_file_meta_data_hash_key" {}

variable "dynamo_file_meta_data_range_key" {}

# variable "dynamo_file_meta_data_read_capacity" {
#   type = number
# }

# variable "dynamo_file_meta_data_write_capacity" {
#   type = number
# }

variable "dynamo_user_meta_data" {}

variable "dynamo_user_meta_data_attributes" {
  type = list(object({
    name = string
    type = string
  }))
}

variable "dynamo_user_meta_data_hash_key" {}

# variable "dynamo_user_meta_data_read_capacity" {
#   type = number
# }

# variable "dynamo_user_meta_data_write_capacity" {
#   type = number
# }

variable "endpoint_type" {}

variable "env_name" {
  type    = string
  default = "prod"
}

variable "function_timeout" {
  type    = number
  default = 60
}

variable "google_client_id" {}

variable "google_client_secret" {}

variable "google_redirect_url" {}

variable "id_token_validity" {
  type = number
}

variable "refresh_token_validity" {
  type = number
}
variable "s3_file_upload" {}

variable "secret_manager_aws_region" {
  type    = string
  default = "ap-south_1"
}

variable "upload_signed_url_env" {
  type = map(string)
}



variable "backend_domain_name" {}
