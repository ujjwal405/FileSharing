data "aws_acm_certificate" "cert" {
  domain      = var.domain_name # Replace with your domain
  statuses    = ["ISSUED"]
  most_recent = true
}

resource "aws_api_gateway_domain_name" "custom_domain" {
  domain_name              = var.domain_name                   # Replace with your custom domain
  regional_certificate_arn = data.aws_acm_certificate.cert.arn # Replace with ACM certificate ARN
  endpoint_configuration {
    types = ["REGIONAL"]
  }
}

resource "aws_api_gateway_base_path_mapping" "custom_domain_mapping" {
  domain_name = aws_api_gateway_domain_name.custom_domain.id
  api_id      = var.rest_api_id
  stage_name  = var.stage_name # Replace with your API stage name
}
