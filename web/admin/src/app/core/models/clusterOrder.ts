import { ClusterEnvironment } from './clusterEnvironment';
import { Criticality, Sensitivity } from './vulnerabilityReport';

export interface ClusterOrder {
  apiVersion: string;
  kind: string;
  metadata: ClusterOrderMetadata;
  spec: ClusterOrderSpec;
  status: ClusterOrderStatus;
}

export interface ClusterOrderMetadata {
  name: string;
  resourceVersion: string;
  createTimestamp: string;
  uid: string;
}

export interface ClusterOrderSpec {
  provider: string;
  cluster: string;
  projectId: string;
  orderBy: string;
  environment: ClusterEnvironment;
  criticality: Criticality;
  sensitivity: Sensitivity;
  highAvailability: boolean;
  nodePools: NodePool[];
  providerConfig: ProviderConfig;
  ownerGroup: string;
}

export interface ClusterOrderStatus {
  status: string;
  phase: string;
  conditions: any[];
  createdTime: string;
  updatedTime: string;
  lastObservedTime: string;
}

export interface NodePool {
  machineClass: string;
  count: number;
}

export interface ProviderConfig {}

export interface TanzuConfig extends ProviderConfig {
  namespaceId: string;
  datacenterId: string;
}

export interface AzureConfig extends ProviderConfig {
  subscriptionId: string;
  resourceGroupName: string;
  region: string;
}

export interface ClusterOrderModel {
  orderType: ClusterOrderType;
  projectId: string;
  orderBy: string;
  provider: string;
  cluster: string;
  environment: number;
  criticality: number;
  sensitivity: number;
  highAvailability: boolean;
  nodePools: NodePool[];
  serviceTags: any;
  ownerGroup: string;
  providerconfig: ProviderConfig;
  k8sVersion: string;
}

export enum ClusterOrderType {
  Unknown = '',
  Create = 'Create',
  Update = 'Update',
  Delete = 'Delete',
}
