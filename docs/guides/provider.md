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



## Nullable arguments

Because kOps sometimes uses pointers to hold data and terraform doesn't offer a way to
differentiate between unset arguments and their default value in a configuration, it can
be necessary wrap those arguments in a nested resource to account for the `null` value.

An example of this is the `anonymous_auth` argument in the `kube_api_server_config` or `kubelet_config_spec`
resources. The `null` value cannot be considered equivalent to `false` on the kOps side, but terraform won't
let us know when it is set or not in the configuration and therefore will provide `false` in case it is unset
in the configuration.

To workaround this limitation, the nullable type is a resource with a single `value` argument. It wraps the
actual value of the argument and makes it possible to account for `null`. When using a nullable argument,
you should assign it this way in the configuration:

```hcl
resource "kops_cluster" "cluster" {
  // ...

  kubelet {
    anonymous_auth {
      value = false
    }
  }

  kube_api_server {
    anonymous_auth {
      value = false
    }
  }

  // ...
}
```


## Argument Reference

The following arguments are supported:
- `state_store` - (Required) - String - StateStore defines the state store used by kops.
- `aws` - (Optional) - [aws](#aws) - Aws contains the aws configuration options.
- `openstack` - (Optional) - [openstack](#openstack) - OpenStack contains the openstack configuration options.
- `klog` - (Optional) - [klog](#klog) - Klog contains the klog configuration options.

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

### openstack

#### Argument Reference

The following arguments are supported:

- `tenant_id` - (Optional) - String
- `tenant_name` - (Optional) - String
- `project_id` - (Optional) - String
- `project_name` - (Optional) - String
- `project_domain_id` - (Optional) - String
- `project_domain_name` - (Optional) - String
- `domain_id` - (Optional) - String
- `domain_name` - (Optional) - String
- `username` - (Optional) - String
- `password` - (Optional) - String
- `auth_url` - (Optional) - String
- `region_name` - (Optional) - String
- `application_credential_id` - (Optional) - String
- `application_credential_secret` - (Optional) - String

### klog

#### Argument Reference

The following arguments are supported:

- `verbosity` - (Optional) - Int([Nullable](#nullable-arguments)) - Verbosity defines the verbosity of klog.



