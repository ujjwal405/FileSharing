// creating dynamo db table for storing file_meta_data
module "dynamo_file_meta_data" {
  source         = "./modules/dynamo"
  aws_region     = var.aws_region
  table_name     = var.dynamo_file_meta_data
  hash_key       = var.dynamo_file_meta_data_hash_key
  range_key      = var.dynamo_file_meta_data_range_key
  read_capacity  = var.dynamo_file_meta_data_read_capacity
  write_capacity = var.dynamo_file_meta_data_write_capacity
  attributes     = var.dynamo_file_meta_data_attributes

}


// creating dynamo db table for storing user_meta_data
module "dynamo_user_meta_data" {
  source         = "./modules/dynamo"
  aws_region     = var.aws_region
  table_name     = var.dynamo_user_meta_data
  hash_key       = var.dynamo_user_meta_data_hash_key
  read_capacity  = var.dynamo_user_meta_data_read_capacity
  write_capacity = var.dynamo_user_meta_data_write_capacity
  attributes     = var.dynamo_user_meta_data_attributes

}
