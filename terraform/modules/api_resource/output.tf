# output "resource_ids" {
#   value = { for key, res in aws_api_gateway_resource.this : res.path => res.id }
# }

output "resource_id" {
  value = aws_api_gateway_resource.this.id
}
