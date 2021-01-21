resource "kops_cluster" "cluster" {
  name               = "cluster.example.com"
  admin_ssh_key      = file("${path.module}/../dummy_ssh.pub")
  cloud_provider     = "aws"
  kubernetes_version = "stable"
  dns_zone           = "example.com"
  network_id         = "net-0"

  networking {
    calico {}
  }

  topology {
    masters = "private"
    nodes   = "private"
    bastion {
      bastion_public_name = "bastion.cluster.example.com"
      load_balancer {
        additional_security_groups = ["sg-0"]
      }
    }
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

resource "kops_instance_group" "bastion-0" {
  cluster_name = kops_cluster.cluster.name
  name         = "bastion-0"
  role         = "Bastion"
  min_size     = 1
  max_size     = 1
  machine_type = "t3.medium"
  subnets      = ["private-0"]
  depends_on   = [kops_cluster.cluster]
}

resource "kops_cluster_updater" "updater" {
  cluster_name = kops_cluster.cluster.name

  # ensures rolling update happens after the cluster and instance groups are up to date
  depends_on   = [
    kops_cluster.cluster,
    kops_instance_group.master-0,
    kops_instance_group.master-1,
    kops_instance_group.master-2,
    kops_instance_group.bastion-0
  ]
}
