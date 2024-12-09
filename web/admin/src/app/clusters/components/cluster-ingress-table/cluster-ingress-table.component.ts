import { Component, EventEmitter, inject, Input, OnInit, Output } from '@angular/core';
import { ResourcesV2ListComponent } from '../../../resourcesv2/components/resources-v2-list/resources-v2-list.component';
import { TranslateModule } from '@ngx-translate/core';
import { ResourceQuery } from '@rork8s/ror-resources/models';
import { ResourcesV2QueryService } from '../../../resourcesv2/services/resources-v2-query.service';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-cluster-ingress-table',
  standalone: true,
  imports: [TranslateModule, ResourcesV2ListComponent],
  templateUrl: './cluster-ingress-table.component.html',
  styleUrl: './cluster-ingress-table.component.scss',
})
export class ClusterIngressTableComponent implements OnInit {
  @Input() cluster: any | undefined;
  @Input() resourceTypes: any[] | undefined;
  @Output() resourceSelected = new EventEmitter<any>();

  private resourcesV2QueryService = inject(ResourcesV2QueryService);
  private router = inject(Router);
  private route = inject(ActivatedRoute);

  private ingressResourceQuery: ResourceQuery = {
    versionkind: {
      Version: 'networking.k8s.io/v1',
      Kind: 'Ingress',
      Group: '',
    },
  };

  ngOnInit() {
    this.resourcesV2QueryService.setQuery(this.ingressResourceQuery);
  }

  showSelectedResource(resource: any) {
    if (!resource) {
      return;
    }

    this.router.navigate(['ingress', resource?.metadata?.name], { relativeTo: this.route });
  }
}
