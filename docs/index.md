# kOps Provider

`terraform-provider-kops` brings [kOps](https://github.com/kubernetes/kops)
into terraform in a fully managed way, enabling idempotency through direct
integration with the kOps api:
- No `local_exec`
- No yaml templating
- No CLI invocations

... just **pure go code.**

Currently using kOps `v1.23.1` and compatible with terraform `0.12` and higher.

~> For now, provisioning the network is not supported. The network must
be created separately and given to the provider through cluster attribute
`network_id` and subnets attributes `provider_id`.

!> The provider has been tested only with AWS and calico networking.
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
`kops rolling-update cluster` to recycle instance groups when the change.
With this provider, this is all taken care of and you should never need to run
kOps manually.

## How does it work

The provider declares resources to declare the state of the cluster:
- [kops_cluster](/docs/resources/cluster.md) defines the desired state of a cluster
- [kops_instance_group](/docs/resources/instance_group.md) defines the desired state of a cluster instance group

The provider also declares data sources to fetch the state of the cluster and
use it in your terraform code:
- [kops_cluster](/docs/data-sources/cluster) fetches the current state of a cluster
- [kops_instance_group](/docs/data-sources/instance_group) fetches the current state of a cluster instance group
- [kops_cluster_status](/docs/data-sources/cluster_status) fetches the current status of a cluster
- [kops_kube_config](/docs/data-sources/kube_config) fetches the kube config infos of a cluster

Finally, a special resource takes care of the cluster lifecyle:
- [kops_cluster_updater](/docs/resources/cluster_updater) manages cluster updates and/or rolling updates

[Provider configuration](/docs/guides/provider) holds cloud provider
authentication settings, currently only AWS is supported.

## Example usage

```hcl
provider "kops" {
  state_store = "s3://cluster.example.com"
  // optionally set up your cloud provider access config
  aws {
    profile = "example_profile"
  }
}

locals {
  masterType  = "t3.medium"
  nodeType    = "t3.medium"
  clusterName = "cluster.example.com"
  dnsZone     = "example.com"
  vpcId       = "vpc-id"
  privateSubnets = [
    { subnetId = "private-subnet-0", zone = "zone-0" },
    { subnetId = "private-subnet-1", zone = "zone-1" },
    { subnetId = "private-subnet-2", zone = "zone-2" }
  ]
  utilitySubnets = [
    { subnetId = "utility-subnet-0", zone = "zone-0" },
    { subnetId = "utility-subnet-1", zone = "zone-1" },
    { subnetId = "utility-subnet-2", zone = "zone-2" }
  ]
}

resource "kops_cluster" "cluster" {
  name               = local.clusterName
  admin_ssh_key      = file("${path.module}/../dummy_ssh.pub")
  cloud_provider     = "aws"
  kubernetes_version = "stable"
  dns_zone           = local.dnsZone
  network_id         = local.vpcId

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

  # private subnets
  subnet {
    name        = "private-0"
    type        = "Private"
    provider_id = local.privateSubnets[0].subnetId
    zone        = local.privateSubnets[0].zone
  }
  subnet {
    name        = "private-1"
    type        = "Private"
    provider_id = local.privateSubnets[1].subnetId
    zone        = local.privateSubnets[1].zone
  }
  subnet {
    name        = "private-2"
    type        = "Private"
    provider_id = local.privateSubnets[2].subnetId
    zone        = local.privateSubnets[2].zone
  }
  subnet {
    name        = "utility-0"
    type        = "Utility"
    provider_id = local.utilitySubnets[0].subnetId
    zone        = local.utilitySubnets[0].zone
  }
  subnet {
    name        = "utility-1"
    type        = "Utility"
    provider_id = local.utilitySubnets[1].subnetId
    zone        = local.utilitySubnets[1].zone
  }
  subnet {
    name        = "utility-2"
    type        = "Utility"
    provider_id = local.utilitySubnets[2].subnetId
    zone        = local.utilitySubnets[2].zone
  }

  # etcd clusters
  etcd_cluster {
    name = "main"
    member {
      name           = "master-0"
      instance_group = "master-0"
    }
    member {
      name           = "master-1"
      instance_group = "master-1"
    }
    member {
      name           = "master-2"
      instance_group = "master-2"
    }
  }
  etcd_cluster {
    name = "events"
    member {
      name           = "master-0"
      instance_group = "master-0"
    }
    member {
      name           = "master-1"
      instance_group = "master-1"
    }
    member {
      name           = "master-2"
      instance_group = "master-2"
    }
  }
}

resource "kops_instance_group" "master-0" {
  cluster_name = kops_cluster.cluster.id
  name         = "master-0"
  role         = "Master"
  min_size     = 1
  max_size     = 1
  machine_type = local.masterType
  subnets      = ["private-0"]
}

resource "kops_instance_group" "master-1" {
  cluster_name = kops_cluster.cluster.id
  name         = "master-1"
  role         = "Master"
  min_size     = 1
  max_size     = 1
  machine_type = local.masterType
  subnets      = ["private-1"]
}

resource "kops_instance_group" "master-2" {
  cluster_name = kops_cluster.cluster.id
  name         = "master-2"
  role         = "Master"
  min_size     = 1
  max_size     = 1
  machine_type = local.masterType
  subnets      = ["private-2"]
}

resource "kops_instance_group" "node-0" {
  cluster_name = kops_cluster.cluster.id
  name         = "node-0"
  role         = "Node"
  min_size     = 1
  max_size     = 2
  machine_type = local.nodeType
  subnets      = ["private-0"]
}

resource "kops_instance_group" "node-1" {
  cluster_name = kops_cluster.cluster.id
  name         = "node-1"
  role         = "Node"
  min_size     = 1
  max_size     = 2
  machine_type = local.nodeType
  subnets      = ["private-1"]
}

resource "kops_instance_group" "node-2" {
  cluster_name = kops_cluster.cluster.id
  name         = "node-2"
  role         = "Node"
  min_size     = 1
  max_size     = 2
  machine_type = local.nodeType
  subnets      = ["private-2"]
}

resource "kops_cluster_updater" "updater" {
  cluster_name = kops_cluster.cluster.id

  keepers = {
    cluster  = kops_cluster.cluster.revision,
    master-0 = kops_instance_group.master-0.revision
    master-1 = kops_instance_group.master-1.revision
    master-2 = kops_instance_group.master-2.revision
    node-0   = kops_instance_group.node-0.revision
    node-1   = kops_instance_group.node-1.revision
    node-2   = kops_instance_group.node-2.revision
  }
}
```

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
      cluster_name = "cluster.example.com"
      name         = "ig-0"
      
      // ....
    }
    ```

1. Run `terraform import`:

    ```shell
    terraform import kops_instance_group.ig-0 cluster.example.com/ig-0
    ```

~> The id of the instance group to be imported must be given in the 
`cluster name/instance group name` format.

## Getting kubeconfig file

To retrieve the `kubeconfig` file for the cluster, run the following command:

```shell
kops export kubecfg --admin --name cluster.example.com --state s3://cluster.example.com
```
