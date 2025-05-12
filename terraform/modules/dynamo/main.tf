provider "aws" {
  region = var.aws_region
}


resource "aws_dynamodb_table" "my_table" {
  name           = var.table_name
  hash_key       = var.hash_key
  billing_mode   = "PAY_PER_REQUEST"
  read_capacity  = var.read_capacity
  write_capacity = var.write_capacity
  range_key      = var.range_key != null ? var.range_key : null

  dynamic "attribute" {
    for_each = var.attributes
    content {
      name = attribute.value.name
      type = attribute.value.type
    }
  }
  lifecycle {
    prevent_destroy = false
  }
}
