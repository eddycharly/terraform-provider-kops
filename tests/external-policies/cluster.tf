resource "kops_cluster" "cluster" {
  name               = local.clusterName
  admin_ssh_key      = file("${path.module}/../id_rsa.pub")
  cloud_provider     = "aws"
  kubernetes_version = "1.19.12"
  dns_zone           = local.dnsZone
  network_id         = local.vpcId

  api {
    dns {}
  }

  authorization {
    rbac {}
  }

  iam {
    allow_container_registry = true
  }

  external_policies {
    key = "node"
    value = [
      "arn:aws:iam::123456789000:policy/test-policy-1",
      "arn:aws:iam::123456789000:policy/test-policy-2"
    ]
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
    name        = "utility-0"
    type        = "Utility"
    provider_id = local.utilitySubnets[0].subnetId
    zone        = local.utilitySubnets[0].zone
  }

  # etcd clusters
  etcd_cluster {
    name = "main"
    member {
      name           = "master-0"
      instance_group = "master-0"
    }
  }
  etcd_cluster {
    name = "events"
    member {
      name           = "master-0"
      instance_group = "master-0"
    }
  }

  kubelet {
    anonymous_auth {
      value = false
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

resource "kops_instance_group" "node-0" {
  cluster_name = kops_cluster.cluster.id
  name         = "node-0"
  role         = "Node"
  min_size     = 1
  max_size     = 2
  machine_type = local.nodeType
  subnets      = ["private-0"]
}
