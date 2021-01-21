# kops_cluster_status

Provides a kOps cluster status data source.

## Example usage

```hcl
data "kops_cluster_status" "status" {
  cluster_name = "cluster.example.com"
}
```

## Argument Reference

The following arguments are supported:
- `cluster_name` - (Required) - String - The target cluster name.
- `exists` - (Computed) - Bool - Tells if the cluster exists.
- `is_valid` - (Computed) - Bool - Tells if the cluster is valid.
- `needs_update` - (Computed) - Bool - Tells if the cluster needs a rolling update.




