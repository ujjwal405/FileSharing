output "user_pool_id" {
  value       = aws_cognito_user_pool.myPool.id
  description = "The ID of the Cognito App Client"
}


output "app_client_id" {
  value       = aws_cognito_user_pool_client.myClient.id
  description = "The ID of the Cognito App Client"
  sensitive   = true

}

output "user_pool_arn" {
  value     = aws_cognito_user_pool.myPool.arn
  sensitive = true
}
