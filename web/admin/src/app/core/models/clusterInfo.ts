import { ClusterMetadata } from './clusterMetadata';
export interface ClusterInfo {
  id: string;
  clusterId: string;
  clusterName: string;
  metadata: ClusterMetadata;
}
