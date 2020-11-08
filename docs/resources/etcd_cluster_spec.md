# kops_etcd_cluster_spec

| attribute | type | optional | required | computed |
| --- | --- | --- | --- | --- |
| `name` | String | Yes |  |  |
| `provider` | String |  | Yes |  |
| `members` | List([EtcdMemberSpec](./EtcdMemberSpec.md)) | Yes |  |  |
| `enable_etcd_tls` | Bool |  | Yes |  |
| `enable_tls_auth` | Bool |  | Yes |  |
| `version` | String |  | Yes |  |
| `leader_election_timeout` | Duration |  | Yes |  |
| `heartbeat_interval` | Duration |  | Yes |  |
| `image` | String |  | Yes |  |
| `backups` | [EtcdBackupSpec](./EtcdBackupSpec.md) |  | Yes |  |
| `manager` | [EtcdManagerSpec](./EtcdManagerSpec.md) |  | Yes |  |
| `memory_request` | Quantity |  | Yes |  |
| `cpu_request` | Quantity |  | Yes |  |
