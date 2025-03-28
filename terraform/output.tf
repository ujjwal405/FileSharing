resource "null_resource" "get_lambda_dirs" {
     triggers = {
    always_run = timestamp()
  }
  provisioner "local-exec" {
    command     = "./get_lambdas.sh"
  }
}

data "local_file" "lambda_dirs" {
  filename   = "${path.module}/lambda_dirs.json"
  depends_on = [null_resource.get_lambda_dirs]
}

locals {
  lambda_dirs = jsondecode(data.local_file.lambda_dirs.content)
}

output "lambda_dirs" {
  value = local.lambda_dirs
}