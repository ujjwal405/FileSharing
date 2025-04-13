resource "aws_iam_policy" "this" {
  name = "${var.policy_name}-${var.env_name}-policy"


  policy = jsonencode({
    Version   = "2012-10-17"
    Statement = var.policy_statements # Dynamically inject statements
  })
}
