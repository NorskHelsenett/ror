import { Injectable } from '@angular/core';
import { ColumnDefinition } from '../models/columnDefinition';

@Injectable()
export class ColumnFactoryService {
  constructor() {}

  getColumnDefinitions(apiVersion, kind: string): ColumnDefinition[] {
    switch (kind) {
      case 'ClusterOrder':
        return this.getClusterOrder();
      case 'Pod':
        return this.getPod();
      case 'Node':
        return this.getNode();
      case 'Namespace':
        return this.getNamespace();
      case 'PersistentVolumeClaim':
        return this.getPersistentVolumeClaim();
      case 'Service':
        return this.getService();
      case 'Deployment':
        return this.getDeployment();
      case 'ReplicaSet':
        return this.getReplicaSet();
      case 'StatefulSet':
        return this.getStatefulSet();
      case 'DeamonSet':
        return this.getDeamonSet();
      case 'Ingress':
        return this.getIngress();
      case 'IngressClass':
        return this.getIngressClass();
      case 'StorageClass':
        return this.getStorageClass();
      case 'Application':
        return this.getApplication();
      case 'AppProject':
        return this.getAppProject();
      default:
        return [];
    }
  }

  private getClusterOrder(): ColumnDefinition[] {
    return [
      {
        field: 'spec.orderType',
        header: 'OrderType',
        type: 'text',
        enabled: true,
      },
      {
        field: 'spec.cluster',
        header: 'Cluster name',
        type: 'text',
        enabled: true,
      },
      {
        field: 'spec.provider',
        header: 'Provider',
        type: 'text',
        enabled: true,
      },
      {
        field: 'spec.orderBy',
        header: 'Ordered by',
        type: 'text',
        enabled: true,
      },
      {
        field: 'spec.ownerGroup',
        header: 'Owner group',
        type: 'text',
        enabled: true,
      },
      {
        field: 'spec.highAvailability',
        header: 'High availability',
        type: 'boolean',
        enabled: true,
      },
    ];
  }

  private getPod(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'Pod name',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'spec.nodeName',
        header: 'Node name',
        type: 'text',
        enabled: true,
      },
      {
        field: 'spec.serviceAccountName',
        header: 'Service account name',
        type: 'text',
        enabled: true,
      },
    ];
  }

  private getNode(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'Node name',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'spec.podcidr',
        header: 'Pod CIDR',
        type: 'text',
        enabled: true,
      },
      {
        field: 'status.capacity.cpu',
        header: 'CPU capasity',
        type: 'numeric',
        enabled: true,
      },
      {
        field: 'status.capacity.memory',
        header: 'Memory capasity',
        type: 'text',
        enabled: true,
      },
      {
        field: 'status.capacity.pods',
        header: 'Pods capasity',
        type: 'numeric',
        enabled: true,
      },
    ];
  }

  private getNamespace(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'Namespace name',
        type: 'text',
        enabled: true,
      },
    ];
  }

  private getPersistentVolumeClaim(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'PVC name',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'spec.storageClassName',
        header: 'Storage class name',
        type: 'text',
        enabled: true,
      },
      {
        field: 'spec.volumeName',
        header: 'Volume name',
        type: 'text',
        enabled: true,
      },
      {
        field: 'spec.volumeMode',
        header: 'Volume mode',
        type: 'text',
        enabled: true,
      },
    ];
  }

  private getService(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'Service name',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'spec.type',
        header: 'Type',
        type: 'text',
        enabled: true,
      },
      {
        field: 'spec.clusterIP',
        header: 'Cluster IP',
        type: 'text',
        enabled: true,
      },
    ];
  }

  private getDeployment(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'Deployment name',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'status.replicas',
        header: 'Replicas',
        type: 'numeric',
        enabled: true,
      },
      {
        field: 'status.availableReplicas',
        header: 'Available replicas',
        type: 'numeric',
        enabled: true,
      },
      {
        field: 'status.readyReplicas',
        header: 'Ready replicas',
        type: 'numeric',
        enabled: true,
      },
      {
        field: 'status.updatedReplicas',
        header: 'Updated replicas',
        type: 'numeric',
        enabled: true,
      },
    ];
  }

  private getReplicaSet(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'ReplicaSet name',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'spec.replicas',
        header: 'Replicas',
        type: 'numeric',
        enabled: true,
      },
    ];
  }

  private getStatefulSet(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'StatefulSet name',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'status.replicas',
        header: 'Replicas',
        type: 'numeric',
        enabled: true,
      },
      {
        field: 'status.availableReplicas',
        header: 'Available replicas',
        type: 'numeric',
        enabled: true,
      },
      {
        field: 'status.currentReplicas',
        header: 'Current replicas',
        type: 'numeric',
        enabled: true,
      },
      {
        field: 'status.readyReplicas',
        header: 'Ready replicas',
        type: 'numeric',
        enabled: true,
      },
      {
        field: 'status.updatedReplicas',
        header: 'Updated replicas',
        type: 'numeric',
        enabled: true,
      },
    ];
  }

  private getDeamonSet(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'DeamonSet name',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'status.numberReady',
        header: 'Number ready',
        type: 'numeric',
        enabled: true,
      },
      {
        field: 'status.numberUnavailable',
        header: 'Number unavailable',
        type: 'numeric',
        enabled: true,
      },
      {
        field: 'status.currentReplicas',
        header: 'Current replicas',
        type: 'numeric',
        enabled: true,
      },
      {
        field: 'status.numberAvailable',
        header: 'Number available',
        type: 'numeric',
        enabled: true,
      },
      {
        field: 'status.updatedNumberScheduled',
        header: 'Updated number scheduled',
        type: 'numeric',
        enabled: true,
      },
      {
        field: 'status.desiredNumberScheduled',
        header: 'Desired number scheduled',
        type: 'numeric',
        enabled: true,
      },
      {
        field: 'status.currentNumberScheduled',
        header: 'Current number scheduled',
        type: 'numeric',
        enabled: true,
      },
    ];
  }

  private getIngressClass(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'IngressClass name',
        type: 'text',
        enabled: true,
      },
      {
        field: 'spec.controller',
        header: 'Controller',
        type: 'text',
        enabled: true,
      },
    ];
  }

  private getIngress(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'Ingress name',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'spec.ingressClassName',
        header: 'Ingress class name',
        type: 'text',
        enabled: true,
      },
    ];
  }

  private getStorageClass(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'StorageClass name',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'provisioner',
        header: 'Provisioner',
        type: 'text',
        enabled: true,
      },
      {
        field: 'reclaimPolicy',
        header: 'Reclaim policy',
        type: 'text',
        enabled: true,
      },
      {
        field: 'volumeBindingMode',
        header: 'Volume binding mode',
        type: 'text',
        enabled: true,
      },
      {
        field: 'allowVolumeExpansion',
        header: 'Allow volume expansion',
        type: 'boolean',
        enabled: true,
      },
    ];
  }

  private getApplication(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'Application name',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'spec.destination.name',
        header: 'Destination name',
        type: 'text',
        enabled: true,
      },
      {
        field: 'spec.project',
        header: 'Project',
        type: 'text',
        enabled: true,
      },
    ];
  }

  private getAppProject(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'Project name',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'spec.description',
        header: 'Description',
        type: 'text',
        enabled: true,
      },
    ];
  }
}
