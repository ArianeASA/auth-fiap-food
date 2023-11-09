terraform {

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.23.1"
    }
  }

  required_version = "~> 1.2"

  backend "s3" {
    bucket = "terraforms-bucket"
    key    = "states-lambda"
    region = var.aws_region
    assume_role = {
      role_arn = "arn:aws:iam::${var.account_id}:role/Terraform"
    }
  }
}

data "terraform_remote_state" "network" {
  backend = "s3"
  config = {
    bucket = "terraforms-bucket"
    key    = "states-lambda/terraform.tfstate"
    region = var.aws_region
  }
}
