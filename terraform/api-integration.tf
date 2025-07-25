module "codeVerification_integration" {
  source      = "./modules/api_integration"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.codeVerification_resource.resource_id
  http_method = module.codeVerification_method.http_method

  invoke_arn = module.lambda_code_verification.invocation_arn
}

module "confirmPassword_integration" {
  source      = "./modules/api_integration"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.confirmPassword_resource.resource_id
  http_method = module.confirmPassword_method.http_method

  invoke_arn = module.lambda_confirm_password.invocation_arn
}

module "downloadSignedURL_integration" {
  source      = "./modules/api_integration"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.downloadSignedURL_resource.resource_id
  http_method = module.downloadSignedURL_method.http_method

  invoke_arn = module.lambda_download_signed_url.invocation_arn
}

module "forgetPassword_integration" {
  source      = "./modules/api_integration"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.forgetPassword_resource.resource_id
  http_method = module.forgetPassword_method.http_method

  invoke_arn = module.lambda_forget_password.invocation_arn
}

module "getCode_integration" {
  source      = "./modules/api_integration"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.getCode_resource.resource_id
  http_method = module.getCode_method.http_method

  invoke_arn = module.lambda_get_code.invocation_arn
}

module "googleCallback_integration" {
  source      = "./modules/api_integration"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.googleCallback_resource.resource_id
  http_method = module.googleCallback_method.http_method

  invoke_arn = module.lambda_google_callback.invocation_arn
}

module "logout_integration" {
  source      = "./modules/api_integration"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.logout_resource.resource_id
  http_method = module.logout_method.http_method

  invoke_arn = module.lambda_logout.invocation_arn

}

module "myfiles_integration" {
  source      = "./modules/api_integration"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.myfiles_resource.resource_id
  http_method = module.myfiles_method.http_method

  invoke_arn = module.lambda_myfiles.invocation_arn

}

module "signin_integration" {
  source      = "./modules/api_integration"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.signin_resource.resource_id
  http_method = module.signin_method.http_method

  invoke_arn = module.lambda_signin.invocation_arn

}

module "signinGoogle_integration" {
  source      = "./modules/api_integration"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.signinGoogle_resource.resource_id
  http_method = module.signinGoogle_method.http_method

  invoke_arn = module.lambda_signin_google.invocation_arn

}

module "signup_integration" {
  source      = "./modules/api_integration"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.signup_resource.resource_id
  http_method = module.signup_method.http_method

  invoke_arn = module.lambda_signup.invocation_arn

}

module "uploadMetaData_integration" {
  source      = "./modules/api_integration"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.uploadMetaData_resource.resource_id
  http_method = module.uploadMetaData_method.http_method

  invoke_arn = module.lambda_upload_metadata.invocation_arn

}


module "uploadSignedUrl_integration" {
  source      = "./modules/api_integration"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.uploadSignedUrl_resource.resource_id
  http_method = module.uploadSignedUrl_method.http_method

  invoke_arn = module.lambda_upload_signed_url.invocation_arn

}
