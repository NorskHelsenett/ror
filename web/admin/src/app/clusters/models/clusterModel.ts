import { ClusterEnvironment } from '../../core/models/clusterEnvironment';
import { ClusterCapasity } from './clusterCapasity';
import { ClusterProvider } from './clusterProvider';
import { KubernetesVersions } from './kubernetesVersions';

export interface ClusterModel {
  provider: ClusterProvider;
  datacenter: string;
  workspace: string;
  environment: ClusterEnvironment;
  clusterName: string;
  tags?: string[];
  kubernetesVersion: KubernetesVersions;
  accessGroups?: string[];
  capasity: ClusterCapasity[];
  project?: string;
  responsible: string;
}
