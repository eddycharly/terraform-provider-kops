provider "kops" {
  state_store = "s3://cluster.example.com"
  aws_profile = "profile"
}
