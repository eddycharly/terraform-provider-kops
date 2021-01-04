terraform {
  required_providers {
    kops = {
      source  = "github/eddycharly/kops"
      version = "0.0.1"
    }
  }
}

provider "kops" {
  state_store = "s3://cluster.example.com"

  aws {
    assume_role {
      role_arn = "arn:aws:iam::12345:role/admin"
    }
  }
}
