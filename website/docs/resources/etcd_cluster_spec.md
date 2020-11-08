# kops_etcd_cluster_spec

| attribute | type | optional | required | computed |
| --- | --- | --- | --- | --- |
| `name` | String | :white_check_mark: |  |  |
| `provider` | String |  | :white_check_mark: |  |
| `members` | List([EtcdMemberSpec](./EtcdMemberSpec.md)) | :white_check_mark: |  |  |
| `enable_etcd_tls` | Bool |  | :white_check_mark: |  |
| `enable_tls_auth` | Bool |  | :white_check_mark: |  |
| `version` | String |  | :white_check_mark: |  |
| `leader_election_timeout` | Duration |  | :white_check_mark: |  |
| `heartbeat_interval` | Duration |  | :white_check_mark: |  |
| `image` | String |  | :white_check_mark: |  |
| `backups` | [EtcdBackupSpec](./EtcdBackupSpec.md) |  | :white_check_mark: |  |
| `manager` | [EtcdManagerSpec](./EtcdManagerSpec.md) |  | :white_check_mark: |  |
| `memory_request` | Quantity |  | :white_check_mark: |  |
| `cpu_request` | Quantity |  | :white_check_mark: |  |
