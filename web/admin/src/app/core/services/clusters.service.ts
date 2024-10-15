import { Filter } from '../models/apiFilter';
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { map, Observable } from 'rxjs';
import { PaginationResult } from '../models/paginatedResult';
import { Cluster } from '../models/cluster';
import { ClusterMetadata } from '../models/clusterMetadata';
import { VulnerabilityReportsView } from '../models/vulnerabilityReport';
import { PolicyReportView } from '../models/policyReport';
import { ConfigService } from './config.service';

@Injectable({
  providedIn: 'root',
})
export class ClustersService {
  constructor(
    private httpClient: HttpClient,
    private configService: ConfigService,
  ) {}

  getByFilter(filter: Filter): Observable<PaginationResult<Cluster>> {
    const url = `${this.configService.config.rorApi}/v1/clusters/filter`;
    return this.httpClient.post<PaginationResult<Cluster>>(url, filter);
  }

  getByWorkspace(workspaceName: string, filter: Filter): Observable<PaginationResult<any>> {
    const url = `${this.configService.config.rorApi}/v1/clusters/workspace/${workspaceName}/filter`;
    return this.httpClient.post<PaginationResult<any>>(url, filter);
  }

  getById(clusterId: string): Observable<any> {
    const url = `${this.configService.config.rorApi}/v1/cluster/${clusterId}`;
    return this.httpClient.get<any[]>(url);
  }

  getMetadata(): Observable<Map<string, string[]>> {
    let url = `${this.configService.config.rorApi}/v1/clusters/metadata`;
    return this.httpClient.get<Map<string, string[]>>(url);
  }

  getPolicyreports(clusterId: string): Observable<PolicyReportView> {
    let url = `${this.configService.config.rorApi}/v1/clusters/${clusterId}/views/policyreports`;
    return this.httpClient.get<PolicyReportView>(url);
  }

  getVulnerabilityReports(clusterId: string): Observable<VulnerabilityReportsView> {
    let url = `${this.configService.config.rorApi}/v1/clusters/${clusterId}/views/vulnerabilityreports`;
    return this.httpClient.get<VulnerabilityReportsView>(url);
  }

  exists(clusterId: string): Observable<boolean> {
    const url = `${this.configService.config.rorApi}/v1/cluster/${clusterId}/exists`;
    return this.httpClient.get<boolean[]>(url).pipe(
      map((data: any) => {
        if (!data) {
          return false;
        }
        return data.exists === true;
      }),
    );
  }

  updateMetadata(clusterId: string, metadata: ClusterMetadata): Observable<void> {
    const url = `${this.configService.config.rorApi}/v1/cluster/${clusterId}/metadata`;
    return this.httpClient.patch<void>(url, metadata);
  }
}
