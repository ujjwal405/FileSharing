resource "aws_s3_bucket" "this" {
  bucket        = var.bucket_name
  force_destroy = var.force_destroy
}

resource "aws_s3_bucket_server_side_encryption_configuration" "this" {
  count  = var.enable_sse ? 1 : 0
  bucket = aws_s3_bucket.this.bucket

  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}


resource "aws_s3_bucket_ownership_controls" "this" {
  bucket = aws_s3_bucket.this.id
  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}



resource "aws_s3_bucket_public_access_block" "this" {
  bucket                  = aws_s3_bucket.this.id
  block_public_acls       = var.block_public_acls
  block_public_policy     = var.block_public_policy
  ignore_public_acls      = var.ignore_public_acls
  restrict_public_buckets = var.restrict_public_buckets
}

# resource "aws_s3_bucket_acl" "this" {
#   count  = var.enable_public_access ? 1 : 0
#   bucket = aws_s3_bucket.this.id
#   acl    = "public-read"

#   depends_on = [
#     aws_s3_bucket_ownership_controls.this,
#     aws_s3_bucket_public_access_block.this,
#   ]
# }



# resource "aws_s3_bucket_website_configuration" "this" {
#   count  = var.enable_static_website ? 1 : 0
#   bucket = aws_s3_bucket.this.id

#   dynamic "index_document" {
#     for_each = var.index_document != "" ? [var.index_document] : []
#     content {
#       suffix = var.index_document
#     }
#   }

#   dynamic "error_document" {
#     for_each = var.error_document != "" ? [var.error_document] : []
#     content {
#       key = var.error_document
#     }
#   }
# }


# resource "aws_s3_bucket_policy" "public_read_access" {
#   count  = var.enable_public_access ? 1 : 0
#   bucket = aws_s3_bucket.this.id
#   policy = jsonencode({
#     Version = "2012-10-17"
#     Statement = [
#       {
#         Effect    = "Allow"
#         Principal = "*"
#         Action    = "s3:GetObject"
#         Resource = [
#           "${aws_s3_bucket.this.arn}/*"
#         ]
#       }
#     ]
#   })
# }





resource "aws_s3_bucket_lifecycle_configuration" "lifecycle_rules" {
  count  = var.enable_lifecycle ? 1 : 0
  bucket = aws_s3_bucket.this.id

  rule {
    id     = "expire-old-objects"
    status = "Enabled"
    filter {
      # empty prefix = apply to all objects
      prefix = ""

    }

    transition {
      days          = 30
      storage_class = "STANDARD_IA"
    }

    transition {
      days          = 90
      storage_class = "GLACIER"
    }

    expiration {
      days = 365
    }
  }
}
