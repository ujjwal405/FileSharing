
locals {
  lambda_basic_execution_policy = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}
// IAM role for authorizer_lambda
module "iam_role_authorizer" {
  source      = "./modules/iam/lambda"
  lambda_name = "authorizer"
  env_name    = var.env_name
  lambda_policy_arns = [
    module.iam_policy_dynamo_get_item_user_table.policy_arn,
    module.iam_policy_secret_manager_get_item_cognito_region.policy_arn,
    module.iam_policy_secret_manager_get_item_app_client_id.policy_arn,
    module.iam_policy_secret_manager_get_item_user_pool_id.policy_arn,
    module.iam_policy_secret_manager_get_item_dynamo_region.policy_arn,
    module.iam_policy_secret_manager_get_item_dynamo_user_table_name.policy_arn,
    module.iam_policy_cognito_admin_delete.policy_arn,
    module.iam_policy_cognito_initiate_auth.policy_arn,
    local.lambda_basic_execution_policy

  ]

}


// IAM role for code_verification_lambda
module "iam_role_code_verification" {
  source      = "./modules/iam/lambda"
  lambda_name = "code_verification"
  env_name    = var.env_name
  lambda_policy_arns = [
    module.iam_policy_cognito_confirm_sign_up.policy_arn,
    local.lambda_basic_execution_policy
  ]
}


// IAM role for confirm_password
module "iam_role_confirm_password" {
  source      = "./modules/iam/lambda"
  lambda_name = "confirm_password"
  env_name    = var.env_name
  lambda_policy_arns = [
    module.iam_policy_secret_manager_get_item_cognito_region.policy_arn,
    module.iam_policy_secret_manager_get_item_app_client_id.policy_arn,
    module.iam_policy_cognito_confirm_forget_password.policy_arn,
    local.lambda_basic_execution_policy

  ]
}


// IAM role for download_signed_url
module "iam_role_download_signed_url" {
  source      = "./modules/iam/lambda"
  lambda_name = "download_signed_url"
  env_name    = var.env_name
  lambda_policy_arns = [
    module.iam_policy_s3_user_file_get_object.policy_arn,
    local.lambda_basic_execution_policy

  ]
}



// IAM role for forget_password
module "iam_role_forget_password" {
  source      = "./modules/iam/lambda"
  lambda_name = "forget_password"
  env_name    = var.env_name
  lambda_policy_arns = [
    module.iam_policy_secret_manager_get_item_cognito_region.policy_arn,
    module.iam_policy_secret_manager_get_item_app_client_id.policy_arn,
    module.iam_policy_cognito_forget_password.policy_arn,
    local.lambda_basic_execution_policy

  ]
}


// IAM role for logout
module "iam_role_logout" {
  source      = "./modules/iam/lambda"
  lambda_name = "logout"
  env_name    = var.env_name
  lambda_policy_arns = [
    module.iam_policy_secret_manager_get_item_cognito_region.policy_arn,
    module.iam_policy_secret_manager_get_item_app_client_id.policy_arn,
    module.iam_policy_secret_manager_get_item_user_pool_id.policy_arn,
    module.iam_policy_secret_manager_get_item_dynamo_region.policy_arn,
    module.iam_policy_secret_manager_get_item_dynamo_user_table_name.policy_arn,
    module.iam_policy_cognito_admin_delete.policy_arn,
    module.iam_policy_cognito_revoke_token.policy_arn,
    local.lambda_basic_execution_policy

  ]
}

// IAM role for signIn
module "iam_role_sign_in" {
  source      = "./modules/iam/lambda"
  lambda_name = "signin"
  env_name    = var.env_name
  lambda_policy_arns = [
    module.iam_policy_secret_manager_get_item_cognito_region.policy_arn,
    module.iam_policy_secret_manager_get_item_app_client_id.policy_arn,
    module.iam_policy_secret_manager_get_item_dynamo_region.policy_arn,
    module.iam_policy_secret_manager_get_item_dynamo_user_table_name.policy_arn,
    module.iam_policy_cognito_initiate_auth.policy_arn,
    module.iam_policy_dynamo_put_item_user_table.policy_arn,
    local.lambda_basic_execution_policy

  ]
}



// IAM role for signUp
module "iam_role_sign_up" {
  source      = "./modules/iam/lambda"
  lambda_name = "signup"
  env_name    = var.env_name
  lambda_policy_arns = [
    module.iam_policy_secret_manager_get_item_cognito_region.policy_arn,
    module.iam_policy_secret_manager_get_item_app_client_id.policy_arn,
    module.iam_policy_secret_manager_get_item_user_pool_id.policy_arn,
    module.iam_policy_cognito_admin_get_user.policy_arn,
    module.iam_policy_cognito_sign_up.policy_arn,
    local.lambda_basic_execution_policy

  ]
}


// IAM  role for upload_metadata
module "iam_role_upload_metadata" {
  source      = "./modules/iam/lambda"
  lambda_name = "upload_metadata"
  env_name    = var.env_name
  lambda_policy_arns = [
    module.iam_policy_secret_manager_get_item_dynamo_region.policy_arn,
    module.iam_policy_secret_manager_get_item_dynamo_file_name.policy_arn,
    module.iam_policy_dynamo_put_item_file_table.policy_arn,
    local.lambda_basic_execution_policy

  ]
}



// IAM  role for upload_signed_url
module "iam_role_upload_signed_url" {
  source      = "./modules/iam/lambda"
  lambda_name = "upload_signed_url"
  env_name    = var.env_name
  lambda_policy_arns = [
    module.iam_policy_s3_user_file_put_object.policy_arn,
    local.lambda_basic_execution_policy

  ]
}



// IAM  role for myfiles
module "iam_role_my_files" {
  source      = "./modules/iam/lambda"
  lambda_name = "myfiles"
  env_name    = var.env_name
  lambda_policy_arns = [
    module.iam_policy_secret_manager_get_item_dynamo_region.policy_arn,
    module.iam_policy_secret_manager_get_item_dynamo_file_name.policy_arn,
    module.iam_policy_dynamo_query_file_table.policy_arn,
    local.lambda_basic_execution_policy

  ]

}



// IAM  role for sign_in_google
module "iam_role_sign_in_google" {
  source      = "./modules/iam/lambda"
  lambda_name = "signin_google"
  env_name    = var.env_name
  lambda_policy_arns = [
    module.iam_policy_secret_manager_get_item_google_credentials.policy_arn,

    local.lambda_basic_execution_policy

  ]

}


// IAM role for google_callback
module "iam_role_google_callback" {
  source      = "./modules/iam/lambda"
  lambda_name = "google_callback"
  env_name    = var.env_name
  lambda_policy_arns = [
    module.iam_policy_secret_manager_get_item_cognito_region.policy_arn,
    module.iam_policy_secret_manager_get_item_app_client_id.policy_arn,
    module.iam_policy_secret_manager_get_item_user_pool_id.policy_arn,
    module.iam_policy_secret_manager_get_item_dynamo_region.policy_arn,
    module.iam_policy_secret_manager_get_item_dynamo_user_table_name.policy_arn,
    module.iam_policy_secret_manager_get_item_google_credentials.policy_arn,

    module.iam_policy_cognito_initiate_auth.policy_arn,
    module.iam_policy_cognito_admin_delete_create_get_user.arn,
    module.iam_policy_dynamo_put_item_user_table.policy_arn,
    local.lambda_basic_execution_policy

  ]
}




// IAM  role for get_code
module "iam_role_get_code" {
  source      = "./modules/iam/lambda"
  lambda_name = "get_code"
  env_name    = var.env_name
  lambda_policy_arns = [
    local.lambda_basic_execution_policy

  ]

}
