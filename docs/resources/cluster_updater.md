# kops_cluster_updater

This resource applies the cluster state on the target cloud provider and performs a rolling update.

The rolling update and cluster validation can be disabled and/or configured through resource attributes.

~> This resource will trigger based on the `keepers` map attribute.
Thats is, if something changes in the attribute, the resource update handler will fire and an apply/rolling update/validate cycle will run.
A good candidate for `keepers` is to use the `revision` coming from `kops_cluster` and `kops_instance_group` resources.

## Example usage

```hcl
resource "kops_cluster" "cluster" {
  name                 = "cluster.example.com"

  // ....
}

resource "kops_instance_group" "master-0" {
  cluster_name = kops_cluster.cluster.name
  name         = "master-0"
  role         = "Master"
  min_size     = 1
  max_size     = 1
  machine_type = "t3.medium"
  subnets      = ["private-0"]
  depends_on   = [kops_cluster.cluster]
}

resource "kops_instance_group" "master-1" {
  cluster_name = kops_cluster.cluster.name
  name         = "master-1"
  role         = "Master"
  min_size     = 1
  max_size     = 1
  machine_type = "t3.medium"
  subnets      = ["private-1"]
  depends_on   = [kops_cluster.cluster]
}

resource "kops_instance_group" "master-2" {
  cluster_name = kops_cluster.cluster.name
  name         = "master-2"
  role         = "Master"
  min_size     = 1
  max_size     = 1
  machine_type = "t3.medium"
  subnets      = ["private-2"]
  depends_on   = [kops_cluster.cluster]
}

resource "kops_cluster_updater" "updater" {
  cluster_name        = kops_cluster.cluster.name

  keepers = {
    cluster  = kops_cluster.cluster.revision,
    master-0 = kops_instance_group.master-0.revision,
    master-1 = kops_instance_group.master-1.revision,
    master-2 = kops_instance_group.master-2.revision
    // ...
  }

  rolling_update {
    skip                = false
    fail_on_drain_error = true
    fail_on_validate    = true
    validate_count      = 1

    // ...
  }

  validate {
    skip = false

    // ...
  }

  # ensures rolling update happens after the cluster and instance groups are up to date
  depends_on   = [
    kops_cluster.cluster,
    kops_instance_group.master-0,
    kops_instance_group.master-1,
    kops_instance_group.master-2
  ]
}
```

## Argument Reference

The following arguments are supported:
- `revision` - (Computed) - Int - Revision is incremented every time the resource changes, this is useful for triggering cluster updater.
- `cluster_name` - (Required) - String - ClusterName is the target cluster name.
- `keepers` - (Optional) - Map(String) - Keepers contains arbitrary strings used to update the resource when one changes.
- `apply` - (Optional) - [apply_options](#apply_options) - Apply holds cluster apply options.
- `rolling_update` - (Optional) - [rolling_update_options](#rolling_update_options) - RollingUpdate holds cluster rolling update options.
- `validate` - (Optional) - [validate_options](#validate_options) - Validate holds cluster validation options.

## Nested resources

### apply_options

#### Argument Reference

The following arguments are supported:

- `skip` - (Optional) - Bool - Skip allows skipping cluster apply.

### rolling_update_options

#### Argument Reference

The following arguments are supported:

- `skip` - (Optional) - Bool - Skip allows skipping cluster rolling update.
- `master_interval` - (Optional) - Duration - MasterInterval is the amount of time to wait after stopping a master instance.
- `node_interval` - (Optional) - Duration - NodeInterval is the amount of time to wait after stopping a non-master instance.
- `bastion_interval` - (Optional) - Duration - BastionInterval is the amount of time to wait after stopping a bastion instance.
- `fail_on_drain_error` - (Optional) - Bool - FailOnDrainError will fail when a drain error occurs.
- `fail_on_validate` - (Optional) - Bool - FailOnValidate will fail when a validation error occurs.
- `post_drain_delay` - (Optional) - Duration - PostDrainDelay is the duration we wait after draining each node.
- `validation_timeout` - (Optional) - Duration - ValidationTimeout is the maximum time to wait for the cluster to validate, once we start validation.
- `validate_count` - (Optional) - Int - ValidateCount is the amount of time that a cluster needs to be validated after single node update.

### validate_options

#### Argument Reference

The following arguments are supported:

- `skip` - (Optional) - Bool - Skip allows skipping cluster validation.
- `timeout` - (Optional) - Duration - Timeout defines the maximum time to wait until the cluster becomes valid.
- `poll_interval` - (Optional) - Duration - PollInterval defines the interval between validation attempts.



