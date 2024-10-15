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
      },
      {
        field: 'spec.cluster',
        header: 'Cluster name',
        type: 'text',
      },
      {
        field: 'spec.provider',
        header: 'Provider',
        type: 'text',
      },
      {
        field: 'spec.orderBy',
        header: 'Ordered by',
        type: 'text',
      },
      {
        field: 'spec.ownerGroup',
        header: 'Owner group',
        type: 'text',
      },
      {
        field: 'spec.highAvailability',
        header: 'High availability',
        type: 'boolean',
      },
    ];
  }

  private getPod(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'Pod name',
        type: 'text',
      },
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
      },
      {
        field: 'spec.nodeName',
        header: 'Node name',
        type: 'text',
      },
      {
        field: 'spec.serviceAccountName',
        header: 'Service account name',
        type: 'text',
      },
    ];
  }

  private getNode(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'Node name',
        type: 'text',
      },
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
      },
      {
        field: 'spec.podcidr',
        header: 'Pod CIDR',
        type: 'text',
      },
      {
        field: 'status.capacity.cpu',
        header: 'CPU capasity',
        type: 'numeric',
      },
      {
        field: 'status.capacity.memory',
        header: 'Memory capasity',
        type: 'text',
      },
      {
        field: 'status.capacity.pods',
        header: 'Pods capasity',
        type: 'numeric',
      },
    ];
  }

  private getNamespace(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'Namespace name',
        type: 'text',
      },
    ];
  }

  private getPersistentVolumeClaim(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'PVC name',
        type: 'text',
      },
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
      },
      {
        field: 'spec.storageClassName',
        header: 'Storage class name',
        type: 'text',
      },
      {
        field: 'spec.volumeName',
        header: 'Volume name',
        type: 'text',
      },
      {
        field: 'spec.volumeMode',
        header: 'Volume mode',
        type: 'text',
      },
    ];
  }

  private getService(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'Service name',
        type: 'text',
      },
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
      },
      {
        field: 'spec.type',
        header: 'Type',
        type: 'text',
      },
      {
        field: 'spec.clusterIP',
        header: 'Cluster IP',
        type: 'text',
      },
    ];
  }

  private getDeployment(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'Deployment name',
        type: 'text',
      },
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
      },
      {
        field: 'status.replicas',
        header: 'Replicas',
        type: 'numeric',
      },
      {
        field: 'status.availableReplicas',
        header: 'Available replicas',
        type: 'numeric',
      },
      {
        field: 'status.readyReplicas',
        header: 'Ready replicas',
        type: 'numeric',
      },
      {
        field: 'status.updatedReplicas',
        header: 'Updated replicas',
        type: 'numeric',
      },
    ];
  }

  private getReplicaSet(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'ReplicaSet name',
        type: 'text',
      },
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
      },
      {
        field: 'spec.replicas',
        header: 'Replicas',
        type: 'numeric',
      },
    ];
  }

  private getStatefulSet(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'StatefulSet name',
        type: 'text',
      },
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
      },
      {
        field: 'status.replicas',
        header: 'Replicas',
        type: 'numeric',
      },
      {
        field: 'status.availableReplicas',
        header: 'Available replicas',
        type: 'numeric',
      },
      {
        field: 'status.currentReplicas',
        header: 'Current replicas',
        type: 'numeric',
      },
      {
        field: 'status.readyReplicas',
        header: 'Ready replicas',
        type: 'numeric',
      },
      {
        field: 'status.updatedReplicas',
        header: 'Updated replicas',
        type: 'numeric',
      },
    ];
  }

  private getDeamonSet(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'DeamonSet name',
        type: 'text',
      },
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
      },
      {
        field: 'status.numberReady',
        header: 'Number ready',
        type: 'numeric',
      },
      {
        field: 'status.numberUnavailable',
        header: 'Number unavailable',
        type: 'numeric',
      },
      {
        field: 'status.currentReplicas',
        header: 'Current replicas',
        type: 'numeric',
      },
      {
        field: 'status.numberAvailable',
        header: 'Number available',
        type: 'numeric',
      },
      {
        field: 'status.updatedNumberScheduled',
        header: 'Updated number scheduled',
        type: 'numeric',
      },
      {
        field: 'status.desiredNumberScheduled',
        header: 'Desired number scheduled',
        type: 'numeric',
      },
      {
        field: 'status.currentNumberScheduled',
        header: 'Current number scheduled',
        type: 'numeric',
      },
    ];
  }

  private getIngressClass(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'IngressClass name',
        type: 'text',
      },
      {
        field: 'spec.controller',
        header: 'Controller',
        type: 'text',
      },
    ];
  }

  private getIngress(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'Ingress name',
        type: 'text',
      },
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
      },
      {
        field: 'spec.ingressClassName',
        header: 'Ingress class name',
        type: 'text',
      },
    ];
  }

  private getStorageClass(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'StorageClass name',
        type: 'text',
      },
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
      },
      {
        field: 'provisioner',
        header: 'Provisioner',
        type: 'text',
      },
      {
        field: 'reclaimPolicy',
        header: 'Reclaim policy',
        type: 'text',
      },
      {
        field: 'volumeBindingMode',
        header: 'Volume binding mode',
        type: 'text',
      },
      {
        field: 'allowVolumeExpansion',
        header: 'Allow volume expansion',
        type: 'boolean',
      },
    ];
  }

  private getApplication(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'Application name',
        type: 'text',
      },
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
      },
      {
        field: 'spec.destination.name',
        header: 'Destination name',
        type: 'text',
      },
      {
        field: 'spec.project',
        header: 'Project',
        type: 'text',
      },
    ];
  }

  private getAppProject(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.name',
        header: 'Project name',
        type: 'text',
      },
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
      },
      {
        field: 'spec.description',
        header: 'Description',
        type: 'text',
      },
    ];
  }
}
