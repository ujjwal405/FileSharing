terraform {
  backend "s3" {
    bucket       = "your-terraform-state-bucket"
    key          = "filesharing/terraform.tfstate"
    region       = "ap-south-1"
    encrypt      = true
    use_lockfile = true #S3 native locking
  }
}
