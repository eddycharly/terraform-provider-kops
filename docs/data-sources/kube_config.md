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
- `server` - (Computed) - String - Kubernetes server url.
- `context` - (Computed) - String - Kubernetes context.
- `namespace` - (Computed) - String - Kubernetes namespace.
- `kube_bearer_token` - (Computed) - String - Kubernetes bearer token.
- `kube_user` - (Computed) - String - Kubernetes user.
- `kube_password` - (Computed) - String - Kubernetes password.
- `ca_cert` - (Computed) - String - Kubernetes cluster certificate.
- `client_cert` - (Computed) - String - Kubernetes client certificate.
- `client_key` - (Computed) - String - Kubernetes client key.

## Nested resources



