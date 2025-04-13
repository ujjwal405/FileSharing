output "permission_id" {
  description = "ID of the created Lambda permission"
  value       = aws_lambda_permission.this.id
}
