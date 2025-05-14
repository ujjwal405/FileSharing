locals {
  # resource_path = [
  #   "/codeVerification",
  #   "/confirmPassword",
  #   "/downloadSignedUrl",
  #   "/forgetPassword",
  #   "/getCode",
  #   "/googleCallback",
  #   "/logout",
  #   "/myfiles",
  #   "/signin",
  #   "/signinGoogle",
  #   "/signup",
  #   "/uploadMetaData",
  #   "/uploadSignedUrl"
  # ]
  full_resource_map = module.api_resource.resource_ids

  cors_map = {
    "/codeVerification"  = lookup(local.full_resource_map, "/codeVerification", null)
    "/confirmPassword"   = lookup(local.full_resource_map, "/confirmPassword", null)
    "/downloadSignedUrl" = lookup(local.full_resource_map, "/downloadSignedUrl", null)
    "/forgetPassword"    = lookup(local.full_resource_map, "/forgetPassword", null)
    "/getCode"           = lookup(local.full_resource_map, "/getCode", null)
    "/googleCallback"    = lookup(local.full_resource_map, "/googleCallback", null)
    "/logout"            = lookup(local.full_resource_map, "/logout", null)
    "/myfiles"           = lookup(local.full_resource_map, "/myfiles", null)
    "/signin"            = lookup(local.full_resource_map, "/signin", null)
    "/signinGoogle"      = lookup(local.full_resource_map, "/signinGoogle", null)
    "/signup"            = lookup(local.full_resource_map, "/signup", null)
    "/uploadMetaData"    = lookup(local.full_resource_map, "/uploadMetaData", null)
    "/uploadSignedUrl"   = lookup(local.full_resource_map, "/uploadSignedUrl", null)
  }

}


// creating api_gateway
module "file_sharing_gateway" {
  source        = "./modules/api_gateway"
  api_name      = var.api_gateway_name
  endpoint_type = var.endpoint_type
}


// creating api_gateway_authorizer
module "api_authorizer" {
  source                 = "./modules/api_authorizer"
  rest_api_id            = module.file_sharing_gateway.rest_api_id
  authorizer_uri         = module.lambda_authorizer.invoke_arn
  aws_region             = var.aws_region
  account_id             = data.aws_caller_identity.current.account_id
  lambda_authorizer_name = module.lambda_authorizer.function_name
}




//creating api_gateway_resource
module "api_resource" {
  source      = "./modules/api_resource"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resources = [
    {
      name      = "codeVerification"
      parent_id = module.file_sharing_gateway.root_resource_id
    },
    {
      name      = "confirmPassword"
      parent_id = module.file_sharing_gateway.root_resource_id
    },
    {
      name      = "downloadSignedUrl"
      parent_id = module.file_sharing_gateway.root_resource_id
    },
    {
      name      = "forgetPassword"
      parent_id = module.file_sharing_gateway.root_resource_id
    },
    {
      name      = "getCode"
      parent_id = module.file_sharing_gateway.root_resource_id
    },
    {
      name      = "googleCallback"
      parent_id = module.file_sharing_gateway.root_resource_id
      }, {
      name      = "logout"
      parent_id = module.file_sharing_gateway.root_resource_id
    },
    {
      name      = "myfiles"
      parent_id = module.file_sharing_gateway.root_resource_id
    },
    {
      name      = "signin"
      parent_id = module.file_sharing_gateway.root_resource_id
    },
    {
      name      = "signinGoogle"
      parent_id = module.file_sharing_gateway.root_resource_id
    },
    {
      name      = "signup"
      parent_id = module.file_sharing_gateway.root_resource_id
    },
    {
      name      = "uploadMetaData"
      parent_id = module.file_sharing_gateway.root_resource_id
    },
    {
      name      = "uploadSignedUrl"
      parent_id = module.file_sharing_gateway.root_resource_id
    }

  ]

}



// creating api_method
module "api_method" {
  source        = "./modules/api_method"
  rest_api_id   = module.file_sharing_gateway.rest_api_id
  resource_ids  = module.api_resource.resource_ids
  authorizer_id = module.api_authorizer.authorizer_id
  methods = [
    {
      resource_name = "/codeVerification" # 👈 Must match the `res.path` key in output!
      http_method   = "POST"
      lambda_arn    = module.lambda_code_verification.invoke_arn
    },
    {
      resource_name = "/confirmPassword"
      http_method   = "POST"
      lambda_arn    = module.lambda_confirm_password.invoke_arn

    },
    {
      resource_name = "/downloadSignedUrl"
      http_method   = "GET"
      lambda_arn    = module.lambda_download_signed_url.invoke_arn
    },
    {
      resource_name = "/forgetPassword"
      http_method   = "POST"
      lambda_arn    = module.lambda_forget_password.invoke_arn

    },
    {
      resource_name = "/getCode"
      http_method   = "POST"
      lambda_arn    = module.lambda_get_code.invoke_arn


    },
    {
      resource_name = "/googleCallback"
      http_method   = "POST"
      lambda_arn    = module.lambda_google_callback.invoke_arn

    },
    {
      resource_name  = "/logout"
      http_method    = "POST"
      lambda_arn     = module.lambda_logout.invoke_arn
      use_authorizer = true

    },
    {
      resource_name  = "/myfiles"
      http_method    = "GET"
      lambda_arn     = module.lambda_myfiles.invoke_arn
      use_authorizer = true

    },
    {
      resource_name = "/signin"
      http_method   = "POST"
      lambda_arn    = module.lambda_signin.invoke_arn

    },
    {
      resource_name = "/signinGoogle"
      http_method   = "POST"
      lambda_arn    = module.lambda_signin_google.invoke_arn

    },
    {
      resource_name = "/signup"
      http_method   = "POST"
      lambda_arn    = module.lambda_signup.invoke_arn

    },

    {
      resource_name  = "/uploadMetaData"
      http_method    = "POST"
      lambda_arn     = module.lambda_upload_metadata.invoke_arn
      use_authorizer = true

    },
    {
      resource_name  = "/uploadSignedUrl"
      http_method    = "GET"
      lambda_arn     = module.lambda_upload_signed_url.invoke_arn
      use_authorizer = true

    },
  ]
}


// api_gateway_deployment
module "api_deployment" {
  source      = "./modules/api_deployment"
  rest_api_id = module.file_sharing_gateway.rest_api_id
}


// api_gateway_stage
module "api_stage" {
  source        = "./modules/api_gateway_stage"
  rest_api_id   = module.file_sharing_gateway.rest_api_id
  deployment_id = module.api_deployment.deployment_id
  stage_name    = var.api_stage_name

}


// enable cors
# module "cors" {
#   for_each = module.api_resource.resource_ids

#   source  = "squidfunk/api-gateway-enable-cors/aws"
#   version = "0.3.3"

#   api_id          = module.file_sharing_gateway.rest_api_id
#   api_resource_id = each.value
#   allow_headers   = var.allow_headers
#   allow_methods   = var.allow_methods
#   allow_origin    = "https://filesharing.ujjwalsilwal123.com.np"
# }


module "cors" {
  source   = "squidfunk/api-gateway-enable-cors/aws"
  version  = "0.3.3"
  for_each = local.cors_map

  api_id          = module.file_sharing_gateway.rest_api_id
  api_resource_id = each.value
  allow_headers   = var.allow_headers
  allow_methods   = var.allow_methods
  allow_origin    = "https://filesharing.ujjwalsilwal123.com.np"
}
