Provides a kOps cluster resource.

## Example usage

```hcl
resource "kops_cluster" "cluster" {
  name                 = "cluster.example.com"
  admin_ssh_key        = file("path to ssh public key file")
  kubernetes_version   = "stable"
  dns_zone             = "example.com"
  network_id           = "net-0"

  cloud_provider {
    aws {}
  }

  iam {
    allow_container_registry = true
  }

  networking {
    calico {}
  }

  topology {
    masters = "private"
    nodes   = "private"

    dns {
      type = "Private"
    }
  }

  # cluster subnets
  subnet {
    name        = "private-0"
    provider_id = "subnet-0"
    type        = "Private"
    zone        = "zone-0"
  }

  subnet {
    name        = "private-1"
    provider_id = "subnet-1"
    type        = "Private"
    zone        = "zone-1"
  }

  subnet {
    name        = "private-2"
    provider_id = "subnet-2"
    type        = "Private"
    zone        = "zone-2"
  }

  # etcd clusters
  etcd_cluster {
    name            = "main"

    member {
      name             = "master-0"
      instance_group   = "master-0"
    }

    member {
      name             = "master-1"
      instance_group   = "master-1"
    }

    member {
      name             = "master-2"
      instance_group   = "master-2"
    }
  }

  etcd_cluster {
    name            = "events"

    member {
      name             = "master-0"
      instance_group   = "master-0"
    }

    member {
      name             = "master-1"
      instance_group   = "master-1"
    }

    member {
      name             = "master-2"
      instance_group   = "master-2"
    }
  }
}
```