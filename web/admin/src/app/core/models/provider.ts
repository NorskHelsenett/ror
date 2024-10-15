import { ClusterProvider } from '../../clusters/models/clusterProvider';

export interface Provider {
  name: string;
  type: ClusterProvider;
  disabled: boolean;
}

export interface ProviderKubernetesVersion {
  name: string;
  version: string;
  disabled: boolean;
}
