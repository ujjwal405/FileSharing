output "s3_bucket_arn" {
  description = "The ARN of the S3 bucket"
  value       = aws_s3_bucket.this.arn
  sensitive   = true
}

output "website_endpoint" {
  value = aws_s3_bucket.this.website_endpoint
}
