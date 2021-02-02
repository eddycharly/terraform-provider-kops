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
