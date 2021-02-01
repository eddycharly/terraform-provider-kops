The provider configuration contains kOps state store and authentication settings.

## Example usage

### State store configuration

```hcl
provider "kops" {
  state_store = "s3://cluster.example.com"
}
```

### Authentication using an AWS profile

```hcl
provider "kops" {
  state_store = "s3://cluster.example.com"

  // optionally set up your cloud provider access config
  aws {
    profile = "example_profile"
  }
}
```

### Authentication using an AWS assumed role

```hcl
provider "kops" {
  state_store = "s3://cluster.example.com"

  // optionally set up your cloud provider access config
  aws {
    region  = "eu-west-1"
    assume_role {
      role_arn = "arn:aws:iam::0123456789:role/admin"
    }
  }
}
```

