# terraform-provider-kops

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/eddycharly/terraform-provider-kops/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/eddycharly/terraform-provider-kops)](https://goreportcard.com/report/eddycharly/terraform-provider-kops)

`terraform-provider-kops` brings [kOps](https://github.com/kubernetes/kops)
into terraform in a fully managed way, enabling idempotency through direct
integration with the kOps api:
- No `local_exec`
- No yaml templating
- No CLI invocations

... just **pure go code.**

Currently using kOps `v1.19.2` and compatible with terraform `0.12` and higher.

**NOTES**
- For now, provisioning the network is not supported. The network must
be created separately and given to the provider through cluster attribute
`network_id` and subnets attributes `provider_id`.
- The provider has only been tested with AWS and calico networking.
If you use it with another cloud or networking provider, please let us know so
that we can help troubleshooting if necessary and update the docs.

## Why use it

kOps is an amazing tool but it can be challenging to integrate in an IAC
(infrastructure as code) stack.

Typical solutions usually involve running kOps CLI in shell scripts or generating
kOps templates manually and force syncing them with the kOps store.

In most cases, getting something idempotent is difficult because you need to
somewhat keep the state of the cluster and are responsible for deleting obsolete
instange groups for example.

This is where `terraform` shines in, state management. This provider takes care
of creating, updating and deleting instance groups as they evolve over time.

Even if kOps provides `kops update cluster --target terraform` to create the
terraform configuration for a kOps cluster, it is still necessary to run
`kops rolling-update cluster` to recycle instance groups when something changes
in the cluster.
With this provider, this is all taken care of and you should never need to invoke
kOps manually.

## How does it work

The provider declares resources to declare the state of the cluster:
- [kops_cluster](/docs/resources/cluster.md) defines the desired state of a cluster
- [kops_instance_group](/docs/resources/instance_group.md) defines the desired state of a cluster instance group

The provider also declares data sources to fetch the state of the cluster and
use it in your terraform code:
- [kops_cluster](/docs/data-sources/cluster.md) fetches the current state of a cluster
- [kops_instance_group](/docs/data-sources/instance_group.md) fetches the current state of a cluster instance group
- [kops_cluster_status](/docs/data-sources/cluster_status.md) fetches the current status of a cluster
- [kops_kube_config](/docs/data-sources/kube_config.md) fetches the kube config infos of a cluster

Finally, a special resource takes care of the cluster lifecyle:
- [kops_cluster_updater](/docs/resources/cluster_updater.md) manages cluster updates and/or rolling updates

[Provider configuration](/docs/guides/provider.md) holds cloud provider
authentication settings, currently only AWS is supported.

## Docs

The full documentation is available in the [docs](/docs/index.md) folder or
on the [terraform registry](https://registry.terraform.io/providers/eddycharly/kops/latest/docs)
provider page.

## Installing the provider

To install the provider, add it in the terraform `required_providers` set.

```hcl
terraform {
  required_providers {
    kops = {
      source  = "eddycharly/kops"
      version = "1.19.0-alpha.1"
    }
  }
}
```

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

## Example usage

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

resource "kops_instance_group" "master-0" {
  cluster_name = kops_cluster.cluster.name
  name         = "master-0"
  role         = "Master"
  min_size     = 1
  max_size     = 1
  machine_type = "t3.medium"
  subnets      = ["private-0"]
  depends_on   = [kops_cluster.cluster]
}

resource "kops_instance_group" "master-1" {
  cluster_name = kops_cluster.cluster.name
  name         = "master-1"
  role         = "Master"
  min_size     = 1
  max_size     = 1
  machine_type = "t3.medium"
  subnets      = ["private-1"]
  depends_on   = [kops_cluster.cluster]
}

resource "kops_instance_group" "master-2" {
  cluster_name = kops_cluster.cluster.name
  name         = "master-2"
  role         = "Master"
  min_size     = 1
  max_size     = 1
  machine_type = "t3.medium"
  subnets      = ["private-2"]
  depends_on   = [kops_cluster.cluster]
}

resource "kops_cluster_updater" "updater" {
  cluster_name = kops_cluster.cluster.name

  keepers = {
    cluster  = kops_cluster.cluster.revision,
    master-0 = kops_instance_group.master-0.revision,
    master-1 = kops_instance_group.master-1.revision,
    master-2 = kops_instance_group.master-2.revision
  }
}
```

More examples are available in the `/examples` folder:
- [Basic example](./examples/basic)
- [Aws profile example](./examples/aws-profile)
- [Aws assume role](./examples/aws-assume-role)
- [Bastion example](./examples/bastion)

## Importing an existing cluster

You can import an existing cluster by creating a `kops_cluster` configuration
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

## Importing an existing instance group

You can import an existing cluster by creating a `kops_instance_group` configuration
and running the `terraform import` command:

1. Create a terraform configuration:

    ```hcl
    provider "kops" {
      state_store = "s3://cluster.example.com"
    }

    resource "kops_instance_group" "ig-0" {
      cluster_namename  = "ig-0"
      name              = "cluster.example.com"
      
      // ....
    }
    ```

1. Run `terraform import`:

    ```shell
    terraform import kops_instance_group.ig-0 cluster.example.com/ig-0
    ```

**NOTE**: the id of the instance group to be imported must be given in the 
`cluster name/instance group name` format.

## Getting kubeconfig file

To retrieve the `kubeconfig` file for the cluster, run the following command:

```shell
kops export kubecfg --name cluster.example.com --state s3://cluster.example.com
```
