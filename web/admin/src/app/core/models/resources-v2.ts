/* Do not change, this code is generated from Golang structs */

export enum VulnerabilityStatus {
  NOT_ASSESSED = 0,
  NEEDS_TRIAGE = 1,
  CONFIRMED = 2,
  DISMISSED = 3,
}
export enum VulnerabilityDismissalReason {
  ACCEPTABLE_RISK = 0,
  FALSE_POSITIVE = 1,
  NOT_APPLICABLE = 2,
}
export interface ResourceNetworkPolicyCondition {
  lastTransitionTime: string;
  message: string;
  observedGeneration: number;
  reason: string;
  status: string;
  type: string;
}
export interface ResourceNetworkPolicyStatus {
  conditions: ResourceNetworkPolicyCondition[];
}
export interface ResourceNetworkPolicyPodSelector {
  matchLabels: { [key: string]: string };
}
export interface ResourceNetworkPolicyIngressRule {
  from: ResourceNetworkPolicyPeer[];
  ports: ResourceNetworkPolicyPort[];
}
export interface ResourceNetworkPolicySelectorExpression {
  key: string;
  operator: string;
  values: string[];
}
export interface ResourceNetworkPolicySelector {
  matchExpressions: ResourceNetworkPolicySelectorExpression[];
  matchLabels: { [key: string]: string };
}
export interface ResourceNetworkPolicyIpBlock {
  cidr: string;
  except: string[];
}
export interface ResourceNetworkPolicyPeer {
  ipBlock?: ResourceNetworkPolicyIpBlock;
  namespaceSelector?: ResourceNetworkPolicySelector;
  podSelector?: ResourceNetworkPolicySelector;
}
export interface ResourceNetworkPolicyPort {
  endPort: number;
  port: IntOrString;
  protocol: string;
}
export interface ResourceNetworkPolicyEgressRule {
  ports: ResourceNetworkPolicyPort[];
  to: ResourceNetworkPolicyPeer[];
}
export interface ResourceNetworkPolicySpec {
  egress: ResourceNetworkPolicyEgressRule[];
  ingress: ResourceNetworkPolicyIngressRule[];
  podSelector: ResourceNetworkPolicyPodSelector;
  policyTypes: string[];
}
export interface ResourceNetworkPolicy {
  spec: ResourceNetworkPolicySpec;
  status: ResourceNetworkPolicyStatus;
}
export interface ResourceEndpointSpecSubsetsPorts {
  appProtocol?: string;
  name?: string;
  port?: number;
  protocol?: string;
}
export interface ResourceEndpointSpecSubsetsNotReadyAddressesTargetRef {
  apiVersion?: string;
  fieldPath?: string;
  kind?: string;
  name?: string;
  namespace?: string;
  resourceVersion?: string;
  uid?: string;
}
export interface ResourceEndpointSpecSubsetsNotReadyAddresses {
  hostname?: string;
  ip?: string;
  nodeName?: string;
  targetRef?: ResourceEndpointSpecSubsetsNotReadyAddressesTargetRef;
}
export interface ResourceEndpointSpecSubsetsAddressesTargetRef {
  apiVersion?: string;
  fieldPath?: string;
  kind?: string;
  name?: string;
  namespace?: string;
  resourceVersion?: string;
  uid?: string;
}
export interface ResourceEndpointSpecSubsetsAddresses {
  hostname?: string;
  ip?: string;
  nodeName?: string;
  targetRef?: ResourceEndpointSpecSubsetsAddressesTargetRef;
}
export interface ResourceEndpointSpecSubsets {
  addresses?: ResourceEndpointSpecSubsetsAddresses[];
  notReadyAddresses?: ResourceEndpointSpecSubsetsNotReadyAddresses[];
  ports?: ResourceEndpointSpecSubsetsPorts[];
}
export interface ResourceEndpoints {
  subsets?: ResourceEndpointSpecSubsets[];
}
export interface ResourceVirtualMachineOperatingSystemStatus {
  id: string;
  name: string;
  version: string;
  hostName: string;
  powerState: string;
  toolVersion: string;
  architecture: string;
}
export interface ResourceVirtualMachineNetworkStatus {
  id: string;
}
export interface ResourceVirtualMachineMemoryStatus {
  id: string;
  unit: string;
  usage: string;
}
export interface ResourceVirtualMachineDiskStatus {
  id: string;
  usageBytes: string;
}
export interface ResourceVirtualMachineCpuStatus {
  id: string;
  unit: string;
  usage: string;
}
export interface ResourceVirtualMachineStatus {
  cpu: ResourceVirtualMachineCpuStatus;
  disks: ResourceVirtualMachineDiskStatus[];
  memory: ResourceVirtualMachineMemoryStatus;
  networks: ResourceVirtualMachineNetworkStatus[];
  operatingSystem: ResourceVirtualMachineOperatingSystemStatus;
}
export interface ResourceVirtualMachineOperatingSystemSpec {
  id: string;
}
export interface ResourceVirtualMachineNetworkSpec {
  id: string;
  dns: string;
  ipv4: string;
  ipv6: string;
  mask: string;
  gateway: string;
}
export interface ResourceVirtualMachineMemorySpec {
  id: string;
  sizeBytes: number;
}
export interface ResourceVirtualMachineDiskSpec {
  id: string;
  name: string;
  type: string;
  sizeBytes: number;
}
export interface ResourceVirtualMachineTagSpec {
  key: string;
  value: string;
  description: string;
}
export interface ResourceVirtualMachineCpuSpec {
  id: string;
  sockets: number;
  coresPerSocket: number;
}
export interface ResourceVirtualMachineSpec {
  cpu: ResourceVirtualMachineCpuSpec;
  tags: ResourceVirtualMachineTagSpec[];
  disks: ResourceVirtualMachineDiskSpec[];
  memory: ResourceVirtualMachineMemorySpec;
  networks: ResourceVirtualMachineNetworkSpec[];
  operatingSystem: ResourceVirtualMachineOperatingSystemSpec;
}
export interface ResourceVirtualMachine {
  id: string;
  name: string;
  spec: ResourceVirtualMachineSpec;
  status: ResourceVirtualMachineStatus;
}
export interface ResourceVulnerabilityEventSpec {
  owner: RorResourceOwnerReference;
  message: string;
}
export interface ResourceVulnerabilityEvent {
  spec: ResourceVulnerabilityEventSpec;
}
export interface ResourceSlackMessageStatus {
  result: number;
  timestamp: Time;
  error: any;
}
export interface ResourceSlackMessageSpec {
  channelId: string;
  message: string;
}
export interface ResourceSlackMessage {
  spec: ResourceSlackMessageSpec;
  status: ResourceSlackMessageStatus[];
}
export interface ResourceRouteSlackReceiver {
  channelId: string;
}
export interface ResourceRouteReceiver {
  slack: ResourceRouteSlackReceiver[];
}
export interface ResourceRouteMessageType {
  apiVersion: string;
  kind: string;
}
export interface ResourceRouteSpec {
  messageType: ResourceRouteMessageType;
  receivers: ResourceRouteReceiver;
}
export interface ResourceRoute {
  spec: ResourceRouteSpec;
}
export interface ResourceClusterVulnerabilityReportReportStatus {
  status: VulnerabilityStatus;
  until?: Time;
  reason?: VulnerabilityDismissalReason;
  comment?: string;
  riskAssessment?: string;
}
export interface ResourceClusterVulnerabilityReportReportOwner {
  digest: string;
  repository: string;
  tag: string;
  resource: string;
  installedVersion: string;
  fixedVersion: string;
  namespace: string;
  ownerReferences: OwnerReference[];
}
export interface Time {}
export interface ResourceClusterVulnerabilityReportReport {
  severity: string;
  score: number;
  title: string;
  primaryLink: string;
  firstObserved: Time;
  lastObserved: Time;
  owners: ResourceClusterVulnerabilityReportReportOwner[];
  status: ResourceClusterVulnerabilityReportReportStatus;
}
export interface ResourceClusterVulnerabilityReportSummary {
  critical: number;
  high: number;
  medium: number;
  low: number;
  unknown: number;
}
export interface ResourceClusterVulnerabilityReport {
  summary: ResourceClusterVulnerabilityReportSummary;
  report: { [key: string]: ResourceClusterVulnerabilityReportReport };
}
export interface ResourceClusterComplianceReport {}
export interface ResourceConfigurationSpec {
  type: string;
  b64enc: boolean;
  data: string;
}
export interface ResourceConfiguration {
  spec: ResourceConfigurationSpec;
}
export interface ResourceProjectSpecRole {
  upn: string;
  name: string;
  role: string;
  email: string;
  phone: string;
}
export interface ResourceProjectSpec {
  projectName: string;
  description: string;
  active: boolean;
  createdTime: string;
  updatedTime: string;
  roles: ResourceProjectSpecRole[];
  workorder: string;
  serviceTag: string;
  tags: string[];
}
export interface ResourceProject {
  spec: ResourceProjectSpec;
}
export interface ResourceClusterOrderStatus {
  status: string;
  phase: string;
  conditions: ResourceKubernetesClusterStatusCondition[];
  createdTime: string;
  updatedTime: string;
  lastObservedTime: string;
}
export interface ResourceClusterOrderSpecNodePool {
  name: string;
  machineClass: string;
  count: number;
}
export interface ResourceClusterOrderSpec {
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
}
export interface ResourceClusterOrder {
  spec: ResourceClusterOrderSpec;
  status: ResourceClusterOrderStatus;
}
export interface ResourceKubernetesClusterStatusCondition {
  type: string;
  status: string;
  lastTransitionTime: string;
  reason: string;
  message: string;
}
export interface ResourceKubernetesClusterStatus {
  status: string;
  phase: string;
  conditions: ResourceKubernetesClusterStatusCondition[];
  kubernetesVersion: string;
  providerStatus: { [key: string]: any };
  createdTime: string;
  updatedTime: string;
  lastObservedTime: string;
}
export interface ResourceKubernetesClusterSpecEndpoint {
  type: string;
  address: string;
}
export interface ResourceKubernetesClusterSpecTopologyWorkers {
  name: string;
  replicas: number;
  version: string;
  machineClass: string;
}
export interface ResourceKubernetesClusterSpecTopologyControlPlane {
  replicas: number;
  version: string;
  machineClass: string;
}
export interface ResourceKubernetesClusterSpecTopology {
  controlPlane: ResourceKubernetesClusterSpecTopologyControlPlane;
  workers: ResourceKubernetesClusterSpecTopologyWorkers[];
}
export interface ResourceKubernetesClusterSpecProviderSpecAzureSpec {
  subscriptionId: string;
  resourceGroup: string;
}
export interface ResourceKubernetesClusterSpecProviderSpecTanzuSpec {
  supervisorClusterName: string;
  namespace: string;
}
export interface ResourceKubernetesClusterSpecProviderSpec {
  tanzuSpec: ResourceKubernetesClusterSpecProviderSpecTanzuSpec;
  azureSpec: ResourceKubernetesClusterSpecProviderSpecAzureSpec;
}
export interface ResourceKubernetesClusterSpecToolingConfig {
  splunkIndex: string;
}
export interface ResourceKubernetesClusterSpec {
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
}
export interface ResourceKubernetesCluster {
  spec: ResourceKubernetesClusterSpec;
  status: ResourceKubernetesClusterStatus;
}
export interface ResourceVirtualMachineClassBindingClassRef {
  apiVersion: string;
  kind: string;
  name: string;
}
export interface ResourceVirtualMachineClassBinding {
  classRef: ResourceVirtualMachineClassBindingClassRef;
}
export interface ResourceVirtualMachineClassSpecHardwareInstanceStorage {
  storageClass: string;
}
export interface ResourceVirtualMachineClassSpecHardware {
  cpus: number;
  instanceStorage: ResourceVirtualMachineClassSpecHardwareInstanceStorage;
}
export interface ResourceVirtualMachineClassSpec {
  description: string;
  hardware: ResourceVirtualMachineClassSpecHardware;
}
export interface ResourceVirtualMachineClass {
  spec: ResourceVirtualMachineClassSpec;
  status: { [key: string]: string };
}
export interface ResourceTanzuKubernetesReleaseStatusCondition {
  lastTransitionTime: string;
  message: string;
  reason: string;
  severity: string;
  status: string;
  type: string;
}
export interface ResourceTanzuKubernetesReleaseStatus {
  conditions: ResourceTanzuKubernetesReleaseStatusCondition[];
}
export interface ResourceTanzuKubernetesReleaseSpecNodeImageRef {
  apiVersion: string;
  fieldPath: string;
  kind: string;
  namespace: string;
  uid: string;
}
export interface ResourceTanzuKubernetesReleaseSpecImage {
  name: string;
  repository: string;
  tag: string;
}
export interface ResourceTanzuKubernetesReleaseSpec {
  images: ResourceTanzuKubernetesReleaseSpecImage[];
  kubernetesVersion: string;
  nodeImageRef: ResourceTanzuKubernetesReleaseSpecNodeImageRef;
  repository: string;
  version: string;
}
export interface ResourceTanzuKubernetesRelease {
  spec: ResourceTanzuKubernetesReleaseSpec;
  status: ResourceTanzuKubernetesReleaseStatus;
}
export interface ResourceTanzuKubernetesClusterStatusConditions {
  lastTransitionTime: string;
  message: string;
  reason: string;
  severity: string;
  status: string;
  type: string;
}
export interface ResourceTanzuKubernetesClusterStatusAPIEndpoints {
  host: string;
  port: number;
}
export interface ResourceTanzuKubernetesClusterStatus {
  apiEndpoints: ResourceTanzuKubernetesClusterStatusAPIEndpoints[];
  conditions: ResourceTanzuKubernetesClusterStatusConditions[];
  phase: string;
  totalWorkerReplicas: number;
  version: string;
}
export interface ResourceTanzuKubernetesClusterSpecTopologyNodePoolsVolumes {
  capasity: { [key: string]: string };
  mountPath: string;
  name: string;
  storageClass: string;
}
export interface ResourceTanzuKubernetesClusterSpecTopologyNodePoolsTkrReference {
  fieldPath: string;
  kind: string;
  name: string;
  namespace: string;
  uid: string;
}
export interface ResourceTanzuKubernetesClusterSpecTopologyNodePoolsTkr {
  reference: ResourceTanzuKubernetesClusterSpecTopologyNodePoolsTkrReference;
}
export interface ResourceTanzuKubernetesClusterSpecTopologyNodePoolsTaints {
  effect: string;
  key: string;
  timeAdded: string;
  value: string;
}
export interface ResourceTanzuKubernetesClusterSpecTopologyNodePools {
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
}
export interface ResourceTanzuKubernetesClusterSpecTopologyControlPlaneTkrReference {
  name: string;
  kind: string;
  namespace: string;
  uid: string;
}
export interface ResourceTanzuKubernetesClusterSpecTopologyControlPlaneTkr {
  reference: ResourceTanzuKubernetesClusterSpecTopologyControlPlaneTkrReference;
}
export interface ResourceTanzuKubernetesClusterSpecTopologyControlPlane {
  nodeDrainTimeout: string;
  replicas: number;
  storageClass: string;
  tkr: ResourceTanzuKubernetesClusterSpecTopologyControlPlaneTkr;
  vmClass: string;
}
export interface ResourceTanzuKubernetesClusterSpecTopology {
  controlPlane: ResourceTanzuKubernetesClusterSpecTopologyControlPlane;
  nodePools: ResourceTanzuKubernetesClusterSpecTopologyNodePools[];
}
export interface ResourceTanzuKubernetesClusterSpecSettingsStorage {
  classes: string[];
  defaultClass: string;
}
export interface ResourceTanzuKubernetesClusterSpecSettingsNetworkTrustAdditionalTrustedCA {
  data: string;
  name: string;
}
export interface ResourceTanzuKubernetesClusterSpecSettingsNetworkTrust {
  additionalTrustedCAs: ResourceTanzuKubernetesClusterSpecSettingsNetworkTrustAdditionalTrustedCA[];
}
export interface ResourceTanzuKubernetesClusterSpecSettingsNetworkServices {
  cidrBlocks: string[];
}
export interface ResourceTanzuKubernetesClusterSpecSettingsNetworkProxy {
  httpProxy: string;
  httpsProxy: string;
  noProxy: string;
}
export interface ResourceTanzuKubernetesClusterSpecSettingsNetworkPods {
  cidrBlocks: string[];
}
export interface ResourceTanzuKubernetesClusterSpecSettingsNetworkCni {
  name: string;
}
export interface ResourceTanzuKubernetesClusterSpecSettingsNetwork {
  cni: ResourceTanzuKubernetesClusterSpecSettingsNetworkCni;
  pods: ResourceTanzuKubernetesClusterSpecSettingsNetworkPods;
  proxy: ResourceTanzuKubernetesClusterSpecSettingsNetworkProxy;
  serviceDomain: string;
  services: ResourceTanzuKubernetesClusterSpecSettingsNetworkServices;
  trust: ResourceTanzuKubernetesClusterSpecSettingsNetworkTrust;
}
export interface ResourceTanzuKubernetesClusterSpecSettings {
  network: ResourceTanzuKubernetesClusterSpecSettingsNetwork;
  storage: ResourceTanzuKubernetesClusterSpecSettingsStorage;
}
export interface ResourceTanzuKubernetesClusterSpecDistribution {
  fullVersion: string;
  version: string;
}
export interface ResourceTanuzKuberntesClusterSpec {
  distribution: ResourceTanzuKubernetesClusterSpecDistribution;
  settings: ResourceTanzuKubernetesClusterSpecSettings;
  topology: ResourceTanzuKubernetesClusterSpecTopology;
}
export interface ResourceTanzuKubernetesCluster {
  spec: ResourceTanuzKuberntesClusterSpec;
  status?: ResourceTanzuKubernetesClusterStatus;
}
export interface ResourceRbacAssessmentReport {
  report: ResourceVulnerabilityReportReport;
}
export interface ResourceConfigAuditReport {
  report: ResourceVulnerabilityReportReport;
}
export interface ResourceExposedSecretReport {
  report: ResourceVulnerabilityReportReport;
}
export interface ResourceVulnerabilityReportReportVulnerability {
  vulnerabilityID: string;
  severity: string;
  score: number;
  title: string;
  resource: string;
  primaryLink: string;
  installedVersion: string;
  fixedVersion: string;
}
export interface ResourceVulnerabilityReportReportArtifact {
  digest: string;
  repository: string;
  tag: string;
}
export interface AquaReportScanner {
  name: string;
  vendor: string;
  version: string;
}
export interface AquaReportSummary {
  criticalCount: number;
  highCount: number;
  lowCount: number;
  mediumCount: number;
  total?: number;
}
export interface ResourceVulnerabilityReportReport {
  summary: AquaReportSummary;
  scanner: AquaReportScanner;
  artifact: ResourceVulnerabilityReportReportArtifact;
  updateTimestamp: string;
  vulnerabilities: ResourceVulnerabilityReportReportVulnerability[];
}
export interface ResourceVulnerabilityReport {
  report: ResourceVulnerabilityReportReport;
}
export interface ResourceIngressClassSpecParameters {
  apiGroup: string;
  kind: string;
  name: string;
  namespace: string;
  scope: string;
}
export interface ResourceIngressClassSpec {
  controller: string;
  parameters: ResourceIngressClassSpecParameters;
}
export interface ResourceIngressClass {
  spec: ResourceIngressClassSpec;
}
export interface ResourceIngressStatusLoadBalancerIngress {
  hostname: string;
  ip: string;
}
export interface ResourceIngressStatusLoadBalancer {
  ingress: ResourceIngressStatusLoadBalancerIngress[];
}
export interface ResourceIngressStatus {
  loadBalancer: ResourceIngressStatusLoadBalancer;
}
export interface ResourceIngressSpecTls {
  hosts: string[];
  secretName: string;
}
export interface ResourceIngressSpecRulesHttpPaths {
  backend: ResourceIngressSpecRulesHttpPathsBackend;
  path: string;
  pathType: string;
}
export interface ResourceIngressSpecRulesHttp {
  paths: ResourceIngressSpecRulesHttpPaths[];
}
export interface ResourceIngressSpecRules {
  apiGroup: string;
  http: ResourceIngressSpecRulesHttp;
}
export interface ResourceIngressSpecBackendServicePort {
  name?: string;
  number?: number;
}
export interface ResourceIngressSpecBackendService {
  name?: string;
  port?: ResourceIngressSpecBackendServicePort;
}
export interface ResourceIngressSpecBackendResource {
  apiGroup?: string;
  kind?: string;
  name?: string;
}
export interface ResourceIngressSpecRulesHttpPathsBackend {
  resource?: ResourceIngressSpecBackendResource;
  service?: ResourceIngressSpecBackendService;
}
export interface ResourceIngressSpec {
  defaultBackend?: ResourceIngressSpecRulesHttpPathsBackend;
  ingressClassName: string;
  rules: ResourceIngressSpecRules[];
  tls: ResourceIngressSpecTls[];
}
export interface ResourceIngress {
  spec: ResourceIngressSpec;
  status: ResourceIngressStatus;
}
export interface ResourceDaemonSetStatus {
  numberReady: number;
  numberUnavailable: number;
  currentReplicas: number;
  numberAvailable: number;
  updatedNumberScheduled: number;
  desiredNumberScheduled: number;
  currentNumberScheduled: number;
}
export interface ResourceDaemonSet {
  status: ResourceDaemonSetStatus;
}
export interface ResourceStatefulSetStatus {
  replicas: number;
  availableReplicas: number;
  currentReplicas: number;
  readyReplicas: number;
  updatedReplicas: number;
}
export interface ResourceStatefulSet {
  status: ResourceStatefulSetStatus;
}
export interface ResourceReplicaSetStatus {
  availableReplicas: number;
  readyReplicas: number;
  replicas: number;
}
export interface ResourceReplicaSetSpecSelectorMatchExpressions {
  key: string;
  operator: string;
  values: string[];
}
export interface ResourceReplicaSetSpecSelector {
  matchExpressions: ResourceReplicaSetSpecSelectorMatchExpressions[];
  matchLabels: { [key: string]: string };
}
export interface ResourceReplicaSetSpec {
  replicas: number;
  selector: ResourceReplicaSetSpecSelector;
}
export interface ResourceReplicaSet {
  spec: ResourceReplicaSetSpec;
  status: ResourceReplicaSetStatus;
}
export interface ResourcePodStatus {
  message?: string;
  phase?: string;
  reason?: string;
  startTime?: string;
}
export interface ResourcePodSpecContainersPorts {
  name?: string;
  containerPort?: number;
  protocol?: string;
}
export interface ResourcePodSpecContainers {
  name?: string;
  image?: string;
  ports?: ResourcePodSpecContainersPorts[];
}
export interface ResourcePodSpec {
  containers?: ResourcePodSpecContainers[];
  serviceAccountName?: string;
  nodeName?: string;
}
export interface ResourcePod {
  spec?: ResourcePodSpec;
  status?: ResourcePodStatus;
}
export interface IntOrString {
  Type: number;
  IntVal: number;
  StrVal: string;
}
export interface ResourceServicePorts {
  appProtocol: string;
  name: string;
  port: number;
  protocol: string;
  targetPort: IntOrString;
}
export interface ResourceServiceSpec {
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
}
export interface ResourceService {
  spec: ResourceServiceSpec;
}
export interface ResourceCertificateStatusCondition {
  lastTransitionTime: string;
  observedGeneration: number;
  message: string;
  reason: string;
  status: string;
  type: string;
}
export interface ResourceCertificateStatus {
  notBefore: string;
  notAfter: string;
  renewalTime: string;
  conditions: ResourceCertificateStatusCondition[];
}
export interface ResourceCertificateSpecIssuerref {
  group: string;
  kind: string;
  name: string;
}
export interface ResourceCertificateSpec {
  dnsNames: string[];
  secretName: string;
  issuerRef: ResourceCertificateSpecIssuerref;
  usages?: string[];
}
export interface ResourceCertificate {
  spec: ResourceCertificateSpec;
  status: ResourceCertificateStatus;
}
export interface ResourceAppProjectSpec {
  description: string;
  sourceRepos: string[];
  destinations: ResourceApplicationSpecDestination[];
}
export interface ResourceAppProject {
  spec: ResourceAppProjectSpec;
}
export interface ResourceApplicationStatusSync {
  revision: string;
  status: string;
}
export interface ResourceApplicationStatusHealth {
  message: string;
  status: string;
}
export interface ResourceApplicationStatusOperationstate {
  startedAt: string;
  finishedAt: string;
  phase: string;
  status: string;
  message: string;
}
export interface ResourceApplicationStatus {
  sourceType: string;
  reconciledAt: string;
  operationState: ResourceApplicationStatusOperationstate;
  health: ResourceApplicationStatusHealth;
  sync: ResourceApplicationStatusSync;
}
export interface ResourceApplicationSpecSyncpolicyRetryBackoff {
  duration: string;
  factor: number;
  maxDuration: string;
}
export interface ResourceApplicationSpecSyncpolicyRetry {
  backoff?: ResourceApplicationSpecSyncpolicyRetryBackoff;
  limit: number;
}
export interface ResourceApplicationSpecSyncpolicyAutomated {
  allowEmpty: boolean;
  prune: boolean;
  selfHeal: boolean;
}
export interface ResourceApplicationSpecSyncpolicy {
  automated?: ResourceApplicationSpecSyncpolicyAutomated;
  retry?: ResourceApplicationSpecSyncpolicyRetry;
  syncOptions: string[];
}
export interface ResourceApplicationSpecSource {
  chart?: string;
  path?: string;
  repoURL: string;
  targetRevision: string;
}
export interface ResourceApplicationSpecDestination {
  name: string;
  namespace: string;
  server: string;
}
export interface ResourceApplicationSpec {
  destination: ResourceApplicationSpecDestination;
  project: string;
  source: ResourceApplicationSpecSource;
  syncPolicy: ResourceApplicationSpecSyncpolicy;
}
export interface ResourceApplication {
  spec: ResourceApplicationSpec;
  status: ResourceApplicationStatus;
}
export interface ResourcePolicyReportSummary {
  error: number;
  fail: number;
  pass: number;
  skip: number;
  warn: number;
}
export interface ResourcePolicyReportResultsResources {
  uid: string;
  apiVersion: string;
  kind: string;
  name: string;
}
export interface ResourcePolicyReportResults {
  policy: string;
  message: string;
  category: string;
  properties: { [key: string]: string };
  severity: string;
  result: string;
  resources: ResourcePolicyReportResultsResources[];
}
export interface ResourcePolicyReport {
  results: ResourcePolicyReportResults[];
  summary: ResourcePolicyReportSummary;
  lastReported?: string;
}
export interface ResourceStorageClass {
  allowVolumeExpansion: boolean;
  provisioner: string;
  reclaimPolicy: string;
  volumeBindingMode: string;
}
export interface ResourceDeploymentStatus {
  replicas: number;
  availableReplicas: number;
  readyReplicas: number;
  updatedReplicas: number;
}
export interface ResourceDeployment {
  status: ResourceDeploymentStatus;
}
export interface ResourcePersistentVolumeClaimStatus {
  accessModes: string[];
  capacity: { [key: string]: string };
  phase: string;
}
export interface ResourcePersistentVolumeClaimSpecResources {
  limits?: { [key: string]: string };
  requests: { [key: string]: string };
}
export interface ResourcePersistentVolumeClaimSpec {
  accessModes: string[];
  resources: ResourcePersistentVolumeClaimSpecResources;
  storageClassName: string;
  volumeMode: string;
  volumeName: string;
}
export interface ResourcePersistentVolumeClaim {
  spec: ResourcePersistentVolumeClaimSpec;
  status: ResourcePersistentVolumeClaimStatus;
}
export interface ResourceNodeStatusNodeinfo {
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
}
export interface ResourceNodeStatusConditions {
  lastHeartbeatTime: string;
  lastTransitionTime: string;
  message: string;
  reason: string;
  status: string;
  type: string;
}
export interface ResourceNodeStatusCapacity {
  cpu: string;
  ephemeralStorage: string;
  memory: string;
  pods: string;
}
export interface ResourceNodeStatusAddresses {
  address: string;
  type: string;
}
export interface ResourceNodeStatus {
  addresses: ResourceNodeStatusAddresses[];
  capacity: ResourceNodeStatusCapacity;
  conditions: ResourceNodeStatusConditions[];
  nodeInfo: ResourceNodeStatusNodeinfo;
}
export interface ResourceNodeSpecTaints {
  effect: string;
  key: string;
}
export interface ResourceNodeSpec {
  podCIDR?: string;
  podCIDRs?: string[];
  providerID?: string;
  taints?: ResourceNodeSpecTaints[];
}
export interface ResourceNode {
  spec: ResourceNodeSpec;
  status: ResourceNodeStatus;
}
export interface ResourceNamespace {}
export interface RorResourceOwnerReference {
  scope: string;
  subject: string;
}
export interface ResourceRorMeta {
  version?: string;
  lastReported?: string;
  internal?: boolean;
  hash?: string;
  ownerref?: RorResourceOwnerReference;
  action?: string;
}
export interface FieldsV1 {}
export interface ManagedFieldsEntry {
  manager?: string;
  operation?: string;
  apiVersion?: string;
  time?: Time;
  fieldsType?: string;
  fieldsV1?: FieldsV1;
  subresource?: string;
}
export interface OwnerReference {
  apiVersion: string;
  kind: string;
  name: string;
  uid: string;
  controller?: boolean;
  blockOwnerDeletion?: boolean;
}
export interface Time {}
export interface ObjectMeta {
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
}
export interface Resource {
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
  vulnerabilityevent?: ResourceVulnerabilityEvent;
  virtualmachine?: ResourceVirtualMachine;
  endpoints?: ResourceEndpoints;
  networkpolicy?: ResourceNetworkPolicy;
}
export interface ResourceSet {
  resources?: Resource[];
}
export interface ResourceQueryFilter {
  field?: string;
  value?: string;
  type?: string;
  operator?: string;
}
export interface ResourceQueryOrder {
  field?: string;
  descending?: boolean;
  index?: number;
}
export interface GroupVersionKind {
  Group: string;
  Version: string;
  Kind: string;
}
export interface ResourceQuery {
  versionkind?: GroupVersionKind;
  uids?: string[];
  ownerrefs?: RorResourceOwnerReference[];
  fields?: string[];
  order?: ResourceQueryOrder[];
  filters?: ResourceQueryFilter[];
  offset?: number;
  limit?: number;
  relatedresources?: ResourceQuery[];
}
