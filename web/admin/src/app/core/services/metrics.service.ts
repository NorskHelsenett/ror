import { Filter } from '../models/apiFilter';
import { MetricsList } from '../models/metricsList';
import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { Observable } from 'rxjs';

import { MetricData } from '../models/overallMetrics';
import { Metrics } from '../models/metrics';
import { MetricsCustom } from '../models/metricsCustom';
import { ConfigService } from './config.service';
import { PaginationResult } from '../models/paginatedResult';

@Injectable({
  providedIn: 'root',
})
export class MetricsService {
  constructor(
    private httpClient: HttpClient,
    private configService: ConfigService,
  ) {}

  getTotalForUser(): Observable<Metrics> {
    const url = `${this.configService.config.rorApi}/v1/metrics`;
    return this.httpClient.get<Metrics>(url);
  }

  getTotal(): Observable<Metrics> {
    const url = `${this.configService.config.rorApi}/v1/metrics/total`;
    return this.httpClient.get<Metrics>(url);
  }

  getForDatacenters(): Observable<MetricsList> {
    const url = `${this.configService.config.rorApi}/v1/metrics/datacenters`;
    return this.httpClient.get<MetricsList>(url);
  }

  getForDatacenter(datacenterName: string): Observable<MetricData> {
    const url = `${this.configService.config.rorApi}/v1/metrics/datacenter/${datacenterName}`;
    return this.httpClient.get<MetricData>(url);
  }

  getForWorkspaces(filter: Filter): Observable<PaginationResult<any>> {
    const url = `${this.configService.config.rorApi}/v1/metrics/workspaces/filter`;
    return this.httpClient.post<PaginationResult<any>>(url, filter);
  }

  getForWorkspacesByDatacenter(datacenter: string, filter: Filter): Observable<PaginationResult<any>> {
    const url = `${this.configService.config.rorApi}/v1/metrics/workspaces/datacenter/${datacenter}/filter`;
    return this.httpClient.post<PaginationResult<any>>(url, filter);
  }

  getForWorkspace(workspaceName: string): Observable<any> {
    const url = `${this.configService.config.rorApi}/v1/metrics/workspace/${workspaceName}`;
    return this.httpClient.get<any>(url);
  }

  getForClusters(): Observable<MetricsList> {
    const url = `${this.configService.config.rorApi}/v1/metrics/clusters`;
    return this.httpClient.get<MetricsList>(url);
  }

  getForClustersByWorkspace(workspaceName: string): Observable<MetricsList> {
    const url = `${this.configService.config.rorApi}/v1/metrics/clusters/workspace/${workspaceName}`;
    return this.httpClient.get<MetricsList>(url);
  }

  getForClusterId(clusterId: string): Observable<MetricData> {
    const url = `${this.configService.config.rorApi}/v1/metrics/cluster/${clusterId}`;
    return this.httpClient.get<MetricData>(url);
  }

  getForClusterByProperty(property: string): Observable<MetricsCustom> {
    const url = `${this.configService.config.rorApi}/v1/metrics/custom/cluster/${property}`;
    return this.httpClient.get<MetricsCustom>(url);
  }
}
