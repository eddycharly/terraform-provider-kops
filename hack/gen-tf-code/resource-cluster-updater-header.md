Performs a cluster rolling update (if needed).

~> This resource is designed to permanently recreate.
The cluster will be rolled out only when necessary.

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
  fail_on_drain_error = true
  fail_on_validate    = true
  validate_count      = 1

  // ...

  # ensures rolling update happens after the cluster and instance groups are up to date
  depends_on   = [
    kops_cluster.cluster,
    kops_instance_group.master-0,
    kops_instance_group.master-1,
    kops_instance_group.master-2
  ]
}
```