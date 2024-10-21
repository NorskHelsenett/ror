import { Injectable } from '@angular/core';
import { ResourceType } from '../../core/models/resources/resourceType';
import { OwnerType } from '../../core/models/resources/ownerType';

@Injectable()
export class TypesService {
  constructor() {}

  getResourceTypes(): ResourceType[] {
    return [
      {
        kind: 'Certificate',
        apiVersion: 'cert-manager.io/v1',
        clusterSpecific: true,
      },
      {
        kind: 'ClusterComplianceReport',
        apiVersion: 'aquasecurity.github.io/v1alpha1',
        clusterSpecific: true,
      },
      {
        kind: 'ClusterVulnerabilityReport',
        apiVersion: 'general.ror.internal/v1alpha1',
        clusterSpecific: true,
      },
      {
        kind: 'ClusterOrder',
        apiVersion: 'general.ror.internal/v1alpha1',
        clusterSpecific: false,
      },
      {
        kind: 'ExposedSecretReport',
        apiVersion: 'aquasecurity.github.io/v1alpha1',
        clusterSpecific: true,
      },
      {
        kind: 'Pod',
        apiVersion: 'v1',
        clusterSpecific: true,
      },
      {
        kind: 'Node',
        apiVersion: 'v1',
        clusterSpecific: true,
      },
      {
        kind: 'Namespace',
        apiVersion: 'v1',
        clusterSpecific: true,
      },
      {
        kind: 'PersistentVolumeClaim',
        apiVersion: 'v1',
        clusterSpecific: true,
      },
      {
        kind: 'Service',
        apiVersion: 'v1',
        clusterSpecific: true,
      },
      {
        kind: 'Deployment',
        apiVersion: 'apps/v1',
        clusterSpecific: true,
      },
      {
        kind: 'ReplicaSet',
        apiVersion: 'apps/v1',
        clusterSpecific: true,
      },
      {
        kind: 'Route',
        apiVersion: 'general.ror.internal/v1alpha1',
        clusterSpecific: true,
      },
      {
        kind: 'StatefulSet',
        apiVersion: 'apps/v1',
        clusterSpecific: true,
      },
      {
        kind: 'DeamonSet',
        apiVersion: 'apps/v1',
        clusterSpecific: true,
      },
      {
        kind: 'Ingress',
        apiVersion: 'networking.k8s.io/v1',
        clusterSpecific: true,
      },
      {
        kind: 'IngressClass',
        apiVersion: 'networking.k8s.io/v1',
        clusterSpecific: true,
      },
      {
        kind: 'RbacAssessmentReport',
        apiVersion: 'aquasecurity.github.io/v1alpha1',
        clusterSpecific: true,
      },
      {
        kind: 'Notification',
        apiVersion: 'general.ror.internal/v1alpha1',
        clusterSpecific: true,
      },
      {
        kind: 'SlackMessage',
        apiVersion: 'general.ror.internal/v1alpha1',
        clusterSpecific: true,
      },
      {
        kind: 'StorageClass',
        apiVersion: 'storage.k8s.io/v1',
        clusterSpecific: true,
      },
      {
        kind: 'Application',
        apiVersion: 'argoproj.io/v1alpha1',
        clusterSpecific: true,
      },
      {
        kind: 'AppProject',
        apiVersion: 'argoproj.io/v1alpha1',
        clusterSpecific: true,
      },
      {
        kind: 'VulnerabilityReport',
        apiVersion: 'aquasecurity.github.io/v1alpha1',
        clusterSpecific: true,
      },
    ];
  }

  getOwnerTypes(): OwnerType[] {
    return [
      {
        scope: 'ror',
        subject: 'globalscope',
        clusterSpesific: false,
      },
      {
        scope: 'cluster',
        subject: null,
        clusterSpesific: true,
      },
    ];
  }
}
