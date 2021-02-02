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
