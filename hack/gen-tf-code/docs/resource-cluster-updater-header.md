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
    cluster  = kops_cluster.cluster.revision
    master-0 = kops_instance_group.master-0.revision
    master-1 = kops_instance_group.master-1.revision
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