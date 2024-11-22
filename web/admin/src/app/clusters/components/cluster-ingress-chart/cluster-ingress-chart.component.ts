import { TooltipModule } from 'primeng/tooltip';
import { ChangeDetectionStrategy, ChangeDetectorRef, Component, effect, inject, ViewEncapsulation } from '@angular/core';
import { ClusterIngressService } from '../../services/cluster-ingress.service';
import { OrganizationChartModule } from 'primeng/organizationchart';
import {
  Resource,
  ResourceIngressSpecRules,
  ResourceIngressSpecRulesHttpPaths,
  ResourceIngressStatusLoadBalancerIngress,
} from '@rork8s/ror-resources/models';
import { TreeNode } from 'primeng/api';
import { TranslateModule } from '@ngx-translate/core';
import { CommonModule, NgOptimizedImage } from '@angular/common';

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
  certificates: Resource[] | undefined;
  pods: Resource[] | undefined;

  private changeDetector = inject(ChangeDetectorRef);
  private clusterIngressService = inject(ClusterIngressService);

  constructor() {
    effect(() => {
      this.cluster = this.clusterIngressService.getCluster();
      this.ingress = this.clusterIngressService.getIngress();
      this.services = this.clusterIngressService.getServices();
      this.pods = this.clusterIngressService.getPods();
      this.certificates = this.clusterIngressService.getCertificates();
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
    loadbalancer.ingress.forEach((lbIngressInfo: ResourceIngressStatusLoadBalancerIngress) => {
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
    this.ingress?.ingress?.spec?.rules?.forEach((ingressRule: ResourceIngressSpecRules) => {
      let pathNodes: TreeNode[] = [];
      ingressRule?.http?.paths?.forEach((path: ResourceIngressSpecRulesHttpPaths) => {
        let pathNode: TreeNode = {
          label: path.path,
          type: 'path',
          data: path,
          expanded: true,
          children: [],
        };

        let podNodes: TreeNode[] = [];
        this.pods?.forEach((pod: Resource) => {
          if (path.backend.service?.name !== pod.metadata.labels['app.kubernetes.io/instance']) {
            return null;
          }
          podNodes.push({
            label: pod.metadata.name,
            type: 'pod',
            data: pod,
            expanded: true,
          });
        });

        let serviceNode: TreeNode = {
          label: path.backend.service?.name,
          type: 'service',
          data: this.services?.find((service: Resource) => service.metadata.name === path.backend.service?.name),
          expanded: true,
          children: podNodes,
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

    ingressNode.children = ruleNodes;
    graph.push(ingressNode);

    this.data = graph;
  }

  // private createRootNodeChildren(): TreeNode[] {
  //   let children: TreeNode[] = [];
  //   this.ingress?.ingress?.spec?.rules?.forEach((ingressRule: any) => {
  //     if (!ingressRule) {
  //       return null;
  //     }

  //     ingressRule?.rules?.forEach((rule: any) => {
  //       let ingressRuleCopy: any = { ...ingressRule, multiple: false };
  //       if (this.ingress?.ingress?.spec?.rules?.length > 1) {
  //         ingressRuleCopy.multiple = true;
  //       }
  //       const ruleNode: TreeNode = {
  //         label: rule?.path,
  //         type: 'rule',
  //         data: ingressRuleCopy,
  //         expanded: true,

  //         children: [],
  //       };

  //       const serviceNode: TreeNode = this.createServiceNode(rule);
  //       ruleNode?.children.push(serviceNode);
  //       children.push(ruleNode);
  //     });
  //   });
  //   return children;
  // }

  // private createServiceNode(rule: any): TreeNode {
  //   const serviceNode: TreeNode = {
  //     label: rule?.service?.name,
  //     type: 'service',
  //     data: rule?.service,
  //     expanded: true,
  //     children: [],
  //   };

  //   rule?.service.endpoints?.forEach((endpoint: any) => {
  //     serviceNode?.children.push({
  //       label: endpoint?.podnamespace,
  //       data: endpoint,
  //       type: 'pod',
  //       style: 'bg-red-500',
  //     });
  //   });
  //   return serviceNode;
  // }
}
