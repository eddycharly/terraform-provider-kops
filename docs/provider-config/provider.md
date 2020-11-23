# kops_provider


## Argument Reference

The following arguments are supported:
- `state_store` - (Required) - String - StateStore defines the state store used by kops.
- `aws` - (Optional) - [aws](#aws) - Aws contains the aws configuration options.
- `rolling_update` - (Optional) - [rolling_update](#rolling_update) - RollingUpdateOptions contains the options used when doing a cluster rolling update.
- `validate` - (Optional) - [validate](#validate) - ValidateOptions contains the options used when validating the cluster.

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

### rolling_update

#### Argument Reference

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

### validate

#### Argument Reference

The following arguments are supported:

- `skip` - (Optional) - Bool
- `timeout` - (Optional) - Duration
- `poll_interval` - (Optional) - Duration

