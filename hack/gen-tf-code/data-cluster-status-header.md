Provides a kOps cluster status data source.

## Example usage

```hcl
data "kops_cluster_status" "status" {
  cluster_name = "cluster.example.com"
}
```
