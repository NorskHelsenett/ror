import { Injectable } from '@angular/core';
import { ResourceType } from '../../core/models/resources/resourceType';
import { OwnerType } from '../../core/models/resources/ownerType';

@Injectable()
export class TypesService {
  constructor() {}

  getResourceTypes(): ResourceType[] {
    return [
      {
        displayName: 'Certificates',
        kind: 'Certificate',
        apiVersion: 'cert-manager.io/v1',
        clusterSpecific: true,
      },
      {
        displayName: 'Cluster compliance reports',
        kind: 'ClusterComplianceReport',
        apiVersion: 'aquasecurity.github.io/v1alpha1',
        clusterSpecific: true,
      },
      {
        displayName: 'Cluster vulnerability reports',
        kind: 'ClusterVulnerabilityReport',
        apiVersion: 'general.ror.internal/v1alpha1',
        clusterSpecific: true,
      },
      {
        displayName: 'Cluster orders',
        kind: 'ClusterOrder',
        apiVersion: 'general.ror.internal/v1alpha1',
        clusterSpecific: false,
      },
      {
        displayName: 'Exposed secret reports',
        kind: 'ExposedSecretReport',
        apiVersion: 'aquasecurity.github.io/v1alpha1',
        clusterSpecific: true,
      },
      {
        displayName: 'Pods',
        kind: 'Pod',
        apiVersion: 'v1',
        clusterSpecific: true,
      },
      {
        displayName: 'Nodes',
        kind: 'Node',
        apiVersion: 'v1',
        clusterSpecific: true,
      },
      {
        displayName: 'Namespaces',
        kind: 'Namespace',
        apiVersion: 'v1',
        clusterSpecific: true,
      },
      {
        displayName: 'Persistent volume claims',
        kind: 'PersistentVolumeClaim',
        apiVersion: 'v1',
        clusterSpecific: true,
      },
      {
        displayName: 'Services',
        kind: 'Service',
        apiVersion: 'v1',
        clusterSpecific: true,
      },
      {
        displayName: 'Deployments',
        kind: 'Deployment',
        apiVersion: 'apps/v1',
        clusterSpecific: true,
      },
      {
        displayName: 'Replica sets',
        kind: 'ReplicaSet',
        apiVersion: 'apps/v1',
        clusterSpecific: true,
      },
      {
        displayName: 'Routes',
        kind: 'Route',
        apiVersion: 'general.ror.internal/v1alpha1',
        clusterSpecific: true,
      },
      {
        displayName: 'Stateful sets',
        kind: 'StatefulSet',
        apiVersion: 'apps/v1',
        clusterSpecific: true,
      },
      {
        displayName: 'Daemon sets',
        kind: 'DeamonSet',
        apiVersion: 'apps/v1',
        clusterSpecific: true,
      },
      {
        displayName: 'Ingresses',
        kind: 'Ingress',
        apiVersion: 'networking.k8s.io/v1',
        clusterSpecific: true,
      },
      {
        displayName: 'Ingress classeses',
        kind: 'IngressClass',
        apiVersion: 'networking.k8s.io/v1',
        clusterSpecific: true,
      },
      {
        displayName: 'RBAC assessment reports',
        kind: 'RbacAssessmentReport',
        apiVersion: 'aquasecurity.github.io/v1alpha1',
        clusterSpecific: true,
      },
      {
        displayName: 'Notifications',
        kind: 'Notification',
        apiVersion: 'general.ror.internal/v1alpha1',
        clusterSpecific: true,
      },
      {
        displayName: 'Slack messages',
        kind: 'SlackMessage',
        apiVersion: 'general.ror.internal/v1alpha1',
        clusterSpecific: true,
      },
      {
        displayName: 'Storage classes',
        kind: 'StorageClass',
        apiVersion: 'storage.k8s.io/v1',
        clusterSpecific: true,
      },
      {
        displayName: 'ArgoCD applications',
        kind: 'Application',
        apiVersion: 'argoproj.io/v1alpha1',
        clusterSpecific: true,
      },
      {
        displayName: 'ArgoCD projects',
        kind: 'AppProject',
        apiVersion: 'argoproj.io/v1alpha1',
        clusterSpecific: true,
      },
      {
        displayName: 'Virtual machines',
        kind: 'Vm',
        apiVersion: 'general.ror.internal/v1alpha1',
        clusterSpecific: true,
      },
      {
        displayName: 'Vulnerability reports',
        kind: 'VulnerabilityReport',
        apiVersion: 'aquasecurity.github.io/v1alpha1',
        clusterSpecific: true,
      },
    ];
  }
}
