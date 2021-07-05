locals {
  masterType  = "t3.medium"
  nodeType    = "t3.medium"
  clusterName = "cluster.example.com"
  dnsZone     = "example.com"
  vpcId       = "vpc-id"
  privateSubnets = [
    { subnetId = "private-subnet-0", zone = "zone-0" },
    { subnetId = "private-subnet-1", zone = "zone-1" },
    { subnetId = "private-subnet-1", zone = "zone-2" }
  ]
  utilitySubnets = [
    { subnetId = "utility-subnet-0", zone = "zone-0" },
    { subnetId = "utility-subnet-1", zone = "zone-1" },
    { subnetId = "utility-subnet-1", zone = "zone-2" }
  ]
}