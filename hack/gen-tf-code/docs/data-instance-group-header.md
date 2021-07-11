Provides a kOps cluster instance group data source.

## Example usage

```hcl
data "kops_instance_group" "ig-0" {
  cluster_name = "cluster.example.com"
  name         = "ig-0"
}
```
