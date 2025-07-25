module "codeVerification_resource" {
  source      = "./modules/api_resource"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  parent_id   = module.file_sharing_gateway.root_resource_id
  path_part   = "codeVerification"
}

module "confirmPassword_resource" {
  source      = "./modules/api_resource"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  parent_id   = module.file_sharing_gateway.root_resource_id
  path_part   = "confirmPassword"
}

module "downloadSignedURL_resource" {
  source      = "./modules/api_resource"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  parent_id   = module.file_sharing_gateway.root_resource_id
  path_part   = "downloadSignedUrl"
}

module "forgetPassword_resource" {
  source      = "./modules/api_resource"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  parent_id   = module.file_sharing_gateway.root_resource_id
  path_part   = "forgetPassword"
}

module "getCode_resource" {
  source      = "./modules/api_resource"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  parent_id   = module.file_sharing_gateway.root_resource_id
  path_part   = "getCode"
}

module "googleCallback_resource" {
  source      = "./modules/api_resource"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  parent_id   = module.file_sharing_gateway.root_resource_id
  path_part   = "googleCallback"
}

module "logout_resource" {
  source      = "./modules/api_resource"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  parent_id   = module.file_sharing_gateway.root_resource_id
  path_part   = "logout"
}

module "myfiles_resource" {
  source      = "./modules/api_resource"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  parent_id   = module.file_sharing_gateway.root_resource_id
  path_part   = "myfiles"
}

module "signin_resource" {
  source      = "./modules/api_resource"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  parent_id   = module.file_sharing_gateway.root_resource_id
  path_part   = "signin"
}

module "signinGoogle_resource" {
  source      = "./modules/api_resource"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  parent_id   = module.file_sharing_gateway.root_resource_id
  path_part   = "signinGoogle"
}

