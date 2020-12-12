# Provider configuration

The kops provider supports configuring the following elements:
- `state_store` - (Required) - String - StateStore defines the state store used by kops.
- `aws` - (Optional) - [aws](#aws) - Aws contains the aws configuration options.
- `rolling_update` - (Optional) - [rolling_update](#rolling_update) - RollingUpdateOptions contains the options used when doing a cluster rolling update.
- `validate` - (Optional) - [validate](#validate) - ValidateOptions contains the options used when validating the cluster.

## Configure state store

```hcl
provider "kops" {
  state_store = "s3://cluster.example.com"
  // optionally set up your cloud provider access config
  aws {
    profile = "example_profile"
  }
}
```

The following arguments are supported:
- `state_store` - (Required) - String - StateStore defines the state store used by kops.

## Configure AWS credentials

```hcl
provider "kops" {
  aws {
    profile = "example_profile"
    region  = "eu-west-1"
    assume_role {
      role_arn = "arn:aws:iam::0123456789:role/admin"
    }
  }
}
```

### aws

The following arguments are supported:

- `profile` - (Optional) - String
- `region` - (Optional) - String
- `assume_role` - (Optional) - [aws_assume_role](#aws_assume_role)

### aws_assume_role

The following arguments are supported:

- `role_arn` - (Optional) - String

## Configure cluster rolling update options

```hcl
provider "kops" {
  rolling_update {
    skip = true
    // ...
  }
}
```

### rolling_update

The following arguments are supported:

- `skip` - (Optional) - Bool
- `master_interval` - (Optional) - Duration
- `node_interval` - (Optional) - Duration
- `bastion_interval` - (Optional) - Duration
- `fail_on_drain_error` - (Optional) - Bool
- `fail_on_validate` - (Optional) - Bool
- `post_drain_delay` - (Optional) - Duration
- `validation_timeout` - (Optional) - Duration
- `validate_count` - (Optional) - Int

## Configure cluster validation options

```hcl
provider "kops" {
  validate {
    skip = true
    // ...
  }
}
```

### validate

The following arguments are supported:

- `skip` - (Optional) - Bool
- `timeout` - (Optional) - Duration
- `poll_interval` - (Optional) - Duration
