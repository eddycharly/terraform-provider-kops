# kops_kube_config

KubeConfig represents the kubernets configuration needed to access a cluster.

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

