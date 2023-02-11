terraform {
    required_providers {
      aws = {
        source = "hashicorp/aws"
        version = "~> 4.0"
      }
    }

    backend "s3" {
        bucket = "angel-owl-tfstate"
        key = "angel-owl.tfstate"
        region = "ap-southeast-1"

        dynamodb_table = "angel-owl-tflock"
        
        profile = "cs301project"
    }
}

provider "aws" {
    region = "ap-southeast-1"
    profile = "cs301project"
}
