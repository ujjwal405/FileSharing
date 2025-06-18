output "stage_invoke_url" {
  value = aws_api_gateway_stage.this.invoke_url
}

output "stage_name" {
  value = aws_api_gateway_stage.this.stage_name
}
