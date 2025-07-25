
# // creating secretfor cognito_region
# module "secret_manager_cognito_region" {
#   source               = "./modules/secret_manager"
#   secret_name          = "COGNITO_REGIONS"
#   secret_value         = var.aws_region
#   recovery_window_days = var.recovery_window_days

# }

# // creating secretfor app_client_id
# module "secret_manager_app_client_id" {
#   source               = "./modules/secret_manager"
#   secret_name          = "APP_CLIENT_IDS"
#   secret_value         = module.cognito_user_pool.app_client_id
#   recovery_window_days = var.recovery_window_days

# }

# // creating secretfor user_pool_id
# module "secret_manager_user_pool_id" {
#   source               = "./modules/secret_manager"
#   secret_name          = "USER_POOL_IDS"
#   secret_value         = module.cognito_user_pool.user_pool_id
#   recovery_window_days = var.recovery_window_days

# }

# // creating secretfor dynamo_region
# module "secret_manager_dynamo_region" {
#   source               = "./modules/secret_manager"
#   secret_name          = "DYNAMO_REGIONS"
#   secret_value         = var.aws_region
#   recovery_window_days = var.recovery_window_days

# }


# // creating secretfor dynamo_user_table_name
# module "secret_manager_dynamo_user_table_name" {
#   source               = "./modules/secret_manager"
#   secret_name          = "DYNAMO_USER_TABLE_NAMES"
#   secret_value         = var.dynamo_user_meta_data
#   recovery_window_days = var.recovery_window_days

# }

# // creating secretfor dynamo_file_name
# module "secret_manager_dynamo_file_name" {
#   source               = "./modules/secret_manager"
#   secret_name          = "DYNAMO_FILE_NAMES"
#   secret_value         = var.dynamo_file_meta_data
#   recovery_window_days = var.recovery_window_days

# }


# // creating secretfor google_client_id
# module "secret_manager_google_client_id" {
#   source               = "./modules/secret_manager"
#   secret_name          = "GOOGLE_CLIENT_IDS"
#   secret_value         = var.google_client_id
#   recovery_window_days = var.recovery_window_days

# }


# // creating secretfor google_client_secret
# module "secret_manager_google_client_secret" {
#   source               = "./modules/secret_manager"
#   secret_name          = "GOOGLE_CLIENT_SECRETS"
#   secret_value         = var.google_client_secret
#   recovery_window_days = var.recovery_window_days

# }


# // creating secretfor google_redirect_url
# module "secret_manager_google_redirect_url" {
#   source               = "./modules/secret_manager"
#   secret_name          = "GOOGLE_REDIRECT_URLS"
#   secret_value         = var.google_redirect_url
#   recovery_window_days = var.recovery_window_days

# }
