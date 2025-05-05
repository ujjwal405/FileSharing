terraform {
  backend "s3" {
    bucket       = "filesharing-terraform-backend"
    key          = "filesharing-backend/terraform.tfstate"
    region       = "ap-south-1"
    encrypt      = true
    use_lockfile = true #S3 native locking
  }
}
