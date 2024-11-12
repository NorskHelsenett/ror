import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { PolicyReportGlobal } from '../models/policyReport';
import { ConfigService } from './config.service';

@Injectable({
  providedIn: 'root',
})
export class PolicyReportsService {
  constructor(
    private httpClient: HttpClient,
    private configService: ConfigService,
  ) {}

  getPolicyReportsGlobal(type: string, clusterID: string): Observable<PolicyReportGlobal[]> {
    let url: string = `${this.configService.config.rorApi}/v1/clusters/views/policyreports?type=${type}&clusterid=${clusterID}`;
    return this.httpClient.get<PolicyReportGlobal[]>(url);
  }
}
