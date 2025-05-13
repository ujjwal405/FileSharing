resource "aws_api_gateway_authorizer" "lambda_authorizer" {
  name            = "lambda-authorizer"
  rest_api_id     = var.rest_api_id
  authorizer_uri  = var.authorizer_uri
  type            = "REQUEST"
  identity_source = "method.request.header.Authorization,method.request.header.X-Id-Token"
}

resource "aws_lambda_permission" "authorizer" {
  statement_id  = "AllowAPIGatewayInvokeAuthorizer"
  action        = "lambda:InvokeFunction"
  function_name = var.lambda_authorizer_name
  principal     = "apigateway.amazonaws.com"
  # Source ARN targets the specific authorizer
  source_arn = "arn:aws:execute-api:${var.aws_region}:${var.account_id}:${var.rest_api_id}/authorizers/${aws_api_gateway_authorizer.lambda_authorizer.id}"
}
