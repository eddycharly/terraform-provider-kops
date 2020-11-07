# terraform-provider-kops

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/eddycharly/terraform-provider-kops/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/eddycharly/terraform-provider-kops)](https://goreportcard.com/report/eddycharly/terraform-provider-kops)

`terraform-provider-kops` brings [KOPS](https://github.com/kubernetes/kops)
into terraform in a fully managed way, enabling idempotency through direct
integration with the KOPS api:
- No `local_exec`
- No yaml templating,
- No CLI invocations

... just **pure go code**.

Currently using KOPS `v1.18.2` and compatible with terraform `0.12` and higher.

**NOTE**: For now, provisioning the network is not supported. The network must
be created separately and given to the provider through cluster attribute
`network_id` and subnets attributes `provider_id`.

## How does it work

The provider declares a single `kops_cluster` resource that holds the whole
cluster state and instance groups.

It is implemented this way because KOPS needs the full cluster state to apply
changes.

## Why use it

KOPS is an amazing tool but it can be challenging to integrate in an IAC
(infrastructure as code) stack.

Typical solutions usually involve running KOPS CLI in shell scripts or generating
KOPS templates manually and force syncing them with the KOPS store.

In most cases, getting something idempotent is difficult because you need to
somewhat keep the state of the cluster and are responsible for deleting obsolete
instange groups for example.

This is where `terraform` shines in, state management. This provider takes care
of creating, updating and deleting instance groups as they evolve over time.

Even if KOPS provides `kops update cluster --target terraform` to create the
terraform configuration for a KOPS cluster, it is still necessary to run
`kops rolling-update cluster` to recycle instance groups when the change.
With this provider, this is all taken care of and you should never need to run
KOPS manually.

## Docs

The full documentation is available in the [docs](./docs/README.md) folder.

## Installing the provider

To install the provider, download the release for your distribution from the
[releases page](https://github.com/eddycharly/terraform-provider-kops/releases)
and refer to the terraform instructions on 
[plugins install](https://www.terraform.io/docs/plugins/basics.html#installing-plugins).

The provider should be published on the [terraform registry](https://registry.terraform.io/)
soon.
I'm just waiting to reach at least the beta stage before getting it advertised
and more generally available.

## Building the provider

To build the provider, clone this repository and run the following command:

```shell
make all
```

If you want to install the built provider after building it, run the following
command instead (working on linux and macos):

```shell
make install
```

## Using the provider

To use the provider you will need to register it in your terraform code:

```hcl
terraform {
  required_providers {
    kops = {
      source   = "github/eddycharly/kops"
      versions = ["0.0.1"]
    }
  }
}

provider "kops" {
  state_store = "s3://cluster.example.com"
  // optionally set up your cloud provider access config
  aws {
    profile = "example_profile"
  }
}
```

## Using the kops_cluster resource

```hcl
resource "kops_cluster" "cluster" {
  name                 = "cluster.example.com"
  admin_ssh_key        = file("path to ssh public key file")
  cloud_provider       = "aws"
  kubernetes_version   = "stable"
  dns_zone             = "example.com"
  network_id           = "net-0"

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

More examples are available in the `/examples` folder:
- [Basic example](./examples/basic)
- [Aws profile example](./examples/aws-profile)
- [Bastion example](./examples/bastion)

## Importing an existing cluster

You can import an existing cluster by creating a `kops_resource` configuration
and running the `terraform import` command:

1. Create a terraform configuration:

    ```hcl
    provider "kops" {
      state_store = "s3://cluster.example.com"
    }

    resource "kops_cluster" "cluster" {
      name        = "cluster.example.com"
      
      // ....
    }
    ```

1. Run `terraform import`:

    ```shell
    terraform import kops_cluster.cluster cluster.example.com
    ```

## Getting kubeconfig file

To retrieve the `kubeconfig` file for the cluster, run the following command:

```shell
kops export kubecfg --name cluster.example.com --state s3://cluster.example.com
```
