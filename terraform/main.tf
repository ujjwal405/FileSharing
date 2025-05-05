
data "aws_caller_identity" "current" {}

resource "null_resource" "get_lambda_dirs" {
  triggers = {
    always_run = timestamp()
  }
  provisioner "local-exec" {
    command = "./get_lambdas.sh"
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
      for dir in ${join(" ", local.lambda_dirs)}; do  # Iterate over each directory in the list
        echo "Processing Lambda in directory: $dir"
        cd "$(pwd)/../backend/lambda/$dir/cmd"  # Change into the Lambda function directory
        GOOS=linux GOARCH=amd64 go build -o bootstrap main.go  # Build the Lambda function
        mkdir -p "$(pwd)/../../$dir/bootstrap" #Create the directory.
        mv bootstrap.zip "$(pwd)/../../$dir/bootstrap/bootstrap.zip"
      done
    EOT
  }
}




