resource "aws_api_gateway_deployment" "this" {
  rest_api_id = var.rest_api_id
  triggers = {
    redeployment = sha1(jsonencode({
      timestamp = timestamp()
    }))
  }

  lifecycle {
    create_before_destroy = true
  }
}
