import { Injectable } from '@angular/core';
import { ColumnDefinition } from '../../resources/models/columnDefinition';

@Injectable({
  providedIn: 'root',
})
export class ColumnFactoryService {
  getColumnDefinitions(apiVersion: string, kind: string): ColumnDefinition[] {
    let columns: ColumnDefinition[] = [];

    if (!kind || !apiVersion) {
      columns = [
        {
          field: 'kind',
          header: 'Kind',
          type: 'text',
          enabled: true,
        },
        {
          field: 'apiVersion',
          header: 'Api version',
          type: 'text',
          enabled: true,
        },
      ];
    }

    if (kind === 'Application' && apiVersion === 'argoproj.io/v1alpha1') {
      columns = this.getApplication();
    }
    if (kind === 'AppProject' && apiVersion === 'argoproj.io/v1alpha1') {
      columns = this.getAppProject();
    }
    if (kind === 'Certificate' && apiVersion === 'cert-manager.io/v1') {
      columns = this.getCertificate();
    }
    if (kind === 'ClusterComplianceReport' && apiVersion === 'aquasecurity.github.io/v1alpha1') {
      columns = this.getClusterComplianceReport();
    }
    if (kind === 'ClusterVuulnerabilityReport' && apiVersion === 'general.ror.internal/v1alpha1') {
      columns = this.getClusterVuulnerabilityReport();
    }
    if (kind === 'ClusterOrder' && apiVersion === 'general.ror.internal/v1alpha1') {
      columns = this.getClusterOrder();
    }
    if (kind === 'ExposedSecretReport' && apiVersion === 'aquasecurity.github.io/v1alpha1') {
      columns = this.getExposedSecretReport();
    }
    if (kind === 'DeamonSet' && apiVersion === 'apps/v1') {
      columns = this.getDeamonSet();
    }
    if (kind === 'Deployment' && apiVersion === 'apps/v1') {
      columns = this.getDeployment();
    }
    if (kind === 'Ingress' && apiVersion === 'networking.k8s.io/v1') {
      columns = this.getIngress();
    }
    if (kind === 'IngressClass' && apiVersion === 'networking.k8s.io/v1') {
      columns = this.getIngressClass();
    }
    if (kind === 'Namespace' && apiVersion === 'v1') {
      columns = this.getNamespace();
    }
    if (kind === 'Node' && apiVersion === 'v1') {
      columns = this.getNode();
    }
    if (kind === 'Notification' && apiVersion === 'general.ror.internal/v1alpha1') {
      columns = this.getNotification();
    }
    if (kind === 'Pod' && apiVersion === 'v1') {
      columns = this.getPod();
    }
    if (kind === 'PersistentVolumeClaim' && apiVersion === 'v1') {
      columns = this.getPersistentVolumeClaim();
    }
    if (kind === 'RbacAssessmentReport' && apiVersion === 'aquasecurity.github.io/v1alpha1') {
      columns = this.getRbacAssessmentReport();
    }
    if (kind === 'ReplicaSet' && apiVersion === 'apps/v1') {
      columns = this.getReplicaSet();
    }
    if (kind === 'Route' && apiVersion === 'general.ror.internal/v1alpha1') {
      columns = this.getRoute();
    }
    if (kind === 'Service' && apiVersion === 'v1') {
      columns = this.getService();
    }
    if (kind === 'SlackMessage' && apiVersion === 'general.ror.internal/v1alpha1') {
      columns = this.getSlackMessage();
    }
    if (kind === 'StatefulSet' && apiVersion === 'apps/v1') {
      columns = this.getStatefulSet();
    }
    if (kind === 'StorageClass' && apiVersion === 'storage.k8s.io/v1') {
      columns = this.getStorageClass();
    }
    if ((kind === 'VirtualMachine' || kind == 'Vm') && apiVersion === 'general.ror.internal/v1alpha1') {
      columns = this.getRORVirtualMachines();
    }
    if (kind === 'VulnerabilityReport' && apiVersion === 'aquasecurity.github.io/v1alpha1') {
      columns = this.getVulnerabilityReport();
    }

    columns = this.postProcessColumnDefinitions(kind, apiVersion, columns);
    return columns;
  }
  getRORVirtualMachines(): ColumnDefinition[] {
    return [
      {
        field: 'vm.name',
        header: 'Name',
        type: 'text',
        enabled: true,
      },
      {
        field: 'vm.config.memorySize',
        header: 'Memory size',
        type: 'numeric',
        enabled: true,
      },
      {
        field: 'vm.config.cpuCount',
        header: 'CPU count',
        type: 'numeric',
        enabled: true,
      },
    ];
  }
  getApplication(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.creationTimestamp',
        header: 'Creation Time',
        type: 'date',
        enabled: true,
      },
      {
        field: 'application.spec.source.chart',
        header: 'Chart',
        type: 'text',
        enabled: true,
      },
      {
        field: 'application.spec.source.path',
        header: 'Path',
        type: 'text',
        enabled: true,
      },
      {
        field: 'application.spec.source.targetRevision',
        header: 'Target Revision',
        type: 'text',
        enabled: true,
      },
      {
        field: 'application.status.health.status',
        header: 'Health status',
        type: 'text',
        enabled: true,
      },
      {
        field: 'application.status.sync.status',
        header: 'Sync status',
        type: 'text',
        enabled: true,
      },
    ];
  }
  getAppProject(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.creationTimestamp',
        header: 'Creation Time',
        type: 'date',
        enabled: true,
      },
    ];
  }
  getCertificate(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.creationTimestamp',
        header: 'Creation Time',
        type: 'date',
        enabled: true,
      },
      {
        field: 'certificate.spec.secretName',
        header: 'Secret Name',
        type: 'text',
        enabled: true,
      },
      {
        field: 'certificate.spec.dnsNames',
        header: 'DNS Names',
        type: 'array',
        enabled: true,
      },
      {
        field: 'certificate.status.notAfter',
        header: 'Not After',
        type: 'date',
        enabled: false,
      },
      {
        field: 'certificate.status.notBefore',
        header: 'Not Before',
        type: 'date',
        enabled: false,
      },
      {
        field: 'certificate.status.renewalTime',
        header: 'Renewal Time',
        type: 'date',
        enabled: false,
      },
    ];
  }
  getClusterComplianceReport(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.creationTimestamp',
        header: 'Creation Time',
        type: 'date',
        enabled: true,
      },
      {
        field: 'clustercompliancereport.report.summary.complianceStatus',
        header: 'Compliance Status',
        type: 'text',
        enabled: true,
      },
      {
        field: 'clustercompliancereport.report.summary.complianceScore',
        header: 'Compliance Score',
        type: 'numeric',
        enabled: true,
      },
    ];
  }
  getClusterVuulnerabilityReport(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.creationTimestamp',
        header: 'Creation Time',
        type: 'date',
        enabled: true,
      },
      {
        field: 'clustervuulnerabilityreport.report.summary.criticalCount',
        header: 'Critical count',
        type: 'numeric',
        enabled: true,
      },
      {
        field: 'clustervuulnerabilityreport.report.summary.highCount',
        header: 'High count',
        type: 'numeric',
        enabled: true,
      },
    ];
  }
  getClusterOrder(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.creationTimestamp',
        header: 'Creation Time',
        type: 'date',
        enabled: true,
      },
      {
        field: 'clusterorder.spec.cluster',
        header: 'Cluster ID',
        type: 'text',
        enabled: true,
      },
    ];
  }
  getExposedSecretReport(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.creationTimestamp',
        header: 'Creation Time',
        type: 'date',
        enabled: true,
      },
      {
        field: 'exposedsecretreport.report.summary.criticalCount',
        header: 'Critical count',
        type: 'numeric',
        enabled: true,
      },
      {
        field: 'exposedsecretreport.report.summary.highCount',
        header: 'High count',
        type: 'numeric',
        enabled: true,
      },
    ];
  }
  getDeamonSet(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.creationTimestamp',
        header: 'Creation Time',
        type: 'date',
        enabled: true,
      },
      {
        field: 'deamonset.spec.selector.matchLabels',
        header: 'Match Labels',
        type: 'array',
        enabled: true,
      },
    ];
  }
  getDeployment(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.creationTimestamp',
        header: 'Creation Time',
        type: 'date',
        enabled: true,
      },
      {
        field: 'deployment.status.replicas',
        header: 'Replicas',
        type: 'numeric',
        enabled: true,
      },
      {
        field: 'deployment.status.readyReplicas',
        header: 'Replicas',
        type: 'numeric',
        enabled: true,
      },
      {
        field: 'deployment.status.availableReplicas',
        header: 'Replicas',
        type: 'numeric',
        enabled: true,
      },
    ];
  }
  getIngress(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.creationTimestamp',
        header: 'Creation Time',
        type: 'date',
        enabled: true,
      },
      {
        field: 'ingress.spec.ingressClassName',
        header: 'Ingress class name',
        type: 'text',
        enabled: true,
      },
    ];
  }
  getIngressClass(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.creationTimestamp',
        header: 'Creation Time',
        type: 'date',
        enabled: true,
      },
      {
        field: 'ingressclass.spec.controller',
        header: 'Controller',
        type: 'text',
        enabled: true,
      },
    ];
  }
  getPersistentVolumeClaim(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.creationTimestamp',
        header: 'Creation Time',
        type: 'date',
        enabled: true,
      },
      {
        field: 'persistentvolumeclaim.spec.volumeName',
        header: 'Volume Name',
        type: 'text',
        enabled: true,
      },
    ];
  }
  getNamespace(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.creationTimestamp',
        header: 'Creation Time',
        type: 'date',
        enabled: true,
      },
    ];
  }
  getNotification(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.creationTimestamp',
        header: 'Creation Time',
        type: 'date',
        enabled: true,
      },
      {
        field: 'notification.spec.message',
        header: 'Message',
        type: 'text',
        enabled: true,
      },
      {
        field: 'notification.spec.severity',
        header: 'Severity',
        type: 'text',
        enabled: true,
      },
      {
        field: 'notification.spec.target',
        header: 'Target',
        type: 'text',
        enabled: true,
      },
      {
        field: 'notification.spec.type',
        header: 'Type',
        type: 'text',
        enabled: true,
      },
    ];
  }
  getPod(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'pod.status.phase',
        header: 'Phase',
        type: 'text',
        enabled: true,
      },
      {
        field: 'pod.status.startTime',
        header: 'Start time',
        type: 'date',
        enabled: true,
      },
    ];
  }
  getNode(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.labels',
        header: 'Labels',
        type: 'array',
        enabled: true,
      },
      {
        field: 'metadata.creationTimestamp',
        header: 'Creation Time',
        type: 'date',
        enabled: true,
      },
      {
        field: 'node.status.nodeInfo.architecture',
        header: 'Architecture',
        type: 'text',
        enabled: true,
      },
      {
        field: 'node.status.nodeInfo.osImage',
        header: 'Os image',
        type: 'text',
        enabled: true,
      },
    ];
  }
  getRbacAssessmentReport(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.creationTimestamp',
        header: 'Creation time',
        type: 'date',
        enabled: true,
      },
      {
        field: 'rbacassessmentreport.report.summary.criticalCount',
        header: 'Critical count',
        type: 'numeric',
        enabled: true,
      },
      {
        field: 'rbacassessmentreport.report.summary.highCount',
        header: 'High count',
        type: 'numeric',
        enabled: true,
      },
    ];
  }
  getReplicaSet(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.creationTimestamp',
        header: 'Creation Time',
        type: 'date',
        enabled: true,
      },
      {
        field: 'replicaset.status.replicas',
        header: 'Replicas',
        type: 'numeric',
        enabled: true,
      },
      {
        field: 'replicaset.status.readyReplicas',
        header: 'Ready Replicas',
        type: 'numeric',
        enabled: true,
      },
      {
        field: 'replicaset.status.availableReplicas',
        header: 'Available Replicas',
        type: 'numeric',
        enabled: true,
      },
    ];
  }
  getRoute(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.creationTimestamp',
        header: 'Creation Time',
        type: 'date',
        enabled: true,
      },
      {
        field: 'route.spec.host',
        header: 'Host',
        type: 'text',
        enabled: true,
      },
      {
        field: 'route.spec.path',
        header: 'Path',
        type: 'text',
        enabled: true,
      },
    ];
  }
  getService(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.creationTimestamp',
        header: 'Creation Time',
        type: 'date',
        enabled: true,
      },
      {
        field: 'service.spec.clusterIP',
        header: 'Cluster IP',
        type: 'text',
        enabled: true,
      },
      {
        field: 'service.spec.type',
        header: 'Type',
        type: 'text',
        enabled: true,
      },
    ];
  }
  getSlackMessage(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.creationTimestamp',
        header: 'Creation Time',
        type: 'date',
        enabled: true,
      },
      {
        field: 'slackmessage.spec.message',
        header: 'Message',
        type: 'text',
        enabled: true,
      },
      {
        field: 'slackmessage.spec.severity',
        header: 'Severity',
        type: 'text',
        enabled: true,
      },
      {
        field: 'slackmessage.spec.target',
        header: 'Target',
        type: 'text',
        enabled: true,
      },
      {
        field: 'slackmessage.spec.type',
        header: 'Type',
        type: 'text',
        enabled: true,
      },
    ];
  }
  getStatefulSet(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.creationTimestamp',
        header: 'Creation Time',
        type: 'date',
        enabled: true,
      },
      {
        field: 'statefulset.status.replicas',
        header: 'Replicas',
        type: 'numeric',
        enabled: true,
      },
      {
        field: 'statefulset.status.readyReplicas',
        header: 'Ready Replicas',
        type: 'numeric',
        enabled: true,
      },
      {
        field: 'statefulset.status.currentReplicas',
        header: 'Current Replicas',
        type: 'numeric',
        enabled: true,
      },
    ];
  }

  getStorageClass(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.creationTimestamp',
        header: 'Creation Time',
        type: 'date',
        enabled: true,
      },
      {
        field: 'storageclass.provisioner',
        header: 'Provisioner',
        type: 'text',
        enabled: true,
      },
      {
        field: 'storageclass.reclaimPolicy',
        header: 'Reclaim Policy',
        type: 'text',
        enabled: true,
      },
      {
        field: 'storageclass.volumeBindingMode',
        header: 'Volume Binding Mode',
        type: 'text',
        enabled: true,
      },
    ];
  }
  getVulnerabilityReport(): ColumnDefinition[] {
    return [
      {
        field: 'metadata.namespace',
        header: 'Namespace',
        type: 'text',
        enabled: true,
      },
      {
        field: 'metadata.creationTimestamp',
        header: 'Creation time',
        type: 'date',
        enabled: true,
      },
      {
        field: 'vulnerabilityreport.report.summary.criticalCount',
        header: 'Critical count',
        type: 'numeric',
        enabled: true,
      },
      {
        field: 'vulnerabilityreport.report.summary.highCount',
        header: 'High count',
        type: 'numeric',
        enabled: true,
      },
    ];
  }

  private postProcessColumnDefinitions(kind: string, apiVersion: string, columns: ColumnDefinition[]): ColumnDefinition[] {
    columns.unshift({
      field: 'metadata.name',
      header: 'Name',
      type: 'text',
      enabled: true,
    });

    columns.push({
      field: 'rormeta.ownerref.subject',
      header: 'Owner',
      type: 'text',
      enabled: true,
    });

    return columns;
  }
}
