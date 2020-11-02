# terraform-provider-kops

`terraform-provider-kops` brings [KOPS](https://github.com/kubernetes/kops)
into terraform in a fully managed way, enabling idempotency through direct
integration with the KOPS api:
- No `local_exec`
- No yaml templating,
- No CLI invocations

... just **pure go code**.

Currently using KOPS `v1.18.2` and compatible with terraform `0.12` and higher.

## How does it work

The provider declares a single `kops_cluster` resource that holds the whole
cluster state and instance groups.

It is implemented this way because KOPS needs the full cluster state to apply
changes.

## Building the provider

To build the provider, clone this repository and run the following command:

```shell
make all
```

If you want to install the built provider after building it, run the following
command instead:

```shell
make install
```

## Using the provider

To use the provider you will need to register it in your terraform code:

```hcl
provider "kops" {
  state_store = "s3://cluster.example.com"
  // optionally use an AWS profile
  aws_profile = "example_profile"
}
```

## Using the kops_cluster resource

```hcl
resource "kops_cluster" "cluster" {
  name                 = "cluster.example.com"
  cloud_provider       = "aws"
  kubernetes_version   = "stable"
  dns_zone             = "example.com"

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
    cidr        = "10.0.64.0/19"
    type        = "Private"
    zone        = "eu-west-3a"
  }

  subnet {
    name        = "private-1"
    cidr        = "10.0.96.0/19"
    type        = "Private"
    zone        = "eu-west-3b"
  }

  subnet {
    name        = "private-2"
    cidr        = "10.0.128.0/19"
    type        = "Private"
    zone        = "eu-west-3c"
  }

  # master instance groups
  instance_group {
    name                     = "master-0"
    role                     = "Master"
    min_size                 = 1
    max_size                 = 1
    machine_type             = "t3.medium"
    subnets                  = ["private-0"]
  }

  instance_group {
    name                     = "master-1"
    role                     = "Master"
    min_size                 = 1
    max_size                 = 1
    machine_type             = "t3.medium"
    subnets                  = ["private-1"]
  }

  instance_group {
    name                     = "master-2"
    role                     = "Master"
    min_size                 = 1
    max_size                 = 1
    machine_type             = "t3.medium"
    subnets                  = ["private-2"]
  }

  # etcd clusters
  etcd_cluster {
    name            = "main"

    members {
      name             = "master-0"
      instance_group   = "master-0"
    }

    members {
      name             = "master-1"
      instance_group   = "master-1"
    }

    members {
      name             = "master-2"
      instance_group   = "master-2"
    }
  }

  etcd_cluster {
    name            = "events"

    members {
      name             = "master-0"
      instance_group   = "master-0"
    }

    members {
      name             = "master-1"
      instance_group   = "master-1"
    }

    members {
      name             = "master-2"
      instance_group   = "master-2"
    }
  }
}
```

## Getting kubeconfig file

To retrieve the `kubeconfig` file for the cluster, run the following command:

```shell
kops export kubecfg --name cluster.example.com --state s3://cluster.example.com
```
