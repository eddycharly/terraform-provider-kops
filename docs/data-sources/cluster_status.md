# kops_cluster_status

Provides a kOps cluster status data source.

## Example usage

```hcl
data "kops_cluster_status" "status" {
  cluster_name = "cluster.example.com"
  apply        = false
}

output "needs_update {
  value = kops_cluster_status.status.needs_update
}
```

## Argument Reference

The following arguments are supported:
- `cluster_name` - (Required) - String - ClusterName defines the target cluster name.
- `apply` - (Required) - Bool - Apply updates the cluster before computing its status.
- `exists` - (Computed) - Bool - Exists indicates if the cluster exists.
- `is_valid` - (Computed) - Bool - IsValid indicates if the cluster is valid.
- `needs_update` - (Computed) - Bool - NeedsUpdate indicates if the cluster needs a rolling update.
- `instance_groups` - (Computed) - List(String) - InstanceGroups contains the name of instance groups to be updated.




