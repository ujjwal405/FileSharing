resource "aws_lambda_permission" "this" {
  statement_id  = "AllowAPIGatewayInvoke-${var.lambda_function_name}-${var.http_method}"
  action        = "lambda:InvokeFunction"
  function_name = var.lambda_function_name
  principal     = "apigateway.amazonaws.com"

  # The source ARN typically includes your API Gateway REST API ID, method, and resource path.
  # Adjust the pattern as needed:
  source_arn = "arn:aws:execute-api:${var.region}:${var.account_id}:${var.rest_api_id}/${var.stage}/${var.http_method}/${var.resource_path}"
}
