output "methods_signature" {
  value = sha1(jsonencode([
    for k, method in var.methods : {
      resource_name  = method.resource_name
      http_method    = method.http_method
      use_authorizer = method.use_authorizer
      lambda_arn     = method.lambda_arn
    }
  ]))
}
