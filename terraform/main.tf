
data "aws_caller_identity" "current" {}

resource "null_resource" "get_lambda_dirs" {
  triggers = {
    always_run = timestamp()
  }
  provisioner "local-exec" {
    command = "chmod +x ./get_lambdas.sh && ./get_lambdas.sh"
  }
}


data "local_file" "lambda_dirs" {
  filename   = "${path.module}/lambda_dirs.json"
  depends_on = [null_resource.get_lambda_dirs]
}



locals {
  lambda_dirs = jsondecode(data.local_file.lambda_dirs.content)
}



resource "null_resource" "build_lambdas" {
  triggers = {
    always_run = timestamp()
  }

  provisioner "local-exec" {
    command = <<EOT
      set -e
      for dir in ${join(" ", local.lambda_dirs)}; do
        LAMBDADIR=$(pwd)/../backend/lambda/$dir
        echo "🔨 Building $dir…"
        cd "$LAMBDADIR/cmd"
        GOOS=linux GOARCH=amd64 go build -o bootstrap main.go

        echo "🗜 Zipping $dir…"
        mkdir -p "$LAMBDADIR/bootstrap"
        mv bootstrap "$LAMBDADIR/bootstrap/"

        echo "✅ Moved to $dir/bootstrap/bootstrap"
      done
    EOT
  }
}




