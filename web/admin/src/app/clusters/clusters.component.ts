import { ChangeDetectionStrategy, ChangeDetectorRef, Component, OnDestroy, OnInit } from '@angular/core';
import { FilterMatchMode, PrimeNGConfig, SelectItem } from 'primeng/api';
import { catchError, finalize, map, Observable, share, Subscription, tap } from 'rxjs';

import { Cluster } from '../core/models/cluster';
import { ClustersService } from '../core/services/clusters.service';
import { ExportService } from '../core/services/export.service';
import { FilterService } from '../core/services/filter.service';
import { Filter } from '../core/models/apiFilter';
import dayjs from 'dayjs';
import dayjsDuration from 'dayjs/plugin/duration';
import { PaginationResult } from '../core/models/paginatedResult';
import { ProjectRole, RoleDefinition } from '../core/models/project';
import { ConfigService } from '../core/services/config.service';
import { AclAccess, AclScopes } from '../core/models/acl-scopes';
import { AclService } from '../core/services/acl.service';
import { SignalService } from '../create/create-cluster/services/signal.service';
import { ClusterEnvironment } from '../core/models/clusterEnvironment';

@Component({
  selector: 'app-clusters',
  templateUrl: './clusters.component.html',
  styleUrls: ['./clusters.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ClustersComponent implements OnInit, OnDestroy {
  clusters$: Observable<PaginationResult<Cluster>> | undefined;
  clustersError: any;

  rowsPerPage = this.configService.config.rowsPerPage;
  rows = this.configService.config.rows;
  loading: boolean;

  selectedColumns: any[] | undefined;
  cols: any[];
  baseCols: any[];

  metadata$: Observable<Map<string, string[]>>;

  filter: Filter;
  lastFilter: Filter;

  datacenters: string[];

  lastLazyLoad: any;
  showExportChoises: boolean;

  matchModeOptions: SelectItem[];
  adminCreate$: Observable<boolean> | undefined;
  aclFetchError: any;

  clusterCreated$: Observable<boolean> | undefined;

  environments: any[] = [
    {
      name: ClusterEnvironment[ClusterEnvironment.Development],
      value: 'dev',
    },
    {
      name: ClusterEnvironment[ClusterEnvironment.Testing],
      value: 'test',
    },
    {
      name: ClusterEnvironment[ClusterEnvironment.QA],
      value: 'qa',
    },
    {
      name: ClusterEnvironment[ClusterEnvironment.Production],
      value: 'prod',
    },
  ];

  private subscriptions = new Subscription();

  constructor(
    private changeDetector: ChangeDetectorRef,
    private clusterService: ClustersService,
    private filterService: FilterService,
    private primengConfig: PrimeNGConfig,
    private exportService: ExportService,
    private configService: ConfigService,
    private aclService: AclService,
    private signalService: SignalService,
  ) {
    dayjs.extend(dayjsDuration);
  }

  ngOnInit(): void {
    this.fetchAcl();
    this.setupColumns();
    this.fetchMetadata();
    this.setupEvents();

    this.primengConfig.ripple = true;
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }

  fetchAcl(): void {
    this.adminCreate$ = this.aclService.check(AclScopes.ROR, AclScopes.Global, AclAccess.Create).pipe(
      share(),
      catchError((error: any) => {
        this.aclFetchError = error;
        this.changeDetector.detectChanges();
        throw error;
      }),
    );
  }

  setupColumns(): void {
    this.selectedColumns = JSON.parse(localStorage.getItem('cluster-table-columns'));
    this.baseCols = [
      {
        field: 'workspace.datacenter.provider',
        header: 'provider',
        searchEN: 'provider',
        searchNO: 'leverandør',
      },
      {
        field: 'clusterName',
        header: 'clusterName',
        searchEN: 'cluster name',
        searchNO: 'clusternavn',
      },
      {
        field: 'healthStatus.health',
        header: 'status',
        searchEN: 'status',
        searchNO: 'status',
      },

      {
        field: 'metrics.cpuPercentage',
        header: 'cpu',
        searchEN: 'cpu',
        searchNO: 'cpu',
      },
      {
        field: 'metrics.memoryPercentage',
        header: 'memory',
        searchEN: 'memory',
        searchNO: 'minne',
      },
      {
        field: 'lastObserved',
        header: 'lastObserved',
        searchEN: 'last heartbeat',
        searchNO: 'sist rapportert',
      },
      {
        field: 'created',
        header: 'created',
        searchEN: 'created',
        searchNO: 'opprettet',
      },
      {
        field: 'versions.nhnTooling.version',
        header: 'nhnTooling',
        searchEN: 'tooling',
        searchNO: 'tooling',
      },
      {
        field: 'metadata.project.name',
        header: 'projectName',
        searchEN: 'project',
        searchNO: 'prosjekt',
      },
    ];
    this.cols = this.baseCols.concat([
      {
        field: 'firstObserved',
        header: 'firstObserved',
        searchEN: 'first heartbeat',
        searchNO: 'først rapportert',
      },
      {
        field: 'workspace.name',
        header: 'workspace',
        searchEN: 'workspace',
        searchNO: 'arbeidsområde',
      },
      {
        field: 'workspace.datacenter.name',
        header: 'datacenter',
        searchEN: 'datacenter',
        searchNO: 'datasenter',
      },
      {
        field: 'versions.nhnTooling.branch',
        header: 'nhnToolingBranch',
        searchEN: 'branch',
        searchNO: 'gren',
      },
      {
        field: 'versions.agent.version',
        header: 'agentVersion',
        searchEN: 'agent',
        searchNO: 'agent',
      },
      {
        field: 'versions.kubernetes',
        header: 'k8sVersion',
        searchEN: 'kubernetes',
        searchNO: 'kubernetes',
      },
      {
        field: 'ingresses.datacenter',
        header: 'ingressesDatacenter',
        searchEN: 'datacenter publications',
        searchNO: 'publikasjoner datasenter',
      },
      {
        field: 'ingresses.health',
        header: 'ingressesHealth',
        searchEN: 'helsenett publications',
        searchNO: 'publikasjoner helsenett',
      },
      {
        field: 'ingresses.internet',
        header: 'ingressesInternet',
        searchEN: 'internet publications',
        searchNO: 'publikasjoner internett',
      },
      {
        field: 'topology.egressIp',
        header: 'egressIP',
        searchEN: 'egress ip',
        searchNO: 'egress ip',
      },
      {
        field: 'environment',
        header: 'environment',
        searchEN: 'environment',
        searchNO: 'miljø',
      },
      {
        field: 'metadata.project.projectMetadata.billing.workorder',
        header: 'workorder',
        searchEN: 'workorder',
        searchNO: 'arbeidsordre',
      },
    ]);
    if (!this.selectedColumns) {
      this.selectedColumns = [...this.baseCols];
    }
    this.matchModeOptions = [
      {
        label: 'Contains',
        value: FilterMatchMode.CONTAINS,
      },
      {
        label: 'Equals',
        value: FilterMatchMode.EQUALS,
      },
    ];
  }

  resetColumns(): void {
    this.selectedColumns = this.baseCols;
    localStorage.removeItem('cluster-table-columns');
    localStorage.removeItem('cluster-table');
    this.fetchClusters(undefined);
    this.changeDetector.detectChanges();
  }

  fetchClusters(event: any): void {
    if (event) {
      this.filter = this.filterService.mapFilter(event);
    }
    this.loading = true;
    this.lastFilter = this.filter;
    this.clustersError = undefined;
    this.clusters$ = this.clusterService.getByFilter(this.filter).pipe(
      share(),
      map((clusters: PaginationResult<any>) => {
        if (clusters?.data == null) {
          clusters.data = [];
        }
        this.datacenters = Array.from(new Set(clusters.data.map((cluster) => cluster['workspace']['datacenter']['name']))).sort();
        return clusters;
      }),
      catchError((error: any) => {
        this.clustersError = error;
        throw error;
      }),
      finalize(() => {
        this.loading = false;
        this.changeDetector.detectChanges();
      }),
    );
  }

  fetchMetadata(): void {
    this.metadata$ = this.clusterService.getMetadata().pipe(
      tap(() => {
        this.changeDetector.detectChanges();
      }),
    );
  }

  getValueFromColumn(cluster: Cluster, column: string): any {
    let nestedColumns: string[] = column.split('.');
    let value: any = cluster[nestedColumns.shift()];
    nestedColumns.forEach((col) => {
      value = value[col];
    });
    return value;
  }

  countPublicationsForType(ingresses: any[], type: string): number {
    let count: number = 0;
    ingresses?.forEach((ingress) => {
      if (ingress?.class?.includes(type)) {
        count++;
      }
    });
    return count;
  }

  diffMinutes(date: Date): number {
    let diff = dayjs(dayjs()).diff(dayjs(date));
    let duration = dayjs.duration(diff);

    if (duration.years() > 0 || duration.months() > 0 || duration.days() > 0) {
      return Number.MAX_VALUE;
    }

    return duration.minutes();
  }
  updateColumns(): void {
    let tableState: Object = JSON.parse(localStorage.getItem('cluster-table'));
    if (!tableState) {
      tableState = { first: 0, rows: this.rows };
    }
    if (!this.selectedColumns || this.selectedColumns?.length === 0) {
      this.selectedColumns = [
        {
          field: 'clusterName',
          header: 'clusterName',
          searchEN: 'cluster name',
          searchNO: 'clusternavn',
        },
      ];
    }
    tableState['columnOrder'] = this.selectedColumns.map((col) => col['field']);
    localStorage.setItem('cluster-table-columns', JSON.stringify(this.selectedColumns));
    localStorage.setItem('cluster-table', JSON.stringify(tableState));
  }

  exportToExcel(): void {
    this.exportData('excel');
  }

  exportToCsv(): void {
    this.exportData('csv');
  }

  private exportData(type: string): void {
    this.subscriptions.add(
      this.clusterService.getByFilter(this.filter).subscribe((clustersPaginated: PaginationResult<Cluster>) => {
        const clusters = this.exportClusters(clustersPaginated?.data);
        if (type === 'csv') {
          this.exportService.exportToCsv(clusters, 'ror-clusters.csv');
        }
        if (type === 'excel') {
          this.exportService.exportAsExcelFile(clusters, 'ror-clusters.xlsx');
        }
      }),
    );
  }

  private exportClusters(clusters: any[]): any[] {
    let exportClusters: any[] = [];
    clusters.forEach((cluster) => {
      let tags: string[] = [];
      if (cluster?.metadata?.serviceTags) {
        const keys = Object.keys(cluster?.metadata?.serviceTags);
        keys.forEach((key: string) => {
          tags.push(key);
        });
      }

      let c: any = {
        clusterId: cluster?.clusterId,
        clusterName: cluster?.clusterName,
        workspaceName: cluster?.workspace?.name,
        datacenterName: cluster?.workspace?.datacenter?.name,

        firstObserved: cluster?.firstObserved,
        lastObserved: cluster?.lastObserved,

        environment: cluster?.environment,
        criticality: cluster?.metadata?.criticality,
        sensitivity: cluster?.metadata?.sensitivity,

        egressIp: cluster?.topology?.egressIp,
        controlPlaneNodesCount: cluster?.topology?.controlPlane?.nodes?.length,
        nodePoolCount: cluster?.topology?.nodePools?.length,

        description: cluster?.metadata?.description,
        workorder: cluster?.metadata?.billing?.workorder || cluster?.metadata?.project?.projectMetadata?.billing?.workorder,
        projectName: cluster?.metadata?.project?.name,

        kubernetesVersion: cluster?.versions?.kubernetes,
        nhnToolingVersion: cluster?.versions?.nhnTooling?.version,

        tags: tags.join(' '),
      };

      let owner: ProjectRole = cluster?.metadata?.roles?.find((role: ProjectRole) => role?.roleDefinition === RoleDefinition.Owner);
      if (owner == null) {
        owner = cluster?.metadata?.project?.projectMetadata?.roles?.find((role: ProjectRole) => role?.roleDefinition === RoleDefinition.Owner);
      }
      c['ownerEmail'] = owner?.contactInfo?.email;
      c['ownerPhone'] = owner?.contactInfo?.phone;

      let responsible: ProjectRole = cluster?.metadata?.roles?.find((role: ProjectRole) => role?.roleDefinition === RoleDefinition.Responsible);
      if (responsible == null) {
        responsible = cluster?.metadata?.project?.projectMetadata?.roles?.find(
          (role: ProjectRole) => role?.roleDefinition === RoleDefinition.Responsible,
        );
      }
      c['responsibleEmail'] = responsible?.contactInfo?.email;
      c['responsiblePhone'] = responsible?.contactInfo?.phone;

      exportClusters.push(c);
    });
    return exportClusters;
  }

  private setupEvents(): void {
    this.clusterCreated$ = this.signalService?.clusterCreated$.pipe(
      tap((created: any) => {
        if (created) {
          this.fetchClusters(this.lastLazyLoad);
          this.changeDetector.detectChanges();
        }
      }),
    );
  }
}
