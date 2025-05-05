terraform {
  required_version = ">= 1.9.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 1.11"
    }
  }
}

provider "aws" {
  region = var.aws_region
}
