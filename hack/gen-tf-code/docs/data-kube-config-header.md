Provides a kOps kube config data source.

## Example usage

```hcl
data "kops_kube_config" "kube_config" {
  cluster_name = "cluster.example.com"
}

# then you can configure another provider based on the kube_config data source

provider "kubectl" {
  host                   = data.kops_kube_config.kube_config.server
  username               = data.kops_kube_config.kube_config.kube_user
  password               = data.kops_kube_config.kube_config.kube_password
  client_certificate     = data.kops_kube_config.kube_config.client_cert
  client_key             = data.kops_kube_config.kube_config.client_key
  cluster_ca_certificate = data.kops_kube_config.kube_config.ca_certs
}
```
