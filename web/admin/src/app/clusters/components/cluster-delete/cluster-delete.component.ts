import { Component, Input, OnInit } from '@angular/core';
import { TranslateService } from '@ngx-translate/core';
import { ClipboardService } from 'ngx-clipboard';
import { MessageService } from 'primeng/api';
import { ResourcesService } from '../../../core/services/resources.service';
import { Observable, catchError, share, tap } from 'rxjs';
import { ResourceNamespace } from '../../../core/models/resources';
import { ResourceQuery } from '../../../core/models/resources-v2';

@Component({
  selector: 'app-cluster-delete',
  templateUrl: './cluster-delete.component.html',
  styleUrl: './cluster-delete.component.scss',
})
export class ClusterDeleteComponent implements OnInit {
  @Input() cluster: any | undefined;

  fetchNamespaces$: Observable<ResourceNamespace[]> | undefined;
  fetchNamespacesError: any;
  canDelete: boolean = false;

  constructor(
    private clipboardService: ClipboardService,
    private messageService: MessageService,
    private translateService: TranslateService,
    private resourceService: ResourcesService,
  ) {}

  ngOnInit(): void {
    this.refresh();
  }

  copy(): void {
    this.clipboardService.copy(this.cluster?.clusterId);
    this.messageService.add({ severity: 'success', summary: this.translateService.instant('common.copied') });
  }

  refresh(): void {
    this.canDelete = false;
    this.fetchNamespacesError = null;
    let query: ResourceQuery = {
      versionkind: {
        Kind: 'Namespace',
        Group: '',
        Version: 'v1',
      },
      ownerrefs: [
        {
          scope: 'cluster',
          subject: this.cluster?.clusterId,
        },
      ],
    };
    this.fetchNamespaces$ = this.resourceService.getResources(query).pipe(
      tap((rs) => {
        rs.resources.forEach((r) => {
          if (r.metadata.name === 'delete-me') {
            this.canDelete = true;
          }
        });
      }),
      catchError((error) => {
        this.fetchNamespacesError = error;
        return [];
      }),
      share(),
    );
  }

  scheduleDeletion(): void {
    // let order: ClusterOrderModel = {
    //   orderType: ClusterOrderType.Delete,
    //   projectId: this.cluster?.
    //   orderBy: '',
    //   provider: this.cluster?.provider,
    //   cluster: this.cluster?.clusterId,
    //   environment: 1,
    // };
    //this.orderService.orderClusterDeletion(order);
  }
}
