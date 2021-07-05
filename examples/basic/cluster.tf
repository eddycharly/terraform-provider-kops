resource "kops_cluster" "cluster" {
  name               = local.clusterName
  admin_ssh_key      = file("${path.module}/../dummy_ssh.pub")
  cloud_provider     = "aws"
  kubernetes_version = "stable"
  dns_zone           = local.dnsZone
  network_id         = local.vpcId

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
  machine_type = local.nodeType
  subnets      = ["private-0"]
}

resource "kops_instance_group" "node-1" {
  cluster_name = kops_cluster.cluster.id
  name         = "node-1"
  role         = "Node"
  min_size     = 1
  machine_type = local.nodeType
  subnets      = ["private-1"]
}

resource "kops_instance_group" "node-2" {
  cluster_name = kops_cluster.cluster.id
  name         = "node-2"
  role         = "Node"
  min_size     = 1
  machine_type = local.nodeType
  subnets      = ["private-2"]
}

resource "kops_cluster_updater" "updater" {
  cluster_name = kops_cluster.cluster.name

  keepers = {
    cluster  = kops_cluster.cluster.revision,
    master-0 = kops_instance_group.master-0.revision
    master-1 = kops_instance_group.master-1.revision
    master-2 = kops_instance_group.master-2.revision
    node-0   = kops_instance_group.node-0.revision
    node-1   = kops_instance_group.node-1.revision
    node-2   = kops_instance_group.node-2.revision
  }

  rolling_update {
    skip                = false
    fail_on_drain_error = true
    fail_on_validate    = true
    validate_count      = 1
  }

  validate {
    skip = false
  }
}

data "kops_kube_config" "kube_config" {
  cluster_name = kops_cluster.cluster.id
}
