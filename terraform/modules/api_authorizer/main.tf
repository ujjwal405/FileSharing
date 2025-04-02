resource "aws_api_gateway_authorizer" "lambda_authorizer" {
  name            = "lambda-authorizer"
  rest_api_id     = var.rest_api_id
  authorizer_uri  = var.authorizer_uri
  type            = "REQUEST"
  identity_source = ["method.request.header.Authorization", "method.request.header.X-Id-Token"]
}
