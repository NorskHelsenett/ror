/* Do not change, this code is generated from Golang structs */

export class ResourceNotificationSpec {
  owner: RorResourceOwnerReference;
  message: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.owner = this.convertValues(source['owner'], RorResourceOwnerReference);
    this.message = source['message'];
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceNotification {
  spec: ResourceNotificationSpec;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.spec = this.convertValues(source['spec'], ResourceNotificationSpec);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceSlackMessageStatus {
  result: number;
  timestamp: Time;
  error: any;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.result = source['result'];
    this.timestamp = this.convertValues(source['timestamp'], Time);
    this.error = source['error'];
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceSlackMessageSpec {
  channelId: string;
  message: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.channelId = source['channelId'];
    this.message = source['message'];
  }
}
export class ResourceSlackMessage {
  spec: ResourceSlackMessageSpec;
  status: ResourceSlackMessageStatus[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.spec = this.convertValues(source['spec'], ResourceSlackMessageSpec);
    this.status = this.convertValues(source['status'], ResourceSlackMessageStatus);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceRouteSlackReceiver {
  channelId: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.channelId = source['channelId'];
  }
}
export class ResourceRouteReceiver {
  slack: ResourceRouteSlackReceiver[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.slack = this.convertValues(source['slack'], ResourceRouteSlackReceiver);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceRouteMessageType {
  apiVersion: string;
  kind: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.apiVersion = source['apiVersion'];
    this.kind = source['kind'];
  }
}
export class ResourceRouteSpec {
  messageType: ResourceRouteMessageType;
  receivers: ResourceRouteReceiver;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.messageType = this.convertValues(source['messageType'], ResourceRouteMessageType);
    this.receivers = this.convertValues(source['receivers'], ResourceRouteReceiver);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceRoute {
  spec: ResourceRouteSpec;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.spec = this.convertValues(source['spec'], ResourceRouteSpec);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class Time {
  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
  }
}
export class ResourceClusterVulnerabilityReportStatus {
  vulnerabilityID: string;
  severity: string;
  score: number;
  title: string;
  resource: string;
  primaryLink: string;
  installedVersion: string;
  fixedVersion: string;
  lastObserved: Time;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.vulnerabilityID = source['vulnerabilityID'];
    this.severity = source['severity'];
    this.score = source['score'];
    this.title = source['title'];
    this.resource = source['resource'];
    this.primaryLink = source['primaryLink'];
    this.installedVersion = source['installedVersion'];
    this.fixedVersion = source['fixedVersion'];
    this.lastObserved = this.convertValues(source['lastObserved'], Time);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceClusterVulnerabilityReportSummary {
  critical: number;
  high: number;
  medium: number;
  low: number;
  unknown: number;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.critical = source['critical'];
    this.high = source['high'];
    this.medium = source['medium'];
    this.low = source['low'];
    this.unknown = source['unknown'];
  }
}
export class ResourceClusterVulnerabilityReportReport {
  summary: ResourceClusterVulnerabilityReportSummary;
  vulnerabilities: { [key: string]: ResourceClusterVulnerabilityReportStatus };

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.summary = this.convertValues(source['summary'], ResourceClusterVulnerabilityReportSummary);
    this.vulnerabilities = this.convertValues(source['vulnerabilities'], ResourceClusterVulnerabilityReportStatus, true);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceClusterVulnerabilityReport {
  report: ResourceClusterVulnerabilityReportReport;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.report = this.convertValues(source['report'], ResourceClusterVulnerabilityReportReport);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceClusterComplianceReport {
  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
  }
}
export class ResourceConfigurationSpec {
  type: string;
  b64enc: boolean;
  data: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.type = source['type'];
    this.b64enc = source['b64enc'];
    this.data = source['data'];
  }
}
export class ResourceConfiguration {
  spec: ResourceConfigurationSpec;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.spec = this.convertValues(source['spec'], ResourceConfigurationSpec);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceProjectSpecRole {
  upn: string;
  name: string;
  role: string;
  email: string;
  phone: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.upn = source['upn'];
    this.name = source['name'];
    this.role = source['role'];
    this.email = source['email'];
    this.phone = source['phone'];
  }
}
export class ResourceProjectSpec {
  projectName: string;
  description: string;
  active: boolean;
  createdTime: string;
  updatedTime: string;
  roles: ResourceProjectSpecRole[];
  workorder: string;
  serviceTag: string;
  tags: string[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.projectName = source['projectName'];
    this.description = source['description'];
    this.active = source['active'];
    this.createdTime = source['createdTime'];
    this.updatedTime = source['updatedTime'];
    this.roles = this.convertValues(source['roles'], ResourceProjectSpecRole);
    this.workorder = source['workorder'];
    this.serviceTag = source['serviceTag'];
    this.tags = source['tags'];
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceProject {
  spec: ResourceProjectSpec;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.spec = this.convertValues(source['spec'], ResourceProjectSpec);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceClusterOrderStatus {
  status: string;
  phase: string;
  conditions: ResourceKubernetesClusterStatusCondition[];
  createdTime: string;
  updatedTime: string;
  lastObservedTime: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.status = source['status'];
    this.phase = source['phase'];
    this.conditions = this.convertValues(source['conditions'], ResourceKubernetesClusterStatusCondition);
    this.createdTime = source['createdTime'];
    this.updatedTime = source['updatedTime'];
    this.lastObservedTime = source['lastObservedTime'];
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceClusterOrderSpecNodePool {
  name: string;
  machineClass: string;
  count: number;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.name = source['name'];
    this.machineClass = source['machineClass'];
    this.count = source['count'];
  }
}
export class ResourceClusterOrderSpec {
  provider: string;
  clusterName: string;
  projectId: string;
  orderBy: string;
  environment: number;
  criticality: number;
  sensitivity: number;
  highAvailability: boolean;
  nodePools: ResourceClusterOrderSpecNodePool[];
  serviceTags?: { [key: string]: string };
  providerConfig?: { [key: string]: any };
  ownerGroup: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.provider = source['provider'];
    this.clusterName = source['clusterName'];
    this.projectId = source['projectId'];
    this.orderBy = source['orderBy'];
    this.environment = source['environment'];
    this.criticality = source['criticality'];
    this.sensitivity = source['sensitivity'];
    this.highAvailability = source['highAvailability'];
    this.nodePools = this.convertValues(source['nodePools'], ResourceClusterOrderSpecNodePool);
    this.serviceTags = source['serviceTags'];
    this.providerConfig = source['providerConfig'];
    this.ownerGroup = source['ownerGroup'];
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceClusterOrder {
  spec: ResourceClusterOrderSpec;
  status: ResourceClusterOrderStatus;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.spec = this.convertValues(source['spec'], ResourceClusterOrderSpec);
    this.status = this.convertValues(source['status'], ResourceClusterOrderStatus);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceKubernetesClusterStatusCondition {
  type: string;
  status: string;
  lastTransitionTime: string;
  reason: string;
  message: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.type = source['type'];
    this.status = source['status'];
    this.lastTransitionTime = source['lastTransitionTime'];
    this.reason = source['reason'];
    this.message = source['message'];
  }
}
export class ResourceKubernetesClusterStatus {
  status: string;
  phase: string;
  conditions: ResourceKubernetesClusterStatusCondition[];
  kubernetesVersion: string;
  providerStatus: { [key: string]: any };
  createdTime: string;
  updatedTime: string;
  lastObservedTime: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.status = source['status'];
    this.phase = source['phase'];
    this.conditions = this.convertValues(source['conditions'], ResourceKubernetesClusterStatusCondition);
    this.kubernetesVersion = source['kubernetesVersion'];
    this.providerStatus = source['providerStatus'];
    this.createdTime = source['createdTime'];
    this.updatedTime = source['updatedTime'];
    this.lastObservedTime = source['lastObservedTime'];
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceKubernetesClusterSpecEndpoint {
  type: string;
  address: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.type = source['type'];
    this.address = source['address'];
  }
}
export class ResourceKubernetesClusterSpecTopologyWorkers {
  name: string;
  replicas: number;
  version: string;
  machineClass: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.name = source['name'];
    this.replicas = source['replicas'];
    this.version = source['version'];
    this.machineClass = source['machineClass'];
  }
}
export class ResourceKubernetesClusterSpecTopologyControlPlane {
  replicas: number;
  version: string;
  machineClass: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.replicas = source['replicas'];
    this.version = source['version'];
    this.machineClass = source['machineClass'];
  }
}
export class ResourceKubernetesClusterSpecTopology {
  controlPlane: ResourceKubernetesClusterSpecTopologyControlPlane;
  workers: ResourceKubernetesClusterSpecTopologyWorkers[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.controlPlane = this.convertValues(source['controlPlane'], ResourceKubernetesClusterSpecTopologyControlPlane);
    this.workers = this.convertValues(source['workers'], ResourceKubernetesClusterSpecTopologyWorkers);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceKubernetesClusterSpecProviderSpecAzureSpec {
  subscriptionId: string;
  resourceGroup: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.subscriptionId = source['subscriptionId'];
    this.resourceGroup = source['resourceGroup'];
  }
}
export class ResourceKubernetesClusterSpecProviderSpecTanzuSpec {
  supervisorClusterName: string;
  namespace: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.supervisorClusterName = source['supervisorClusterName'];
    this.namespace = source['namespace'];
  }
}
export class ResourceKubernetesClusterSpecProviderSpec {
  tanzuSpec: ResourceKubernetesClusterSpecProviderSpecTanzuSpec;
  azureSpec: ResourceKubernetesClusterSpecProviderSpecAzureSpec;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.tanzuSpec = this.convertValues(source['tanzuSpec'], ResourceKubernetesClusterSpecProviderSpecTanzuSpec);
    this.azureSpec = this.convertValues(source['azureSpec'], ResourceKubernetesClusterSpecProviderSpecAzureSpec);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceKubernetesClusterSpecToolingConfig {
  splunkIndex: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.splunkIndex = source['splunkIndex'];
  }
}
export class ResourceKubernetesClusterSpec {
  clusterId: string;
  clusterName: string;
  description: string;
  project: string;
  provider: string;
  createdBy: string;
  toolingConfig: ResourceKubernetesClusterSpecToolingConfig;
  environment: string;
  providerSpec: ResourceKubernetesClusterSpecProviderSpec;
  topology: ResourceKubernetesClusterSpecTopology;
  endpoints: ResourceKubernetesClusterSpecEndpoint[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.clusterId = source['clusterId'];
    this.clusterName = source['clusterName'];
    this.description = source['description'];
    this.project = source['project'];
    this.provider = source['provider'];
    this.createdBy = source['createdBy'];
    this.toolingConfig = this.convertValues(source['toolingConfig'], ResourceKubernetesClusterSpecToolingConfig);
    this.environment = source['environment'];
    this.providerSpec = this.convertValues(source['providerSpec'], ResourceKubernetesClusterSpecProviderSpec);
    this.topology = this.convertValues(source['topology'], ResourceKubernetesClusterSpecTopology);
    this.endpoints = this.convertValues(source['endpoints'], ResourceKubernetesClusterSpecEndpoint);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceKubernetesCluster {
  spec: ResourceKubernetesClusterSpec;
  status: ResourceKubernetesClusterStatus;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.spec = this.convertValues(source['spec'], ResourceKubernetesClusterSpec);
    this.status = this.convertValues(source['status'], ResourceKubernetesClusterStatus);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceVirtualMachineClassBindingClassRef {
  apiVersion: string;
  kind: string;
  name: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.apiVersion = source['apiVersion'];
    this.kind = source['kind'];
    this.name = source['name'];
  }
}
export class ResourceVirtualMachineClassBinding {
  classRef: ResourceVirtualMachineClassBindingClassRef;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.classRef = this.convertValues(source['classRef'], ResourceVirtualMachineClassBindingClassRef);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceVirtualMachineClassSpecHardwareInstanceStorage {
  storageClass: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.storageClass = source['storageClass'];
  }
}
export class ResourceVirtualMachineClassSpecHardware {
  cpus: number;
  instanceStorage: ResourceVirtualMachineClassSpecHardwareInstanceStorage;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.cpus = source['cpus'];
    this.instanceStorage = this.convertValues(source['instanceStorage'], ResourceVirtualMachineClassSpecHardwareInstanceStorage);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceVirtualMachineClassSpec {
  description: string;
  hardware: ResourceVirtualMachineClassSpecHardware;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.description = source['description'];
    this.hardware = this.convertValues(source['hardware'], ResourceVirtualMachineClassSpecHardware);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceVirtualMachineClass {
  spec: ResourceVirtualMachineClassSpec;
  status: { [key: string]: string };

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.spec = this.convertValues(source['spec'], ResourceVirtualMachineClassSpec);
    this.status = source['status'];
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceTanzuKubernetesReleaseStatusCondition {
  lastTransitionTime: string;
  message: string;
  reason: string;
  severity: string;
  status: string;
  type: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.lastTransitionTime = source['lastTransitionTime'];
    this.message = source['message'];
    this.reason = source['reason'];
    this.severity = source['severity'];
    this.status = source['status'];
    this.type = source['type'];
  }
}
export class ResourceTanzuKubernetesReleaseStatus {
  conditions: ResourceTanzuKubernetesReleaseStatusCondition[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.conditions = this.convertValues(source['conditions'], ResourceTanzuKubernetesReleaseStatusCondition);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceTanzuKubernetesReleaseSpecNodeImageRef {
  apiVersion: string;
  fieldPath: string;
  kind: string;
  namespace: string;
  uid: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.apiVersion = source['apiVersion'];
    this.fieldPath = source['fieldPath'];
    this.kind = source['kind'];
    this.namespace = source['namespace'];
    this.uid = source['uid'];
  }
}
export class ResourceTanzuKubernetesReleaseSpecImage {
  name: string;
  repository: string;
  tag: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.name = source['name'];
    this.repository = source['repository'];
    this.tag = source['tag'];
  }
}
export class ResourceTanzuKubernetesReleaseSpec {
  images: ResourceTanzuKubernetesReleaseSpecImage[];
  kubernetesVersion: string;
  nodeImageRef: ResourceTanzuKubernetesReleaseSpecNodeImageRef;
  repository: string;
  version: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.images = this.convertValues(source['images'], ResourceTanzuKubernetesReleaseSpecImage);
    this.kubernetesVersion = source['kubernetesVersion'];
    this.nodeImageRef = this.convertValues(source['nodeImageRef'], ResourceTanzuKubernetesReleaseSpecNodeImageRef);
    this.repository = source['repository'];
    this.version = source['version'];
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceTanzuKubernetesRelease {
  spec: ResourceTanzuKubernetesReleaseSpec;
  status: ResourceTanzuKubernetesReleaseStatus;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.spec = this.convertValues(source['spec'], ResourceTanzuKubernetesReleaseSpec);
    this.status = this.convertValues(source['status'], ResourceTanzuKubernetesReleaseStatus);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceTanzuKubernetesClusterStatusConditions {
  lastTransitionTime: string;
  message: string;
  reason: string;
  severity: string;
  status: string;
  type: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.lastTransitionTime = source['lastTransitionTime'];
    this.message = source['message'];
    this.reason = source['reason'];
    this.severity = source['severity'];
    this.status = source['status'];
    this.type = source['type'];
  }
}
export class ResourceTanzuKubernetesClusterStatusAPIEndpoints {
  host: string;
  port: number;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.host = source['host'];
    this.port = source['port'];
  }
}
export class ResourceTanzuKubernetesClusterStatus {
  apiEndpoints: ResourceTanzuKubernetesClusterStatusAPIEndpoints[];
  conditions: ResourceTanzuKubernetesClusterStatusConditions[];
  phase: string;
  totalWorkerReplicas: number;
  version: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.apiEndpoints = this.convertValues(source['apiEndpoints'], ResourceTanzuKubernetesClusterStatusAPIEndpoints);
    this.conditions = this.convertValues(source['conditions'], ResourceTanzuKubernetesClusterStatusConditions);
    this.phase = source['phase'];
    this.totalWorkerReplicas = source['totalWorkerReplicas'];
    this.version = source['version'];
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceTanzuKubernetesClusterSpecTopologyNodePoolsVolumes {
  capasity: { [key: string]: string };
  mountPath: string;
  name: string;
  storageClass: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.capasity = source['capasity'];
    this.mountPath = source['mountPath'];
    this.name = source['name'];
    this.storageClass = source['storageClass'];
  }
}
export class ResourceTanzuKubernetesClusterSpecTopologyNodePoolsTkrReference {
  fieldPath: string;
  kind: string;
  name: string;
  namespace: string;
  uid: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.fieldPath = source['fieldPath'];
    this.kind = source['kind'];
    this.name = source['name'];
    this.namespace = source['namespace'];
    this.uid = source['uid'];
  }
}
export class ResourceTanzuKubernetesClusterSpecTopologyNodePoolsTkr {
  reference: ResourceTanzuKubernetesClusterSpecTopologyNodePoolsTkrReference;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.reference = this.convertValues(source['reference'], ResourceTanzuKubernetesClusterSpecTopologyNodePoolsTkrReference);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceTanzuKubernetesClusterSpecTopologyNodePoolsTaints {
  effect: string;
  key: string;
  timeAdded: string;
  value: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.effect = source['effect'];
    this.key = source['key'];
    this.timeAdded = source['timeAdded'];
    this.value = source['value'];
  }
}
export class ResourceTanzuKubernetesClusterSpecTopologyNodePools {
  failureDomain: string;
  labels: { [key: string]: string };
  name: string;
  nodeDrainTimeout: string;
  replicas: number;
  storageClass: string;
  taints: ResourceTanzuKubernetesClusterSpecTopologyNodePoolsTaints[];
  tkr: ResourceTanzuKubernetesClusterSpecTopologyNodePoolsTkr;
  vmClass: string;
  volumes: ResourceTanzuKubernetesClusterSpecTopologyNodePoolsVolumes[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.failureDomain = source['failureDomain'];
    this.labels = source['labels'];
    this.name = source['name'];
    this.nodeDrainTimeout = source['nodeDrainTimeout'];
    this.replicas = source['replicas'];
    this.storageClass = source['storageClass'];
    this.taints = this.convertValues(source['taints'], ResourceTanzuKubernetesClusterSpecTopologyNodePoolsTaints);
    this.tkr = this.convertValues(source['tkr'], ResourceTanzuKubernetesClusterSpecTopologyNodePoolsTkr);
    this.vmClass = source['vmClass'];
    this.volumes = this.convertValues(source['volumes'], ResourceTanzuKubernetesClusterSpecTopologyNodePoolsVolumes);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceTanzuKubernetesClusterSpecTopologyControlPlaneTkrReference {
  name: string;
  kind: string;
  namespace: string;
  uid: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.name = source['name'];
    this.kind = source['kind'];
    this.namespace = source['namespace'];
    this.uid = source['uid'];
  }
}
export class ResourceTanzuKubernetesClusterSpecTopologyControlPlaneTkr {
  reference: ResourceTanzuKubernetesClusterSpecTopologyControlPlaneTkrReference;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.reference = this.convertValues(source['reference'], ResourceTanzuKubernetesClusterSpecTopologyControlPlaneTkrReference);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceTanzuKubernetesClusterSpecTopologyControlPlane {
  nodeDrainTimeout: string;
  replicas: number;
  storageClass: string;
  tkr: ResourceTanzuKubernetesClusterSpecTopologyControlPlaneTkr;
  vmClass: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.nodeDrainTimeout = source['nodeDrainTimeout'];
    this.replicas = source['replicas'];
    this.storageClass = source['storageClass'];
    this.tkr = this.convertValues(source['tkr'], ResourceTanzuKubernetesClusterSpecTopologyControlPlaneTkr);
    this.vmClass = source['vmClass'];
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceTanzuKubernetesClusterSpecTopology {
  controlPlane: ResourceTanzuKubernetesClusterSpecTopologyControlPlane;
  nodePools: ResourceTanzuKubernetesClusterSpecTopologyNodePools[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.controlPlane = this.convertValues(source['controlPlane'], ResourceTanzuKubernetesClusterSpecTopologyControlPlane);
    this.nodePools = this.convertValues(source['nodePools'], ResourceTanzuKubernetesClusterSpecTopologyNodePools);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceTanzuKubernetesClusterSpecSettingsStorage {
  classes: string[];
  defaultClass: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.classes = source['classes'];
    this.defaultClass = source['defaultClass'];
  }
}
export class ResourceTanzuKubernetesClusterSpecSettingsNetworkTrustAdditionalTrustedCA {
  data: string;
  name: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.data = source['data'];
    this.name = source['name'];
  }
}
export class ResourceTanzuKubernetesClusterSpecSettingsNetworkTrust {
  additionalTrustedCAs: ResourceTanzuKubernetesClusterSpecSettingsNetworkTrustAdditionalTrustedCA[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.additionalTrustedCAs = this.convertValues(
      source['additionalTrustedCAs'],
      ResourceTanzuKubernetesClusterSpecSettingsNetworkTrustAdditionalTrustedCA,
    );
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceTanzuKubernetesClusterSpecSettingsNetworkServices {
  cidrBlocks: string[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.cidrBlocks = source['cidrBlocks'];
  }
}
export class ResourceTanzuKubernetesClusterSpecSettingsNetworkProxy {
  httpProxy: string;
  httpsProxy: string;
  noProxy: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.httpProxy = source['httpProxy'];
    this.httpsProxy = source['httpsProxy'];
    this.noProxy = source['noProxy'];
  }
}
export class ResourceTanzuKubernetesClusterSpecSettingsNetworkPods {
  cidrBlocks: string[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.cidrBlocks = source['cidrBlocks'];
  }
}
export class ResourceTanzuKubernetesClusterSpecSettingsNetworkCni {
  name: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.name = source['name'];
  }
}
export class ResourceTanzuKubernetesClusterSpecSettingsNetwork {
  cni: ResourceTanzuKubernetesClusterSpecSettingsNetworkCni;
  pods: ResourceTanzuKubernetesClusterSpecSettingsNetworkPods;
  proxy: ResourceTanzuKubernetesClusterSpecSettingsNetworkProxy;
  serviceDomain: string;
  services: ResourceTanzuKubernetesClusterSpecSettingsNetworkServices;
  trust: ResourceTanzuKubernetesClusterSpecSettingsNetworkTrust;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.cni = this.convertValues(source['cni'], ResourceTanzuKubernetesClusterSpecSettingsNetworkCni);
    this.pods = this.convertValues(source['pods'], ResourceTanzuKubernetesClusterSpecSettingsNetworkPods);
    this.proxy = this.convertValues(source['proxy'], ResourceTanzuKubernetesClusterSpecSettingsNetworkProxy);
    this.serviceDomain = source['serviceDomain'];
    this.services = this.convertValues(source['services'], ResourceTanzuKubernetesClusterSpecSettingsNetworkServices);
    this.trust = this.convertValues(source['trust'], ResourceTanzuKubernetesClusterSpecSettingsNetworkTrust);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceTanzuKubernetesClusterSpecSettings {
  network: ResourceTanzuKubernetesClusterSpecSettingsNetwork;
  storage: ResourceTanzuKubernetesClusterSpecSettingsStorage;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.network = this.convertValues(source['network'], ResourceTanzuKubernetesClusterSpecSettingsNetwork);
    this.storage = this.convertValues(source['storage'], ResourceTanzuKubernetesClusterSpecSettingsStorage);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceTanzuKubernetesClusterSpecDistribution {
  fullVersion: string;
  version: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.fullVersion = source['fullVersion'];
    this.version = source['version'];
  }
}
export class ResourceTanuzKuberntesClusterSpec {
  distribution: ResourceTanzuKubernetesClusterSpecDistribution;
  settings: ResourceTanzuKubernetesClusterSpecSettings;
  topology: ResourceTanzuKubernetesClusterSpecTopology;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.distribution = this.convertValues(source['distribution'], ResourceTanzuKubernetesClusterSpecDistribution);
    this.settings = this.convertValues(source['settings'], ResourceTanzuKubernetesClusterSpecSettings);
    this.topology = this.convertValues(source['topology'], ResourceTanzuKubernetesClusterSpecTopology);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceTanzuKubernetesCluster {
  spec: ResourceTanuzKuberntesClusterSpec;
  status?: ResourceTanzuKubernetesClusterStatus;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.spec = this.convertValues(source['spec'], ResourceTanuzKuberntesClusterSpec);
    this.status = this.convertValues(source['status'], ResourceTanzuKubernetesClusterStatus);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceRbacAssessmentReport {
  report: ResourceVulnerabilityReportReport;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.report = this.convertValues(source['report'], ResourceVulnerabilityReportReport);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceConfigAuditReport {
  report: ResourceVulnerabilityReportReport;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.report = this.convertValues(source['report'], ResourceVulnerabilityReportReport);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceExposedSecretReport {
  report: ResourceVulnerabilityReportReport;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.report = this.convertValues(source['report'], ResourceVulnerabilityReportReport);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceVulnerabilityReportReportVulnerability {
  vulnerabilityID: string;
  severity: string;
  score: number;
  title: string;
  resource: string;
  primaryLink: string;
  installedVersion: string;
  fixedVersion: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.vulnerabilityID = source['vulnerabilityID'];
    this.severity = source['severity'];
    this.score = source['score'];
    this.title = source['title'];
    this.resource = source['resource'];
    this.primaryLink = source['primaryLink'];
    this.installedVersion = source['installedVersion'];
    this.fixedVersion = source['fixedVersion'];
  }
}
export class ResourceVulnerabilityReportReportArtifact {
  repository: string;
  tag: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.repository = source['repository'];
    this.tag = source['tag'];
  }
}
export class AquaReportScanner {
  name: string;
  vendor: string;
  version: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.name = source['name'];
    this.vendor = source['vendor'];
    this.version = source['version'];
  }
}
export class AquaReportSummary {
  criticalCount: number;
  highCount: number;
  lowCount: number;
  mediumCount: number;
  total?: number;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.criticalCount = source['criticalCount'];
    this.highCount = source['highCount'];
    this.lowCount = source['lowCount'];
    this.mediumCount = source['mediumCount'];
    this.total = source['total'];
  }
}
export class ResourceVulnerabilityReportReport {
  summary: AquaReportSummary;
  scanner: AquaReportScanner;
  artifact: ResourceVulnerabilityReportReportArtifact;
  updateTimestamp: string;
  vulnerabilities: ResourceVulnerabilityReportReportVulnerability[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.summary = this.convertValues(source['summary'], AquaReportSummary);
    this.scanner = this.convertValues(source['scanner'], AquaReportScanner);
    this.artifact = this.convertValues(source['artifact'], ResourceVulnerabilityReportReportArtifact);
    this.updateTimestamp = source['updateTimestamp'];
    this.vulnerabilities = this.convertValues(source['vulnerabilities'], ResourceVulnerabilityReportReportVulnerability);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceVulnerabilityReport {
  report: ResourceVulnerabilityReportReport;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.report = this.convertValues(source['report'], ResourceVulnerabilityReportReport);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceIngressClassSpecParameters {
  apiGroup: string;
  kind: string;
  name: string;
  namespace: string;
  scope: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.apiGroup = source['apiGroup'];
    this.kind = source['kind'];
    this.name = source['name'];
    this.namespace = source['namespace'];
    this.scope = source['scope'];
  }
}
export class ResourceIngressClassSpec {
  controller: string;
  parameters: ResourceIngressClassSpecParameters;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.controller = source['controller'];
    this.parameters = this.convertValues(source['parameters'], ResourceIngressClassSpecParameters);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceIngressClass {
  spec: ResourceIngressClassSpec;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.spec = this.convertValues(source['spec'], ResourceIngressClassSpec);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceIngressStatusLoadBalancerIngress {
  hostname: string;
  ip: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.hostname = source['hostname'];
    this.ip = source['ip'];
  }
}
export class ResourceIngressStatusLoadBalancer {
  ingress: ResourceIngressStatusLoadBalancerIngress[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.ingress = this.convertValues(source['ingress'], ResourceIngressStatusLoadBalancerIngress);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceIngressStatus {
  loadBalancer: ResourceIngressStatusLoadBalancer;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.loadBalancer = this.convertValues(source['loadBalancer'], ResourceIngressStatusLoadBalancer);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceIngressSpecTls {
  hosts: string[];
  secretName: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.hosts = source['hosts'];
    this.secretName = source['secretName'];
  }
}
export class ResourceIngressSpecRulesHttpPaths {
  backend: ResourceIngressSpecRulesHttpPathsBackend;
  path: string;
  pathType: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.backend = this.convertValues(source['backend'], ResourceIngressSpecRulesHttpPathsBackend);
    this.path = source['path'];
    this.pathType = source['pathType'];
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceIngressSpecRulesHttp {
  paths: ResourceIngressSpecRulesHttpPaths[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.paths = this.convertValues(source['paths'], ResourceIngressSpecRulesHttpPaths);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceIngressSpecRules {
  apiGroup: string;
  http: ResourceIngressSpecRulesHttp;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.apiGroup = source['apiGroup'];
    this.http = this.convertValues(source['http'], ResourceIngressSpecRulesHttp);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceIngressSpecBackendServicePort {
  name?: string;
  number?: number;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.name = source['name'];
    this.number = source['number'];
  }
}
export class ResourceIngressSpecBackendService {
  name?: string;
  port?: ResourceIngressSpecBackendServicePort;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.name = source['name'];
    this.port = this.convertValues(source['port'], ResourceIngressSpecBackendServicePort);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceIngressSpecBackendResource {
  apiGroup?: string;
  kind?: string;
  name?: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.apiGroup = source['apiGroup'];
    this.kind = source['kind'];
    this.name = source['name'];
  }
}
export class ResourceIngressSpecRulesHttpPathsBackend {
  resource?: ResourceIngressSpecBackendResource;
  service?: ResourceIngressSpecBackendService;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.resource = this.convertValues(source['resource'], ResourceIngressSpecBackendResource);
    this.service = this.convertValues(source['service'], ResourceIngressSpecBackendService);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceIngressSpec {
  defaultBackend?: ResourceIngressSpecRulesHttpPathsBackend;
  ingressClassName: string;
  rules: ResourceIngressSpecRules[];
  tls: ResourceIngressSpecTls[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.defaultBackend = this.convertValues(source['defaultBackend'], ResourceIngressSpecRulesHttpPathsBackend);
    this.ingressClassName = source['ingressClassName'];
    this.rules = this.convertValues(source['rules'], ResourceIngressSpecRules);
    this.tls = this.convertValues(source['tls'], ResourceIngressSpecTls);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceIngress {
  spec: ResourceIngressSpec;
  status: ResourceIngressStatus;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.spec = this.convertValues(source['spec'], ResourceIngressSpec);
    this.status = this.convertValues(source['status'], ResourceIngressStatus);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceDaemonSetStatus {
  numberReady: number;
  numberUnavailable: number;
  currentReplicas: number;
  numberAvailable: number;
  updatedNumberScheduled: number;
  desiredNumberScheduled: number;
  currentNumberScheduled: number;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.numberReady = source['numberReady'];
    this.numberUnavailable = source['numberUnavailable'];
    this.currentReplicas = source['currentReplicas'];
    this.numberAvailable = source['numberAvailable'];
    this.updatedNumberScheduled = source['updatedNumberScheduled'];
    this.desiredNumberScheduled = source['desiredNumberScheduled'];
    this.currentNumberScheduled = source['currentNumberScheduled'];
  }
}
export class ResourceDaemonSet {
  status: ResourceDaemonSetStatus;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.status = this.convertValues(source['status'], ResourceDaemonSetStatus);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceStatefulSetStatus {
  replicas: number;
  availableReplicas: number;
  currentReplicas: number;
  readyReplicas: number;
  updatedReplicas: number;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.replicas = source['replicas'];
    this.availableReplicas = source['availableReplicas'];
    this.currentReplicas = source['currentReplicas'];
    this.readyReplicas = source['readyReplicas'];
    this.updatedReplicas = source['updatedReplicas'];
  }
}
export class ResourceStatefulSet {
  status: ResourceStatefulSetStatus;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.status = this.convertValues(source['status'], ResourceStatefulSetStatus);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceReplicaSetStatus {
  availableReplicas: number;
  readyReplicas: number;
  replicas: number;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.availableReplicas = source['availableReplicas'];
    this.readyReplicas = source['readyReplicas'];
    this.replicas = source['replicas'];
  }
}
export class ResourceReplicaSetSpecSelectorMatchExpressions {
  key: string;
  operator: string;
  values: string[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.key = source['key'];
    this.operator = source['operator'];
    this.values = source['values'];
  }
}
export class ResourceReplicaSetSpecSelector {
  matchExpressions: ResourceReplicaSetSpecSelectorMatchExpressions[];
  matchLabels: { [key: string]: string };

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.matchExpressions = this.convertValues(source['matchExpressions'], ResourceReplicaSetSpecSelectorMatchExpressions);
    this.matchLabels = source['matchLabels'];
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceReplicaSetSpec {
  replicas: number;
  selector: ResourceReplicaSetSpecSelector;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.replicas = source['replicas'];
    this.selector = this.convertValues(source['selector'], ResourceReplicaSetSpecSelector);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceReplicaSet {
  spec: ResourceReplicaSetSpec;
  status: ResourceReplicaSetStatus;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.spec = this.convertValues(source['spec'], ResourceReplicaSetSpec);
    this.status = this.convertValues(source['status'], ResourceReplicaSetStatus);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourcePodStatus {
  message?: string;
  phase?: string;
  reason?: string;
  startTime?: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.message = source['message'];
    this.phase = source['phase'];
    this.reason = source['reason'];
    this.startTime = source['startTime'];
  }
}
export class ResourcePodSpecContainersPorts {
  name?: string;
  containerPort?: number;
  protocol?: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.name = source['name'];
    this.containerPort = source['containerPort'];
    this.protocol = source['protocol'];
  }
}
export class ResourcePodSpecContainers {
  name?: string;
  image?: string;
  ports?: ResourcePodSpecContainersPorts[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.name = source['name'];
    this.image = source['image'];
    this.ports = this.convertValues(source['ports'], ResourcePodSpecContainersPorts);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourcePodSpec {
  containers?: ResourcePodSpecContainers[];
  serviceAccountName?: string;
  nodeName?: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.containers = this.convertValues(source['containers'], ResourcePodSpecContainers);
    this.serviceAccountName = source['serviceAccountName'];
    this.nodeName = source['nodeName'];
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourcePod {
  spec?: ResourcePodSpec;
  status?: ResourcePodStatus;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.spec = this.convertValues(source['spec'], ResourcePodSpec);
    this.status = this.convertValues(source['status'], ResourcePodStatus);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class IntOrString {
  Type: number;
  IntVal: number;
  StrVal: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.Type = source['Type'];
    this.IntVal = source['IntVal'];
    this.StrVal = source['StrVal'];
  }
}
export class ResourceServicePorts {
  appProtocol: string;
  name: string;
  port: number;
  protocol: string;
  targetPort: IntOrString;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.appProtocol = source['appProtocol'];
    this.name = source['name'];
    this.port = source['port'];
    this.protocol = source['protocol'];
    this.targetPort = this.convertValues(source['targetPort'], IntOrString);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceServiceSpec {
  type: string;
  selector: { [key: string]: string };
  ports: ResourceServicePorts[];
  clusterIP: string;
  clusterIPs: string[];
  externalIPs?: string[];
  externalName?: string;
  ipFamilies: string[];
  ipFamilyPolicy: string;
  internalTrafficPolicy: string;
  externalTrafficPolicy: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.type = source['type'];
    this.selector = source['selector'];
    this.ports = this.convertValues(source['ports'], ResourceServicePorts);
    this.clusterIP = source['clusterIP'];
    this.clusterIPs = source['clusterIPs'];
    this.externalIPs = source['externalIPs'];
    this.externalName = source['externalName'];
    this.ipFamilies = source['ipFamilies'];
    this.ipFamilyPolicy = source['ipFamilyPolicy'];
    this.internalTrafficPolicy = source['internalTrafficPolicy'];
    this.externalTrafficPolicy = source['externalTrafficPolicy'];
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceService {
  spec: ResourceServiceSpec;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.spec = this.convertValues(source['spec'], ResourceServiceSpec);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceCertificateStatusCondition {
  lastTransitionTime: string;
  observedGeneration: number;
  message: string;
  reason: string;
  status: string;
  type: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.lastTransitionTime = source['lastTransitionTime'];
    this.observedGeneration = source['observedGeneration'];
    this.message = source['message'];
    this.reason = source['reason'];
    this.status = source['status'];
    this.type = source['type'];
  }
}
export class ResourceCertificateStatus {
  notBefore: string;
  notAfter: string;
  renewalTime: string;
  conditions: ResourceCertificateStatusCondition[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.notBefore = source['notBefore'];
    this.notAfter = source['notAfter'];
    this.renewalTime = source['renewalTime'];
    this.conditions = this.convertValues(source['conditions'], ResourceCertificateStatusCondition);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceCertificateSpecIssuerref {
  group: string;
  kind: string;
  name: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.group = source['group'];
    this.kind = source['kind'];
    this.name = source['name'];
  }
}
export class ResourceCertificateSpec {
  dnsNames: string[];
  secretName: string;
  issuerRef: ResourceCertificateSpecIssuerref;
  usages?: string[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.dnsNames = source['dnsNames'];
    this.secretName = source['secretName'];
    this.issuerRef = this.convertValues(source['issuerRef'], ResourceCertificateSpecIssuerref);
    this.usages = source['usages'];
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceCertificate {
  spec: ResourceCertificateSpec;
  status: ResourceCertificateStatus;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.spec = this.convertValues(source['spec'], ResourceCertificateSpec);
    this.status = this.convertValues(source['status'], ResourceCertificateStatus);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceAppProjectSpec {
  description: string;
  sourceRepos: string[];
  destinations: ResourceApplicationSpecDestination[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.description = source['description'];
    this.sourceRepos = source['sourceRepos'];
    this.destinations = this.convertValues(source['destinations'], ResourceApplicationSpecDestination);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceAppProject {
  spec: ResourceAppProjectSpec;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.spec = this.convertValues(source['spec'], ResourceAppProjectSpec);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceApplicationStatusSync {
  revision: string;
  status: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.revision = source['revision'];
    this.status = source['status'];
  }
}
export class ResourceApplicationStatusHealth {
  message: string;
  status: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.message = source['message'];
    this.status = source['status'];
  }
}
export class ResourceApplicationStatusOperationstate {
  startedAt: string;
  finishedAt: string;
  phase: string;
  status: string;
  message: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.startedAt = source['startedAt'];
    this.finishedAt = source['finishedAt'];
    this.phase = source['phase'];
    this.status = source['status'];
    this.message = source['message'];
  }
}
export class ResourceApplicationStatus {
  sourceType: string;
  reconciledAt: string;
  operationState: ResourceApplicationStatusOperationstate;
  health: ResourceApplicationStatusHealth;
  sync: ResourceApplicationStatusSync;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.sourceType = source['sourceType'];
    this.reconciledAt = source['reconciledAt'];
    this.operationState = this.convertValues(source['operationState'], ResourceApplicationStatusOperationstate);
    this.health = this.convertValues(source['health'], ResourceApplicationStatusHealth);
    this.sync = this.convertValues(source['sync'], ResourceApplicationStatusSync);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceApplicationSpecSyncpolicyRetryBackoff {
  duration: string;
  factor: number;
  maxDuration: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.duration = source['duration'];
    this.factor = source['factor'];
    this.maxDuration = source['maxDuration'];
  }
}
export class ResourceApplicationSpecSyncpolicyRetry {
  backoff?: ResourceApplicationSpecSyncpolicyRetryBackoff;
  limit: number;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.backoff = this.convertValues(source['backoff'], ResourceApplicationSpecSyncpolicyRetryBackoff);
    this.limit = source['limit'];
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceApplicationSpecSyncpolicyAutomated {
  allowEmpty: boolean;
  prune: boolean;
  selfHeal: boolean;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.allowEmpty = source['allowEmpty'];
    this.prune = source['prune'];
    this.selfHeal = source['selfHeal'];
  }
}
export class ResourceApplicationSpecSyncpolicy {
  automated?: ResourceApplicationSpecSyncpolicyAutomated;
  retry?: ResourceApplicationSpecSyncpolicyRetry;
  syncOptions: string[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.automated = this.convertValues(source['automated'], ResourceApplicationSpecSyncpolicyAutomated);
    this.retry = this.convertValues(source['retry'], ResourceApplicationSpecSyncpolicyRetry);
    this.syncOptions = source['syncOptions'];
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceApplicationSpecSource {
  chart?: string;
  path?: string;
  repoURL: string;
  targetRevision: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.chart = source['chart'];
    this.path = source['path'];
    this.repoURL = source['repoURL'];
    this.targetRevision = source['targetRevision'];
  }
}
export class ResourceApplicationSpecDestination {
  name: string;
  namespace: string;
  server: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.name = source['name'];
    this.namespace = source['namespace'];
    this.server = source['server'];
  }
}
export class ResourceApplicationSpec {
  destination: ResourceApplicationSpecDestination;
  project: string;
  source: ResourceApplicationSpecSource;
  syncPolicy: ResourceApplicationSpecSyncpolicy;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.destination = this.convertValues(source['destination'], ResourceApplicationSpecDestination);
    this.project = source['project'];
    this.source = this.convertValues(source['source'], ResourceApplicationSpecSource);
    this.syncPolicy = this.convertValues(source['syncPolicy'], ResourceApplicationSpecSyncpolicy);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceApplication {
  spec: ResourceApplicationSpec;
  status: ResourceApplicationStatus;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.spec = this.convertValues(source['spec'], ResourceApplicationSpec);
    this.status = this.convertValues(source['status'], ResourceApplicationStatus);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourcePolicyReportSummary {
  error: number;
  fail: number;
  pass: number;
  skip: number;
  warn: number;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.error = source['error'];
    this.fail = source['fail'];
    this.pass = source['pass'];
    this.skip = source['skip'];
    this.warn = source['warn'];
  }
}
export class ResourcePolicyReportResultsResources {
  uid: string;
  apiVersion: string;
  kind: string;
  name: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.uid = source['uid'];
    this.apiVersion = source['apiVersion'];
    this.kind = source['kind'];
    this.name = source['name'];
  }
}
export class ResourcePolicyReportResults {
  policy: string;
  message: string;
  category: string;
  properties: { [key: string]: string };
  severity: string;
  result: string;
  resources: ResourcePolicyReportResultsResources[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.policy = source['policy'];
    this.message = source['message'];
    this.category = source['category'];
    this.properties = source['properties'];
    this.severity = source['severity'];
    this.result = source['result'];
    this.resources = this.convertValues(source['resources'], ResourcePolicyReportResultsResources);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourcePolicyReport {
  results: ResourcePolicyReportResults[];
  summary: ResourcePolicyReportSummary;
  lastReported?: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.results = this.convertValues(source['results'], ResourcePolicyReportResults);
    this.summary = this.convertValues(source['summary'], ResourcePolicyReportSummary);
    this.lastReported = source['lastReported'];
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceStorageClass {
  allowVolumeExpansion: boolean;
  provisioner: string;
  reclaimPolicy: string;
  volumeBindingMode: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.allowVolumeExpansion = source['allowVolumeExpansion'];
    this.provisioner = source['provisioner'];
    this.reclaimPolicy = source['reclaimPolicy'];
    this.volumeBindingMode = source['volumeBindingMode'];
  }
}
export class ResourceDeploymentStatus {
  replicas: number;
  availableReplicas: number;
  readyReplicas: number;
  updatedReplicas: number;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.replicas = source['replicas'];
    this.availableReplicas = source['availableReplicas'];
    this.readyReplicas = source['readyReplicas'];
    this.updatedReplicas = source['updatedReplicas'];
  }
}
export class ResourceDeployment {
  status: ResourceDeploymentStatus;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.status = this.convertValues(source['status'], ResourceDeploymentStatus);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourcePersistentVolumeClaimStatus {
  accessModes: string[];
  capacity: { [key: string]: string };
  phase: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.accessModes = source['accessModes'];
    this.capacity = source['capacity'];
    this.phase = source['phase'];
  }
}
export class ResourcePersistentVolumeClaimSpecResources {
  limits?: { [key: string]: string };
  requests: { [key: string]: string };

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.limits = source['limits'];
    this.requests = source['requests'];
  }
}
export class ResourcePersistentVolumeClaimSpec {
  accessModes: string[];
  resources: ResourcePersistentVolumeClaimSpecResources;
  storageClassName: string;
  volumeMode: string;
  volumeName: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.accessModes = source['accessModes'];
    this.resources = this.convertValues(source['resources'], ResourcePersistentVolumeClaimSpecResources);
    this.storageClassName = source['storageClassName'];
    this.volumeMode = source['volumeMode'];
    this.volumeName = source['volumeName'];
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourcePersistentVolumeClaim {
  spec: ResourcePersistentVolumeClaimSpec;
  status: ResourcePersistentVolumeClaimStatus;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.spec = this.convertValues(source['spec'], ResourcePersistentVolumeClaimSpec);
    this.status = this.convertValues(source['status'], ResourcePersistentVolumeClaimStatus);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceNodeStatusNodeinfo {
  architecture: string;
  bootID: string;
  containerRuntimeVersion: string;
  kernelVersion: string;
  kubeProxyVersion: string;
  kubeletVersion: string;
  machineID: string;
  operatingSystem: string;
  osImage: string;
  systemUUID: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.architecture = source['architecture'];
    this.bootID = source['bootID'];
    this.containerRuntimeVersion = source['containerRuntimeVersion'];
    this.kernelVersion = source['kernelVersion'];
    this.kubeProxyVersion = source['kubeProxyVersion'];
    this.kubeletVersion = source['kubeletVersion'];
    this.machineID = source['machineID'];
    this.operatingSystem = source['operatingSystem'];
    this.osImage = source['osImage'];
    this.systemUUID = source['systemUUID'];
  }
}
export class ResourceNodeStatusConditions {
  lastHeartbeatTime: string;
  lastTransitionTime: string;
  message: string;
  reason: string;
  status: string;
  type: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.lastHeartbeatTime = source['lastHeartbeatTime'];
    this.lastTransitionTime = source['lastTransitionTime'];
    this.message = source['message'];
    this.reason = source['reason'];
    this.status = source['status'];
    this.type = source['type'];
  }
}
export class ResourceNodeStatusCapacity {
  cpu: string;
  ephemeralStorage: string;
  memory: string;
  pods: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.cpu = source['cpu'];
    this.ephemeralStorage = source['ephemeral-storage'];
    this.memory = source['memory'];
    this.pods = source['pods'];
  }
}
export class ResourceNodeStatusAddresses {
  address: string;
  type: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.address = source['address'];
    this.type = source['type'];
  }
}
export class ResourceNodeStatus {
  addresses: ResourceNodeStatusAddresses[];
  capacity: ResourceNodeStatusCapacity;
  conditions: ResourceNodeStatusConditions[];
  nodeInfo: ResourceNodeStatusNodeinfo;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.addresses = this.convertValues(source['addresses'], ResourceNodeStatusAddresses);
    this.capacity = this.convertValues(source['capacity'], ResourceNodeStatusCapacity);
    this.conditions = this.convertValues(source['conditions'], ResourceNodeStatusConditions);
    this.nodeInfo = this.convertValues(source['nodeInfo'], ResourceNodeStatusNodeinfo);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceNodeSpecTaints {
  effect: string;
  key: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.effect = source['effect'];
    this.key = source['key'];
  }
}
export class ResourceNodeSpec {
  podCIDR?: string;
  podCIDRs?: string[];
  providerID?: string;
  taints?: ResourceNodeSpecTaints[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.podCIDR = source['podCIDR'];
    this.podCIDRs = source['podCIDRs'];
    this.providerID = source['providerID'];
    this.taints = this.convertValues(source['taints'], ResourceNodeSpecTaints);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceNode {
  spec: ResourceNodeSpec;
  status: ResourceNodeStatus;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.spec = this.convertValues(source['spec'], ResourceNodeSpec);
    this.status = this.convertValues(source['status'], ResourceNodeStatus);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceNamespace {
  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
  }
}
export class RorResourceOwnerReference {
  scope: string;
  subject: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.scope = source['scope'];
    this.subject = source['subject'];
  }
}
export class ResourceRorMeta {
  version?: string;
  lastReported?: string;
  internal?: boolean;
  hash?: string;
  ownerref?: RorResourceOwnerReference;
  action?: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.version = source['version'];
    this.lastReported = source['lastReported'];
    this.internal = source['internal'];
    this.hash = source['hash'];
    this.ownerref = this.convertValues(source['ownerref'], RorResourceOwnerReference);
    this.action = source['action'];
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class FieldsV1 {
  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
  }
}
export class ManagedFieldsEntry {
  manager?: string;
  operation?: string;
  apiVersion?: string;
  time?: Time;
  fieldsType?: string;
  fieldsV1?: FieldsV1;
  subresource?: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.manager = source['manager'];
    this.operation = source['operation'];
    this.apiVersion = source['apiVersion'];
    this.time = this.convertValues(source['time'], Time);
    this.fieldsType = source['fieldsType'];
    this.fieldsV1 = this.convertValues(source['fieldsV1'], FieldsV1);
    this.subresource = source['subresource'];
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class OwnerReference {
  apiVersion: string;
  kind: string;
  name: string;
  uid: string;
  controller?: boolean;
  blockOwnerDeletion?: boolean;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.apiVersion = source['apiVersion'];
    this.kind = source['kind'];
    this.name = source['name'];
    this.uid = source['uid'];
    this.controller = source['controller'];
    this.blockOwnerDeletion = source['blockOwnerDeletion'];
  }
}

export class ObjectMeta {
  name?: string;
  generateName?: string;
  namespace?: string;
  selfLink?: string;
  uid?: string;
  resourceVersion?: string;
  generation?: number;
  creationTimestamp?: Time;
  deletionTimestamp?: Time;
  deletionGracePeriodSeconds?: number;
  labels?: { [key: string]: string };
  annotations?: { [key: string]: string };
  ownerReferences?: OwnerReference[];
  finalizers?: string[];
  managedFields?: ManagedFieldsEntry[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.name = source['name'];
    this.generateName = source['generateName'];
    this.namespace = source['namespace'];
    this.selfLink = source['selfLink'];
    this.uid = source['uid'];
    this.resourceVersion = source['resourceVersion'];
    this.generation = source['generation'];
    this.creationTimestamp = this.convertValues(source['creationTimestamp'], Time);
    this.deletionTimestamp = this.convertValues(source['deletionTimestamp'], Time);
    this.deletionGracePeriodSeconds = source['deletionGracePeriodSeconds'];
    this.labels = source['labels'];
    this.annotations = source['annotations'];
    this.ownerReferences = this.convertValues(source['ownerReferences'], OwnerReference);
    this.finalizers = source['finalizers'];
    this.managedFields = this.convertValues(source['managedFields'], ManagedFieldsEntry);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class Resource {
  kind?: string;
  apiVersion?: string;
  metadata?: ObjectMeta;
  rormeta: ResourceRorMeta;
  namespace?: ResourceNamespace;
  node?: ResourceNode;
  persistentvolumeclaim?: ResourcePersistentVolumeClaim;
  deployment?: ResourceDeployment;
  storageclass?: ResourceStorageClass;
  policyreport?: ResourcePolicyReport;
  application?: ResourceApplication;
  appproject?: ResourceAppProject;
  certificate?: ResourceCertificate;
  service?: ResourceService;
  pod?: ResourcePod;
  replicaset?: ResourceReplicaSet;
  statefulset?: ResourceStatefulSet;
  daemonset?: ResourceDaemonSet;
  ingress?: ResourceIngress;
  ingressclass?: ResourceIngressClass;
  vulnerabilityreport?: ResourceVulnerabilityReport;
  exposedsecretreport?: ResourceExposedSecretReport;
  configauditreport?: ResourceConfigAuditReport;
  rbacassessmentreport?: ResourceRbacAssessmentReport;
  tanzukubernetescluster?: ResourceTanzuKubernetesCluster;
  tanzukubernetesrelease?: ResourceTanzuKubernetesRelease;
  virtualmachineclass?: ResourceVirtualMachineClass;
  virtualmachineclassbinding?: ResourceVirtualMachineClassBinding;
  kubernetescluster?: ResourceKubernetesCluster;
  clusterorder?: ResourceClusterOrder;
  project?: ResourceProject;
  configuration?: ResourceConfiguration;
  clustercompliancereport?: ResourceClusterComplianceReport;
  clustervulnerabilityreport?: ResourceClusterVulnerabilityReport;
  route?: ResourceRoute;
  slackmessage?: ResourceSlackMessage;
  notification?: ResourceNotification;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.kind = source['kind'];
    this.apiVersion = source['apiVersion'];
    this.metadata = this.convertValues(source['metadata'], ObjectMeta);
    this.rormeta = this.convertValues(source['rormeta'], ResourceRorMeta);
    this.namespace = this.convertValues(source['namespace'], ResourceNamespace);
    this.node = this.convertValues(source['node'], ResourceNode);
    this.persistentvolumeclaim = this.convertValues(source['persistentvolumeclaim'], ResourcePersistentVolumeClaim);
    this.deployment = this.convertValues(source['deployment'], ResourceDeployment);
    this.storageclass = this.convertValues(source['storageclass'], ResourceStorageClass);
    this.policyreport = this.convertValues(source['policyreport'], ResourcePolicyReport);
    this.application = this.convertValues(source['application'], ResourceApplication);
    this.appproject = this.convertValues(source['appproject'], ResourceAppProject);
    this.certificate = this.convertValues(source['certificate'], ResourceCertificate);
    this.service = this.convertValues(source['service'], ResourceService);
    this.pod = this.convertValues(source['pod'], ResourcePod);
    this.replicaset = this.convertValues(source['replicaset'], ResourceReplicaSet);
    this.statefulset = this.convertValues(source['statefulset'], ResourceStatefulSet);
    this.daemonset = this.convertValues(source['daemonset'], ResourceDaemonSet);
    this.ingress = this.convertValues(source['ingress'], ResourceIngress);
    this.ingressclass = this.convertValues(source['ingressclass'], ResourceIngressClass);
    this.vulnerabilityreport = this.convertValues(source['vulnerabilityreport'], ResourceVulnerabilityReport);
    this.exposedsecretreport = this.convertValues(source['exposedsecretreport'], ResourceExposedSecretReport);
    this.configauditreport = this.convertValues(source['configauditreport'], ResourceConfigAuditReport);
    this.rbacassessmentreport = this.convertValues(source['rbacassessmentreport'], ResourceRbacAssessmentReport);
    this.tanzukubernetescluster = this.convertValues(source['tanzukubernetescluster'], ResourceTanzuKubernetesCluster);
    this.tanzukubernetesrelease = this.convertValues(source['tanzukubernetesrelease'], ResourceTanzuKubernetesRelease);
    this.virtualmachineclass = this.convertValues(source['virtualmachineclass'], ResourceVirtualMachineClass);
    this.virtualmachineclassbinding = this.convertValues(source['virtualmachineclassbinding'], ResourceVirtualMachineClassBinding);
    this.kubernetescluster = this.convertValues(source['kubernetescluster'], ResourceKubernetesCluster);
    this.clusterorder = this.convertValues(source['clusterorder'], ResourceClusterOrder);
    this.project = this.convertValues(source['project'], ResourceProject);
    this.configuration = this.convertValues(source['configuration'], ResourceConfiguration);
    this.clustercompliancereport = this.convertValues(source['clustercompliancereport'], ResourceClusterComplianceReport);
    this.clustervulnerabilityreport = this.convertValues(source['clustervulnerabilityreport'], ResourceClusterVulnerabilityReport);
    this.route = this.convertValues(source['route'], ResourceRoute);
    this.slackmessage = this.convertValues(source['slackmessage'], ResourceSlackMessage);
    this.notification = this.convertValues(source['notification'], ResourceNotification);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
export class ResourceSet {
  resources?: Resource[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.resources = this.convertValues(source['resources'], Resource);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
