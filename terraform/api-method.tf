module "codeVerification_method" {
  source      = "./modules/api_method"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.codeVerification_resource.resource_id
  http_method = "POST"

}

module "confirmPassword_method" {
  source      = "./modules/api_method"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.confirmPassword_resource.resource_id
  http_method = "POST"

}

module "downloadSignedURL_method" {
  source      = "./modules/api_method"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.downloadSignedURL_resource.resource_id
  http_method = "GET"

}

module "forgetPassword_method" {
  source      = "./modules/api_method"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.forgetPassword_resource.resource_id
  http_method = "POST"

}

module "getCode_method" {
  source      = "./modules/api_method"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.getCode_resource.resource_id
  http_method = "POST"

}

module "googleCallback_method" {
  source      = "./modules/api_method"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.googleCallback_resource.resource_id
  http_method = "POST"

}

module "logout_method" {
  source      = "./modules/api_method"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.logout_resource.resource_id
  http_method = "POST"

}

module "myfiles_method" {
  source      = "./modules/api_method"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.myfiles_resource.resource_id
  http_method = "GET"

}

module "signin_method" {
  source      = "./modules/api_method"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.signin_resource.resource_id
  http_method = "POST"

}

module "signinGoogle_method" {
  source      = "./modules/api_method"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.signinGoogle_resource.resource_id
  http_method = "POST"

}

module "signup_method" {
  source      = "./modules/api_method"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.signup_resource.resource_id
  http_method = "POST"

}

module "uploadMetaData_method" {
  source      = "./modules/api_method"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.uploadMetaData_resource.resource_id
  http_method = "POST"

}

module "uploadSignedUrl_method" {
  source      = "./modules/api_method"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.uploadSignedUrl_resource.resource_id
  http_method = "GET"

}
