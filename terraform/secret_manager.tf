
// creating secretfor cognito_region
module "secret_manager_cognito_region" {
  source       = "./modules/secret_manager"
  secret_name  = "COGNITO_REGION"
  secret_value = var.aws_region

}

// creating secretfor app_client_id
module "secret_manager_app_client_id" {
  source       = "./modules/secret_manager"
  secret_name  = "APP_CLIENT_ID"
  secret_value = module.cognito_user_pool.app_client_id

}

// creating secretfor user_pool_id
module "secret_manager_user_pool_id" {
  source       = "./modules/secret_manager"
  secret_name  = "USER_POOL_ID"
  secret_value = module.cognito_user_pool.user_pool_id

}

// creating secretfor dynamo_region
module "secret_manager_dynamo_region" {
  source       = "./modules/secret_manager"
  secret_name  = "DYNAMO_REGION"
  secret_value = var.aws_region

}


// creating secretfor dynamo_user_table_name
module "secret_manager_dynamo_user_table_name" {
  source       = "./modules/secret_manager"
  secret_name  = "DYNAMO_USER_TABLE_NAME"
  secret_value = var.dynamo_user_meta_data

}

// creating secretfor dynamo_file_name
module "secret_manager_dynamo_file_name" {
  source       = "./modules/secret_manager"
  secret_name  = "DYNAMO_FILE_NAME"
  secret_value = var.dynamo_file_meta_data

}


// creating secretfor google_client_id
module "secret_manager_google_client_id" {
  source       = "./modules/secret_manager"
  secret_name  = "GOOGLE_CLIENT_ID"
  secret_value = var.google_client_id

}


// creating secretfor google_client_secret
module "secret_manager_google_client_secret" {
  source       = "./modules/secret_manager"
  secret_name  = "GOOGLE_CLIENT_SECRET"
  secret_value = var.google_client_secret

}


// creating secretfor google_redirect_url
module "secret_manager_google_redirect_url" {
  source       = "./modules/secret_manager"
  secret_name  = "GOOGLE_REDIRECT_URL"
  secret_value = var.google_redirect_url

}
