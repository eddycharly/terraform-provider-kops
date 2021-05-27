# kops_kube_config

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
  cluster_ca_certificate = data.kops_kube_config.kube_config.ca_cert
}
```

## Argument Reference

The following arguments are supported:
- `cluster_name` - (Required) - String - The cluster name.
- `admin` - (Optional) - (Computed) - Int - Admin is the cluster admin user credential lifetime.
- `internal` - (Optional) - (Computed) - Bool - Internal use the cluster's internal DNS name.
- `server` - (Computed) - String - Kubernetes server url.
- `context` - (Computed) - String - Kubernetes context.
- `namespace` - (Computed) - String - Kubernetes namespace.
- `kube_user` - (Computed) - String - Kubernetes user.
- `kube_password` - (Sensitive) - (Computed) - String - Kubernetes password.
- `ca_cert` - (Sensitive) - (Computed) - String - Kubernetes cluster certificate.
- `client_cert` - (Sensitive) - (Computed) - String - Kubernetes client certificate.
- `client_key` - (Sensitive) - (Computed) - String - Kubernetes client key.




