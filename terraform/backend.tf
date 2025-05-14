terraform {
  backend "s3" {
    bucket       = "filesharing-terraform-backend"
    key          = "prod/terraform.tfstate"
    region       = "ap-south-1"
    encrypt      = true
    use_lockfile = true #S3 native locking
  }
}
