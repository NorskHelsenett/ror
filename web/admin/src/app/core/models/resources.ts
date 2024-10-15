export interface ResourceNamespace {
  apiVersion: string;
  kind: string;
  metadata: ResourceMetadata;
}

export interface ResourceMetadata {
  name: string;
  resourceVersion: string;
  creationTimestamp: string;
  labels: { [key: string]: string };
  annotations: { [key: string]: string };
  uid: string;
  namespace: string;
  generation: number;
  ownerReferences: ResourceMetadataOwnerReference[];
}

export interface ResourceMetadataOwnerReference {
  apiVersion: string;
  kind: string;
  name: string;
  uid: string;
}
