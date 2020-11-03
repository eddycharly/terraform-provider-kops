# kops_etcd_cluster_spec

| attribute | type | optional/required | computed |
| --- | --- | --- | --- |
| `name` | String | Required |  |
| `provider` | String | Optional |  |
| `members` | List([EtcdMemberSpec](./EtcdMemberSpec.generated.md)) | Required |  |
| `enable_etcd_tls` | Bool | Optional |  |
| `enable_tls_auth` | Bool | Optional |  |
| `version` | String | Optional |  |
| `leader_election_timeout` | Duration | Optional |  |
| `heartbeat_interval` | Duration | Optional |  |
| `image` | String | Optional |  |
| `backups` | [EtcdBackupSpec](./EtcdBackupSpec.generated.md) | Optional |  |
| `manager` | [EtcdManagerSpec](./EtcdManagerSpec.generated.md) | Optional |  |
| `memory_request` | Quantity | Optional |  |
| `cpu_request` | Quantity | Optional |  |
