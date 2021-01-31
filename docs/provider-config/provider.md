# kops_provider


## Argument Reference

The following arguments are supported:
- `state_store` - (Required) - String - StateStore defines the state store used by kops.
- `aws` - (Optional) - [aws](#aws) - Aws contains the aws configuration options.

## Nested resources

### aws

#### Argument Reference

The following arguments are supported:

- `profile` - (Optional) - String
- `region` - (Optional) - String
- `assume_role` - (Optional) - [aws_assume_role](#aws_assume_role)

### aws_assume_role

#### Argument Reference

The following arguments are supported:

- `role_arn` - (Optional) - String



