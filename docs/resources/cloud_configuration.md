# kops_cloud_configuration

| attribute | type | optional | required | computed |
| --- | --- | --- | --- | --- |
| `multizone` | Bool |  | Yes |  |
| `node_tags` | String |  | Yes |  |
| `node_instance_prefix` | String |  | Yes |  |
| `gce_service_account` | String |  | Yes |  |
| `disable_security_group_ingress` | Bool |  | Yes |  |
| `elb_security_group` | String |  | Yes |  |
| `v_sphere_username` | String |  | Yes |  |
| `v_sphere_password` | String |  | Yes |  |
| `v_sphere_server` | String |  | Yes |  |
| `v_sphere_datacenter` | String |  | Yes |  |
| `v_sphere_resource_pool` | String |  | Yes |  |
| `v_sphere_datastore` | String |  | Yes |  |
| `v_sphere_core_dns_server` | String |  | Yes |  |
| `spotinst_product` | String |  | Yes |  |
| `spotinst_orientation` | String |  | Yes |  |
| `openstack` | [OpenstackConfiguration](./OpenstackConfiguration.md) |  | Yes |  |
