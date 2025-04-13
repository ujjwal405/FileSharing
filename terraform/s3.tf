// creating s3 bucket for storing user files
module "s3_file_upload" {
  source           = "./modules/s3"
  bucket_name      = var.s3_file_upload
  enable_sse       = true
  enable_lifecycle = true
}

// creating s3 bucket for hosting static website
module "s3_static_website" {
  source                  = "./modules/s3"
  bucket_name             = var.s3_static_website
  force_destroy           = true
  block_public_acls       = false
  block_public_policy     = false
  ignore_public_acls      = false
  restrict_public_buckets = false
  enable_static_website   = true
  enable_public_access    = true

}
