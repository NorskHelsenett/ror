import { ChangeDetectionStrategy, ChangeDetectorRef, Component, Input, OnInit } from '@angular/core';
import { LangChangeEvent, TranslateService } from '@ngx-translate/core';
import { TreeNode } from 'primeng/api';
import { Subscription } from 'rxjs';
import { ThemeService } from '../../../core/services/theme.service';

@Component({
  selector: 'app-cluster-ingress-rules',
  templateUrl: './cluster-ingress-rules.component.html',
  styleUrls: ['./cluster-ingress-rules.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ClusterIngressRulesComponent implements OnInit {
  @Input() cluster: any = undefined;
  @Input() ingressName: any = undefined;

  ingresses: any[] = [];
  ingress: any;
  rulenodes: any[];
  links: any[];
  nodes: any;
  thisColor: any;
  thisComment: any;
  subscriptions: Subscription = new Subscription();
  newText: any;
  isDark: boolean;

  data: TreeNode[] = [];

  constructor(
    private changeDetector: ChangeDetectorRef,
    private translateService: TranslateService,
    private themeService: ThemeService,
  ) {}

  ngOnInit(): void {
    if (!this.cluster) {
      return;
    }

    this.subscriptions.add(
      this.translateService.onLangChange.pipe().subscribe((languageChangeEvent: LangChangeEvent) => {
        this.renderGraph();
      }),
    );
    this.subscriptions.add(
      this.themeService.isDark.subscribe((value) => {
        this.isDark = value;
        this.renderGraph();
      }),
    );

    this.ingresses = this.cluster?.ingresses;
    this.ingress = this.ingresses.filter((object) => {
      return object['name'] == this.ingressName;
    })[0];

    this.renderGraph();
  }

  renderGraph(): void {
    if (!this.ingress) {
      return;
    }

    if (
      !this.ingress?.class ||
      !this.ingress?.class?.startsWith('avi') ||
      !['helsenett', 'internett', 'datacenter'].includes(this.ingress?.class.split('-')[this.ingress?.class.split('-').length - 1])
    ) {
      this.ingress.classError = true;
    }

    let root: TreeNode = {
      label: this.ingress?.name,
      type: 'ingress',
      data: this.ingress,
      expanded: true,
    };
    let graph: TreeNode[] = [root];

    root.children = this.createRootNodeChildren();

    this.data = graph;
    this.changeDetector?.detectChanges();
  }

  private createRootNodeChildren(): TreeNode[] {
    let children: TreeNode[] = [];
    this.ingress?.ingressrules?.forEach((ingressRule: any) => {
      if (!ingressRule) {
        return null;
      }

      ingressRule?.rules?.forEach((rule: any) => {
        let ingressRuleCopy: any = { ...ingressRule, multiple: false };
        if (this.ingress?.ingressrules?.length > 1) {
          ingressRuleCopy.multiple = true;
        }
        const ruleNode: TreeNode = {
          label: rule?.path,
          type: 'rule',
          data: ingressRuleCopy,
          expanded: true,

          children: [],
        };

        const serviceNode: TreeNode = this.createServiceNode(rule);
        ruleNode?.children.push(serviceNode);
        children.push(ruleNode);
      });
    });
    return children;
  }

  private createServiceNode(rule: any): TreeNode {
    const serviceNode: TreeNode = {
      label: rule?.service?.name,
      type: 'service',
      data: rule?.service,
      expanded: true,
      children: [],
    };

    rule?.service.endpoints?.forEach((endpoint: any) => {
      serviceNode?.children.push({
        label: endpoint?.podnamespace,
        data: endpoint,
        type: 'pod',
        style: 'bg-red-500',
      });
    });
    return serviceNode;
  }
}
