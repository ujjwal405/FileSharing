resource "aws_api_gateway_method" "methods" {
  for_each      = { for idx, method in var.methods : idx => method }
  rest_api_id   = var.rest_api_id
  resource_id   = var.resource_ids[each.value.resource_name]
  http_method   = each.value.http_method
  authorization = each.value.use_authorizer ? "CUSTOM" : "NONE"
  authorizer_id = each.value.use_authorizer == true ? aws_api_gateway_authorizer.lambda_authorizer.id : null
}

resource "aws_api_gateway_integration" "integrations" {
  for_each                = aws_api_gateway_method.methods
  rest_api_id             = var.rest_api_id
  resource_id             = each.value.resource_id
  http_method             = each.value.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = var.methods[each.key].lambda_arn
}
