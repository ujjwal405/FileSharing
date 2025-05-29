resource "aws_api_gateway_domain_name" "custom_domain" {
  domain_name              = "api.example.com"                                                         # Replace with your custom domain
  regional_certificate_arn = "arn:aws:acm:us-east-1:123456789012:certificate/your-acm-certificate-arn" # Replace with ACM certificate ARN
}

resource "aws_api_gateway_base_path_mapping" "custom_domain_mapping" {
  domain_name = aws_api_gateway_domain_name.custom_domain.id
  rest_api_id = var.rest_api_id
  stage_name  = var.stage_name # Replace with your API stage name
}
