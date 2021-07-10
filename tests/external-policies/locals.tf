locals {
  masterType  = "t3.medium"
  nodeType    = "t3.medium"
  clusterName = "cluster.example.com"
  dnsZone     = "example.com"
  vpcId       = "vpc-12345678"
  privateSubnets = [
    { subnetId = "subnet-1", zone = "us-test-1a" }
  ]
  utilitySubnets = [
    { subnetId = "subnet-2", zone = "us-test-1a" }
  ]
}
