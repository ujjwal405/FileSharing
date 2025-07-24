module "ssm_cognito_region" {
  source          = "./modules/ssm_parameter"
  parameter_name  = "/myapp/cognito/region"
  parameter_value = var.aws_region

}

module "ssm_app_client_id" {
  source          = "./modules/ssm_parameter"
  parameter_name  = "/myapp/cognito/appClientID"
  parameter_value = module.cognito_user_pool.app_client_id

}

module "ssm_user_pool_id" {
  source          = "./modules/ssm_parameter"
  parameter_name  = "/myapp/cognito/userPoolID"
  parameter_value = module.cognito_user_pool.user_pool_id

}

module "ssm_dynamo_region" {
  source          = "./modules/ssm_parameter"
  parameter_name  = "/myapp/dynamo/dynamoRegion"
  parameter_value = var.aws_region
}

module "ssm_dynamo_user_table" {
  source          = "./modules/ssm_parameter"
  parameter_name  = "/myapp/dynamo/dynamoUserTableName"
  parameter_value = var.dynamo_user_meta_data
}

module "ssm_dynamo_file_table" {
  source          = "./modules/ssm_parameter"
  parameter_name  = "/myapp/dynamo/dynamoFileName"
  parameter_value = var.dynamo_file_meta_data
}

module "ssm_google_client_id" {
  source          = "./modules/ssm_parameter"
  parameter_name  = "/myapp/google/googleClientID"
  parameter_value = var.google_client_id
}

module "ssm_google_client_secret" {
  source          = "./modules/ssm_parameter"
  parameter_name  = "/myapp/google/googleClientSecret"
  parameter_value = var.google_client_secret
}

module "ssm_google_redirect_url" {
  source          = "./modules/ssm_parameter"
  parameter_name  = "/myapp/google/googleRedirectURL"
  parameter_value = var.google_redirect_url
}
