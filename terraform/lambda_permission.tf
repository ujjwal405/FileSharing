module "lambda_permission_code_verification" {
  source              = "./modules/iam/lambda_permission"
  rest_api_id         = module.file_sharing_gateway.rest_api_id
  lambda_function_name = module.lambda_code_verification.function_name  # or .invoke_arn if preferred
  http_method         = "POST"
  resource_path       = "codeVerification"   # Without a leading slash
  stage               = var.api_stage_name # e.g., "prod"
  region              = var.aws_region
  account_id          = data.aws_caller_identity.current.account_id
}


module "lambda_permission_confirm_password" {
  source              = "./modules/iam/lambda_permission"
  rest_api_id         = module.file_sharing_gateway.rest_api_id
  lambda_function_name = module.lambda_confirm_password.function_name  # or .invoke_arn if preferred
  http_method         = "POST"
  resource_path       = "confirmPassword"   # Without a leading slash
  stage               = var.api_stage_name # e.g., "prod"
  region              = var.aws_region
  account_id          = data.aws_caller_identity.current.account_id
}

module "lambda_permission_download_signed_url" {
  source              = "./modules/iam/lambda_permission"
  rest_api_id         = module.file_sharing_gateway.rest_api_id
  lambda_function_name = module.lambda_download_signed_url.function_name  # or .invoke_arn if preferred
  http_method         = "GET"
  resource_path       = "downloadSignedUrl"   # Without a leading slash
  stage               = var.api_stage_name # e.g., "prod"
  region              = var.aws_region
  account_id          = data.aws_caller_identity.current.account_id
}

module "lambda_permission_forget_password" {
  source              = "./modules/iam/lambda_permission"
  rest_api_id         = module.file_sharing_gateway.rest_api_id
  lambda_function_name = module.lambda_forget_password.function_name  # or .invoke_arn if preferred
  http_method         = "POST"
  resource_path       = "forgetPassword"   # Without a leading slash
  stage               = var.api_stage_name # e.g., "prod"
  region              = var.aws_region
  account_id          = data.aws_caller_identity.current.account_id
}

module "lambda_permission_get_code" {
  source              = "./modules/iam/lambda_permission"
  rest_api_id         = module.file_sharing_gateway.rest_api_id
  lambda_function_name = module.lambda_get_code.function_name  # or .invoke_arn if preferred
  http_method         = "GET"
  resource_path       = "getCode"   # Without a leading slash
  stage               = var.api_stage_name # e.g., "prod"
  region              = var.aws_region
  account_id          = data.aws_caller_identity.current.account_id
}

module "lambda_permission_google_callback" {
  source              = "./modules/iam/lambda_permission"
  rest_api_id         = module.file_sharing_gateway.rest_api_id
  lambda_function_name = module.lambda_google_callback.function_name  # or .invoke_arn if preferred
  http_method         = "GET"
  resource_path       = "googleCallback"   # Without a leading slash
  stage               = var.api_stage_name # e.g., "prod"
  region              = var.aws_region
  account_id          = data.aws_caller_identity.current.account_id
}

module "lambda_permission_logout" {
  source              = "./modules/iam/lambda_permission"
  rest_api_id         = module.file_sharing_gateway.rest_api_id
  lambda_function_name = module.lambda_logout.function_name  # or .invoke_arn if preferred
  http_method         = "POST"
  resource_path       = "logout"   # Without a leading slash
  stage               = var.api_stage_name # e.g., "prod"
  region              = var.aws_region
  account_id          = data.aws_caller_identity.current.account_id
}

module "lambda_permission_my_files" {
  source              = "./modules/iam/lambda_permission"
  rest_api_id         = module.file_sharing_gateway.rest_api_id
  lambda_function_name = module.lambda_myfiles.function_name  # or .invoke_arn if preferred
  http_method         = "GET"
  resource_path       = "myfiles"   # Without a leading slash
  stage               = var.api_stage_name # e.g., "prod"
  region              = var.aws_region
  account_id          = data.aws_caller_identity.current.account_id
}

module "lambda_permission_sign_in" {
  source              = "./modules/iam/lambda_permission"
  rest_api_id         = module.file_sharing_gateway.rest_api_id
  lambda_function_name = module.lambda_signin.function_name  # or .invoke_arn if preferred
  http_method         = "POST"
  resource_path       = "signin"   # Without a leading slash
  stage               = var.api_stage_name # e.g., "prod"
  region              = var.aws_region
  account_id          = data.aws_caller_identity.current.account_id
}


module "lambda_permission_sign_in_google" {
  source              = "./modules/iam/lambda_permission"
  rest_api_id         = module.file_sharing_gateway.rest_api_id
  lambda_function_name = module.lambda_signin_google.function_name  # or .invoke_arn if preferred
  http_method         = "POST"
  resource_path       = "signinGoogle"   # Without a leading slash
  stage               = var.api_stage_name # e.g., "prod"
  region              = var.aws_region
  account_id          = data.aws_caller_identity.current.account_id
}

module "lambda_permission_signup" {
  source              = "./modules/iam/lambda_permission"
  rest_api_id         = module.file_sharing_gateway.rest_api_id
  lambda_function_name = module.lambda_signup.function_name  # or .invoke_arn if preferred
  http_method         = "POST"
  resource_path       = "signup"   # Without a leading slash
  stage               = var.api_stage_name # e.g., "prod"
  region              = var.aws_region
  account_id          = data.aws_caller_identity.current.account_id
}

module "lambda_permission_upload_meta_data" {
  source              = "./modules/iam/lambda_permission"
  rest_api_id         = module.file_sharing_gateway.rest_api_id
  lambda_function_name = module.lambda_upload_metadata.function_name  # or .invoke_arn if preferred
  http_method         = "POST"
  resource_path       = "uploadMetaData"   # Without a leading slash
  stage               = var.api_stage_name # e.g., "prod"
  region              = var.aws_region
  account_id          = data.aws_caller_identity.current.account_id
}

module "lambda_permission_upload_signed_url" {
  source              = "./modules/iam/lambda_permission"
  rest_api_id         = module.file_sharing_gateway.rest_api_id
  lambda_function_name = module.lambda_upload_signed_url.function_name  # or .invoke_arn if preferred
  http_method         = "GET"
  resource_path       = "uploadSignedUrl"   # Without a leading slash
  stage               = var.api_stage_name # e.g., "prod"
  region              = var.aws_region
  account_id          = data.aws_caller_identity.current.account_id
}