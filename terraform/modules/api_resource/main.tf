
resource "aws_api_gateway_resource" "this" {
  # for_each    = { for idx, res in var.resources : idx => res }
  rest_api_id = var.rest_api_id
  parent_id   = var.parent_id
  path_part   = var.path_part
}


