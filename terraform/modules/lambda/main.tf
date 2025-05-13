locals {
  source_dir = "${path.module}/../../../backend/lambda/${var.lambda_name}/bootstrap"
  #   output_path = "${path.module}/../../../backend/lambda/${var.lambda_name}"
  output_path = "${path.module}/../../../backend/lambda/${var.lambda_name}"

}


resource "aws_lambda_function" "this" {
  function_name = "${var.lambda_name}-${var.env_name}"
  handler       = "bootstrap" # Go functions use "bootstrap" as the handler
  runtime       = var.function_runtime
  timeout       = var.function_timeout


  filename         = "${local.output_path}/${var.lambda_name}.zip"
  source_code_hash = data.archive_file.this.output_base64sha256


  # filename         = "${local.output_path}/bootstrap/bootstrap.zip"
  # source_code_hash = data.local_file.lambda_zip.content_base64sha256

  role = var.lambda_role_arn

  dynamic "environment" {
    for_each = length(var.environment_variables) > 0 ? [1] : []
    content {
      variables = var.environment_variables
    }
  }


}


data "archive_file" "this" {
  source_dir  = local.source_dir
  type        = "zip"
  output_path = "${local.output_path}/${var.lambda_name}.zip"
}

# data "local_file" "lambda_zip" {
#   filename = "${local.output_path}/bootstrap/bootstrap.zip"
# }
