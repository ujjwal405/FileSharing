
data "aws_acm_certificate" "this" {
  domain      = var.domain_name
  statuses    = ["ISSUED"]
  most_recent = true
}

resource "aws_api_gateway_domain_name" "this" {
  domain_name     = var.domain_name
  certificate_arn = data.aws_acm_certificate.this.arn

  endpoint_configuration {
    types = [var.endpoint_type]
  }
}

resource "aws_api_gateway_base_path_mapping" "this" {
  api_id      = var.rest_api_id
  stage_name  = var.stage_name
  domain_name = aws_api_gateway_domain_name.this.domain_name
}
