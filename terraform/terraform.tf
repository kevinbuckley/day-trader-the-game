terraform {
  required_version = "~> 1.7"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.7.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.6.0"
    }
  }
  // cloud would go here
}

provider "aws" {
  region = var.region
}
