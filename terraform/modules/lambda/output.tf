output "invoke_arn" {
  value = aws_lambda_function.this.arn
}

output "function_name" {
  value = aws_lambda_function.this.function_name
}

output "invocation_arn" {
  value = aws_lambda_function.this.invoke_arn
}
