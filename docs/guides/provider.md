# kops_provider

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

## Argument Reference

The following arguments are supported:
- `state_store` - (Required) - String - StateStore defines the state store used by kops.
- `aws` - (Optional) - [aws](#aws) - Aws contains the aws configuration options.

## Nested resources

### aws

#### Argument Reference

The following arguments are supported:

- `profile` - (Optional) - String - Profile defines the AWS profile to load when calling aws services.
- `region` - (Optional) - String - Region defines the AWS region.
- `access_key` - (Optional) - String - Region defines the AWS access key.
- `secret_key` - (Optional) - String - Region defines the AWS secret key.
- `assume_role` - (Optional) - [aws_assume_role](#aws_assume_role) - AssumeRole defines the AWS IAM role to be assumed.
- `s3_endpoint` - (Optional) - String - S3Endpoint defines S3 compatible endpoint.
- `s3_region` - (Optional) - String - S3Region defines S3 compatible endpoint region.
- `s3_access_key` - (Optional) - String - S3AccessKey defines S3 compatible endpoint access key.
- `s3_secret_key` - (Optional) - String - S3SecretKey defines S3 compatible endpoint secret key.
- `skip_region_check` - (Optional) - Bool - SkipRegionCheck skips validating region check.

### aws_assume_role

#### Argument Reference

The following arguments are supported:

- `role_arn` - (Optional) - String - RoleArn defines the arn of the AWS IAM role to assume.



