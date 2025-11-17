# Roadmap

### Release

| function              | milestone | done | description                                                                                               |
| --------------------- | --------- | ---- | --------------------------------------------------------------------------------------------------------- |
| create cluster        | 1.0       | 90%  | Create cluster from sclusterspec, allows integration with other teams managementssolution.                |
| upgrade/scale cluster | 1.0       | 20%  | Upgrade/scal cluster from changing sclusterspec, allows integration with other teams managementssolution. |
| decomission cluster   | 1.0       | 20%  | Decomission cluster, allows integration with other teams managementssolution.                             |

### Operate

| function                 | milestone | done               | description                                                       |
| ------------------------ | --------- | ------------------ | ----------------------------------------------------------------- |
| cluster inventory        | 1.0       | :heavy_check_mark: | Collect basic clustrerinformation                                 |
| cluster status           | 1.0       | :heavy_check_mark: | Calculate healthscore for the cluster based on established rules. |
| Configuration management | 1.0       |                    |                                                                   |
| - Dex integrations       | 1.0       | :heavy_check_mark: | Automatic provisioning of dex clientid and clientsecrets          |
| - tooling-config         | 1.0       | 90%                | Full config of tooling from ROR                                   |
| - Authorization          | 1.0       | :heavy_check_mark: | Configure access to clusters from ROR                             |

### Report

| function        | milestone | done               | description                                                         |
| --------------- | --------- | ------------------ | ------------------------------------------------------------------- |
| - Metrics       | 1.0       | :heavy_check_mark: | Gather basic cluster metrics                                        |
| - Security scan | 1.0       | :heavy_check_mark: | Run scans of cluster components like config, images certificates... |
