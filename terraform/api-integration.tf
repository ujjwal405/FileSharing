module "codeVerification_integration" {
  source      = "./modules/api_integration"
  rest_api_id = module.file_sharing_gateway.rest_api_id
  resource_id = module.codeVerification_resource.resource_id
  http_method = module.codeVerification_method.http_method

  invoke_arn = module.lambda_code_verification.invocation_arn
}
