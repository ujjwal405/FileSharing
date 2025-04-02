output "secret_ids" {
  description = "A map of secret names to their IDs"
  value       = { for k, v in aws_secretsmanager_secret.this : k => v.id }
  sensitive   = true
}
