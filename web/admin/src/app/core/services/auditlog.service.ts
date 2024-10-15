import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Filter } from '../models/apiFilter';
import { AuditLog } from '../models/auditlog';
import { PaginationResult } from '../models/paginatedResult';
import { ConfigService } from './config.service';

@Injectable({
  providedIn: 'root',
})
export class AuditlogService {
  constructor(
    private httpClient: HttpClient,
    private configService: ConfigService,
  ) {}

  getAll(): Observable<AuditLog[]> {
    let url = `${this.configService.config.rorApi}/v1/auditlogs`;
    return this.httpClient.get<AuditLog[]>(url);
  }

  getById(id: string): Observable<AuditLog> {
    let url = `${this.configService.config.rorApi}/v1/auditlogs/${id}`;
    return this.httpClient.get<AuditLog>(url);
  }

  getByFilter(filter: Filter): Observable<PaginationResult<AuditLog>> {
    let url = `${this.configService.config.rorApi}/v1/auditlogs/filter`;
    return this.httpClient.post<PaginationResult<AuditLog>>(url, filter);
  }

  getMetadata(): Observable<Map<string, string[]>> {
    let url = `${this.configService.config.rorApi}/v1/auditlogs/metadata`;
    return this.httpClient.get<Map<string, string[]>>(url);
  }
}
