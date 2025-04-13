resource "aws_iam_role" "this" {
  name = "${var.lambda_name}-${var.env_name}"

  assume_role_policy = jsonencode({
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "this" {
  for_each   = toset(var.lambda_policy_arns)
  role       = aws_iam_role.this.name
  policy_arn = each.value
}


