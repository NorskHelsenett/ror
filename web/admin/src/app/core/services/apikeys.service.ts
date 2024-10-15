import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { Filter } from '../models/apiFilter';
import { PaginationResult } from '../models/paginatedResult';
import { ApiKey } from '../models/apikey';
import { ConfigService } from './config.service';

@Injectable({
  providedIn: 'root',
})
export class ApikeysService {
  constructor(
    private httpClient: HttpClient,
    private configService: ConfigService,
  ) {}

  getById(id: string): Observable<ApiKey> {
    let url = `${this.configService.config.rorApi}/v1/apikeys/${id}`;
    return this.httpClient.get<ApiKey>(url);
  }

  getByFilter(filter: Filter): Observable<PaginationResult<ApiKey>> {
    if (!filter) {
      return of(null);
    }

    let url = `${this.configService.config.rorApi}/v1/apikeys/filter`;
    return this.httpClient.post<PaginationResult<ApiKey>>(url, filter);
  }

  delete(apikeyId: string): Observable<boolean> {
    let url = `${this.configService.config.rorApi}/v1/apikeys/${apikeyId}`;
    return this.httpClient.delete<boolean>(url);
  }

  create(apikey: ApiKey): Observable<string> {
    let url = `${this.configService.config.rorApi}/v1/apikeys`;
    return this.httpClient.post<string>(url, apikey);
  }
}
