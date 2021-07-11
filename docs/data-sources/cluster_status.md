# kops_cluster_status

Provides a kOps cluster status data source.

This data source assumes the cluster has been applied (or updated) see [kops_cluster_updater](/docs/resources/cluster_updater).

## Example usage

```hcl
resource "kops_cluster_updater" "updater" {
  cluster_name = "cluster.example.com"

  // ...
}

data "kops_cluster_status" "status" {
  cluster_name = "cluster.example.com"
  depends_on   = [kops_cluster_updater.updater]
}

output "needs_update {
  value = kops_cluster_status.status.needs_update
}
```


## Argument Reference

The following arguments are supported:
- `cluster_name` - (Required) - String - ClusterName defines the target cluster name.
- `exists` - (Computed) - Bool - Exists indicates if the cluster exists.
- `is_valid` - (Computed) - Bool - IsValid indicates if the cluster is valid.
- `needs_update` - (Computed) - Bool - NeedsUpdate indicates if the cluster needs a rolling update.
- `instance_groups` - (Computed) - List(String) - InstanceGroups contains the name of instance groups to be updated.




