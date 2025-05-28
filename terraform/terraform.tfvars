access_token_validity = 1

allow_headers = ["Content-Type", "Authorization", "X-Id-Token"]
allow_methods = ["GET", "POST", "OPTIONS"]

api_gateway_name = "FileUpload-GTW"
api_stage_name   = "v1"

download_signed_url_env = {
  BUCKET_NAME = "filesharing-user-file-bucket"
  ENVIRONMENT = "prod"
  REGION      = "ap-south-1"
}

dynamo_file_meta_data = "FileTable"
dynamo_file_meta_data_attributes = [
  # { name = "s3filename", type = "S" },
  { name = "created_at", type = "S" },
  { name = "email", type = "S" },
  # { name = "filename", type = "S" }
]
dynamo_file_meta_data_hash_key  = "email"
dynamo_file_meta_data_range_key = "created_at"
# dynamo_file_meta_data_read_capacity  = 5
# dynamo_file_meta_data_write_capacity = 5

dynamo_user_meta_data = "UserTable"
dynamo_user_meta_data_attributes = [
  {
    name = "email"
    type = "S"
  },
  # {
  #   name = "refresh_token"
  #   type = "S"
  # },
  # {
  #   name = "expires_at"
  #   type = "N"
  # }
]
dynamo_user_meta_data_hash_key = "email"
# dynamo_user_meta_data_read_capacity  = 5
# dynamo_user_meta_data_write_capacity = 5

endpoint_type = "REGIONAL"

id_token_validity = 1

refresh_token_validity = 1

s3_file_upload = "filesharing-user-file-bucket"

upload_signed_url_env = {
  BUCKET_NAME = "filesharing-user-file-bucket"
  ENVIRONMENT = "prod"
  REGION      = "ap-south-1"
}
