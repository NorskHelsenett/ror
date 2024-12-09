import { TooltipModule } from 'primeng/tooltip';
import { ChangeDetectionStrategy, ChangeDetectorRef, Component, effect, inject, ViewEncapsulation } from '@angular/core';
import { ClusterIngressService } from '../../services/cluster-ingress.service';
import { OrganizationChartModule } from 'primeng/organizationchart';
import {
  Resource,
  ResourceEndpointSpecSubsetsAddresses,
  ResourceEndpointSpecSubsetsNotReadyAddresses,
  ResourceIngressSpecRules,
  ResourceIngressSpecRulesHttpPaths,
  ResourceIngressStatusLoadBalancerIngress,
} from '@rork8s/ror-resources/models';
import { TreeNode } from 'primeng/api';
import { TranslateModule, TranslateService } from '@ngx-translate/core';
import { CommonModule, NgOptimizedImage } from '@angular/common';
import { HealthStatus } from '../../../core/models/healthstatus';

@Component({
  selector: 'app-cluster-ingress-chart',
  standalone: true,
  imports: [CommonModule, NgOptimizedImage, TranslateModule, OrganizationChartModule, TooltipModule],
  templateUrl: './cluster-ingress-chart.component.html',
  styleUrl: './cluster-ingress-chart.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
  encapsulation: ViewEncapsulation.Emulated,
})
export class ClusterIngressChartComponent {
  data: TreeNode[] = [];

  cluster: any | undefined;
  ingress: Resource | undefined;
  services: Resource[] | undefined;
  endpoints: Resource[] | undefined;
  certificates: Resource[] | undefined;
  pods: Resource[] | undefined;

  private changeDetector = inject(ChangeDetectorRef);
  private translateService = inject(TranslateService);
  private clusterIngressService = inject(ClusterIngressService);

  constructor() {
    effect(() => {
      this.cluster = this.clusterIngressService.getCluster();
      this.ingress = this.clusterIngressService.getIngress();
      this.services = this.clusterIngressService.getServices();
      this.endpoints = this.clusterIngressService.getEndpoints();
      this.pods = this.clusterIngressService.getPods();
      this.certificates = this.clusterIngressService.getCertificates();
      this.setGraphData();
      this.changeDetector.detectChanges();
    });

    this.translateService.onLangChange.subscribe(() => {
      this.setGraphData();
      this.changeDetector.detectChanges();
    });
  }

  private setGraphData() {
    let graph: TreeNode[] = [];

    let ingressNode: TreeNode = {
      label: this.ingress?.metadata?.name,
      type: 'ingress',
      data: this.ingress,
      expanded: true,
      children: [],
    };

    let loadbalancer = this.ingress?.ingress?.status?.loadBalancer;
    loadbalancer?.ingress?.forEach((lbIngressInfo: ResourceIngressStatusLoadBalancerIngress) => {
      let loadbalancerNode: TreeNode = {
        label: lbIngressInfo?.hostname,
        type: 'loadbalancer',
        data: lbIngressInfo,
        expanded: true,
        children: [ingressNode],
      };
      graph.push(loadbalancerNode);
    });

    let ruleNodes: TreeNode[] = [];
    let pathNodes: TreeNode[] = [];
    this.ingress?.ingress?.spec?.rules?.forEach((ingressRule: ResourceIngressSpecRules) => {
      ingressRule?.http?.paths?.forEach((path: ResourceIngressSpecRulesHttpPaths) => {
        let pathNode: TreeNode = {
          label: path.path,
          type: 'path',
          data: path,
          expanded: true,
          children: [],
        };

        let epNodes: TreeNode[] = [];
        this.endpoints?.forEach((endpoint: any) => {
          if (path?.backend?.service?.name !== endpoint?.metadata?.name) {
            return null;
          }

          endpoint?.endpoints?.subsets?.forEach((subset: any) => {
            subset?.addresses?.forEach((address: ResourceEndpointSpecSubsetsAddresses) => {
              let pods = this.pods?.filter((pod: Resource) => {
                return pod?.metadata?.name === address?.targetRef?.name && address?.targetRef.kind === 'Pod';
              });
              epNodes.push({
                label: `${endpoint?.metadata?.name} (${address?.ip})`,
                type: 'endpoints',
                data: { address, ready: true },
                expanded: true,
                children: pods?.map((pod: Resource) => {
                  return {
                    label: pod?.metadata?.name,
                    type: 'pod',
                    data: pod,
                    expanded: true,
                  };
                }),
              });
            });
            subset?.notReadyAddresses?.forEach((address: ResourceEndpointSpecSubsetsNotReadyAddresses) => {
              let pods = this.pods?.filter((pod: Resource) => {
                return pod.metadata.name === address.targetRef.name && address.targetRef.kind === 'Pod';
              });
              epNodes.push({
                label: `${endpoint?.metadata?.name} (${address?.ip})`,
                type: 'endpoints',
                data: { address, ready: false },
                expanded: true,
                children: pods?.map((pod: Resource) => {
                  return {
                    label: pod.metadata.name,
                    type: 'pod',
                    data: pod,
                    expanded: true,
                  };
                }),
              });
            });
          });
        });

        let serviceNode: TreeNode = {
          label: path.backend.service?.name,
          type: 'service',
          data: this.services?.find((service: Resource) => {
            return service.metadata.name === path.backend.service?.name;
          }),
          expanded: true,
          children: epNodes,
        };

        pathNode.children.push(serviceNode);
        pathNodes.push(pathNode);
      });
      let rule: TreeNode = {
        label: 'rule',
        type: 'rule',
        data: ingressRule,
        expanded: true,
        children: [],
      };
      rule.children = pathNodes;
      ruleNodes.push(rule);
    });

    ingressNode.children = pathNodes;
    graph.push(ingressNode);

    this.data = graph;
  }

  getHealthStatus(): HealthStatus {
    return this.clusterIngressService.getHealthStatus();
  }

  getHealthStatusForIngress(): HealthStatus {
    return this.clusterIngressService.getHealthStatusForIngress();
  }

  getHealthStatusForIngressHttpPath(path: ResourceIngressSpecRulesHttpPaths): HealthStatus {
    return this.clusterIngressService.getHealthStatusForIngressPath(path);
  }

  getHealthStatusForServices(): HealthStatus {
    return this.clusterIngressService.getHealthStatusForServices();
  }

  getHealthStatusForEndpoints(): HealthStatus {
    return this.clusterIngressService.getHealthStatusForEndpoints();
  }

  getHealthStatusForPod(pod: Resource): HealthStatus {
    return this.clusterIngressService.getHealthStatusForPod(pod);
  }
}
