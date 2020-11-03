resource "kops_cluster" "cluster" {
  name               = "cluster.example.com"
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
      bastion_public_name        = "bastion.cluster.example.com"
      additional_security_groups = ["sg-0"]
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

  # master instance groups
  instance_group {
    name         = "master-0"
    role         = "Master"
    min_size     = 1
    max_size     = 1
    machine_type = "t3.medium"
    subnets      = ["private-0"]
  }

  instance_group {
    name         = "master-1"
    role         = "Master"
    min_size     = 1
    max_size     = 1
    machine_type = "t3.medium"
    subnets      = ["private-1"]
  }

  instance_group {
    name         = "master-2"
    role         = "Master"
    min_size     = 1
    max_size     = 1
    machine_type = "t3.medium"
    subnets      = ["private-2"]
  }

  # bastion instance groups
  instance_group {
    name         = "master-0"
    role         = "Bastion"
    min_size     = 1
    max_size     = 1
    machine_type = "t3.medium"
    subnets      = ["private-0"]
  }

  # etcd clusters
  etcd_cluster {
    name = "main"

    members {
      name           = "master-0"
      instance_group = "master-0"
    }

    members {
      name           = "master-1"
      instance_group = "master-1"
    }

    members {
      name           = "master-2"
      instance_group = "master-2"
    }
  }

  etcd_cluster {
    name = "events"

    members {
      name           = "master-0"
      instance_group = "master-0"
    }

    members {
      name           = "master-1"
      instance_group = "master-1"
    }

    members {
      name           = "master-2"
      instance_group = "master-2"
    }
  }
}
