// Lambda function for authorizer
module "lambda_authorizer" {
  source           = "./modules/lambda"
  lambda_name      = "authorizer"
  env_name         = var.env_name
  function_timeout = var.function_timeout
  lambda_role_arn  = module.iam_role_authorizer.role_arn
  environment_variables = {
    AWS_REGION = var.secret_manager_aws_region
  }

}


// Lambda function for code_verification
module "lambda_code_verification" {
  source           = "./modules/lambda"
  lambda_name      = "code_verification"
  env_name         = var.env_name
  function_timeout = var.function_timeout
  lambda_role_arn  = module.iam_role_code_verification.role_arn

  environment_variables = {
    AWS_REGION = var.secret_manager_aws_region
  }

}

// Lambda function for confirm_password
module "lambda_confirm_password" {
  source           = "./modules/lambda"
  lambda_name      = "confirm_password"
  env_name         = var.env_name
  function_timeout = var.function_timeout
  lambda_role_arn  = module.iam_role_confirm_password.role_arn

  environment_variables = {
    AWS_REGION = var.secret_manager_aws_region
  }
}



// Lambda function for download_signed_url
module "lambda_download_signed_url" {
  source           = "./modules/lambda"
  lambda_name      = "download_signed_url"
  env_name         = var.env_name
  function_timeout = var.function_timeout
  lambda_role_arn  = module.iam_role_download_signed_url.role_arn

  environment_variables = var.download_signed_url_env
}



// Lambda function for forget_password
module "lambda_forget_password" {
  source           = "./modules/lambda"
  lambda_name      = "forget_password"
  env_name         = var.env_name
  function_timeout = var.function_timeout
  lambda_role_arn  = module.iam_role_forget_password.role_arn

  environment_variables = {
    AWS_REGION = var.secret_manager_aws_region
  }
}



// Lambda function for google_callback
module "lambda_google_callback" {
  source           = "./modules/lambda"
  lambda_name      = "google_callback"
  env_name         = var.env_name
  function_timeout = var.function_timeout
  lambda_role_arn  = module.iam_role_google_callback.role_arn

  environment_variables = {
    AWS_REGION = var.secret_manager_aws_region
  }
}


// Lambda function for logout
module "lambda_logout" {
  source           = "./modules/lambda"
  lambda_name      = "logout"
  env_name         = var.env_name
  function_timeout = var.function_timeout
  lambda_role_arn  = module.iam_role_logout.role_arn

  environment_variables = {
    AWS_REGION = var.secret_manager_aws_region
  }
}



// Lambda function for myfiles
module "lambda_myfiles" {
  source           = "./modules/lambda"
  lambda_name      = "myfiles"
  env_name         = var.env_name
  function_timeout = var.function_timeout
  lambda_role_arn  = module.iam_role_my_files.role_arn

  environment_variables = {
    AWS_REGION = var.secret_manager_aws_region
  }
}


// Lambda function for signin
module "lambda_signin" {
  source           = "./modules/lambda"
  lambda_name      = "signin"
  env_name         = var.env_name
  function_timeout = var.function_timeout
  lambda_role_arn  = module.iam_role_sign_in.role_arn

  environment_variables = {
    AWS_REGION = var.secret_manager_aws_region
  }
}


// Lambda function for signin_google
module "lambda_signin_google" {
  source           = "./modules/lambda"
  lambda_name      = "signin_google"
  env_name         = var.env_name
  function_timeout = var.function_timeout
  lambda_role_arn  = module.iam_role_sign_in_google.role_arn

  environment_variables = {
    AWS_REGION = var.secret_manager_aws_region
  }
}


// Lambda function for signup
module "lambda_signup" {
  source           = "./modules/lambda"
  lambda_name      = "signup"
  env_name         = var.env_name
  function_timeout = var.function_timeout
  lambda_role_arn  = module.iam_role_sign_up.role_arn

  environment_variables = {
    AWS_REGION = var.secret_manager_aws_region
  }
}

// Lambda function for upload_metadata
module "lambda_upload_metadata" {
  source           = "./modules/lambda"
  lambda_name      = "upload_metadata"
  env_name         = var.env_name
  function_timeout = var.function_timeout
  lambda_role_arn  = module.iam_role_upload_metadata.role_arn

  environment_variables = {
    AWS_REGION = var.secret_manager_aws_region
  }
}



// Lambda function for upload_signed_url
module "lambda_upload_signed_url" {
  source           = "./modules/lambda"
  lambda_name      = "upload_signed_url"
  env_name         = var.env_name
  function_timeout = var.function_timeout
  lambda_role_arn  = module.iam_role_upload_signed_url.role_arn

  environment_variables = var.upload_signed_url_env
}

// Lambda function for get_code
module "lambda_get_code" {
  source           = "./modules/lambda"
  lambda_name      = "get_code"
  env_name         = var.env_name
  function_timeout = var.function_timeout
  lambda_role_arn  = module.iam_role_get_code.role_arn

}
