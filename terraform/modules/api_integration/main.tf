resource "aws_api_gateway_integration" "lambda_get" {
  rest_api_id = var.rest_api_id
  resource_id = var.resource_id
  http_method = var.http_method

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = var.invoke_arn
}
