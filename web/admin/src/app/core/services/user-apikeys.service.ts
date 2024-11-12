import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { Filter } from '../models/apiFilter';
import { HttpClient } from '@angular/common/http';
import { ApiKey } from '../models/apikey';
import { PaginationResult } from '../models/paginatedResult';
import { ConfigService } from './config.service';

@Injectable({
  providedIn: 'root',
})
export class UserApikeysService {
  constructor(
    private httpClient: HttpClient,
    private configService: ConfigService,
  ) {}

  getByFilter(filter: Filter): Observable<PaginationResult<ApiKey>> {
    if (!filter) {
      return of(null);
    }

    let url = `${this.configService.config.rorApi}/v1/users/self/apikeys/filter`;
    return this.httpClient.post<PaginationResult<ApiKey>>(url, filter);
  }

  create(apikey: ApiKey): Observable<string> {
    let url = `${this.configService.config.rorApi}/v1/users/self/apikeys`;
    return this.httpClient.post<string>(url, apikey);
  }

  delete(apikeyId: string): Observable<boolean> {
    let url = `${this.configService.config.rorApi}/v1/users/self/apikeys/${apikeyId}`;
    return this.httpClient.delete<boolean>(url);
  }
}
