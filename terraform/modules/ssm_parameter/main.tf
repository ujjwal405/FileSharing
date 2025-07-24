resource "aws_ssm_parameter" "this" {
  name        = var.parameter_name
  type        = var.parameter_type
  value       = var.parameter_value

}
