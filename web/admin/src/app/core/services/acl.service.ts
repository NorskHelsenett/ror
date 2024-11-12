import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, catchError, map, of } from 'rxjs';
import { Filter } from '../models/apiFilter';
import { PaginationResult } from '../models/paginatedResult';
import { AclV2 } from '../models/aclv2';
import { AclAccess, AclScopes } from '../models/acl-scopes';
import { ConfigService } from './config.service';

@Injectable({
  providedIn: 'root',
})
export class AclService {
  constructor(
    private httpClient: HttpClient,
    private configService: ConfigService,
  ) {}

  getById(id: string): Observable<AclV2> {
    let url = `${this.configService.config.rorApi}/v1/acl/${id}`;
    return this.httpClient.get<AclV2>(url);
  }

  getScopes(): Observable<string[]> {
    let url = `${this.configService.config.rorApi}/v1/acl/scopes`;
    return this.httpClient.get<string[]>(url);
  }

  getByFilter(filter: Filter): Observable<PaginationResult<AclV2>> {
    if (!filter) {
      return of(null);
    }

    let url = `${this.configService.config.rorApi}/v1/acl/filter`;
    return this.httpClient.post<PaginationResult<AclV2>>(url, filter);
  }

  create(acl: AclV2): Observable<AclV2> {
    let url = `${this.configService.config.rorApi}/v1/acl`;
    acl.version = 2;
    return this.httpClient.post<AclV2>(url, acl);
  }

  update(id: string, acl: AclV2): Observable<AclV2> {
    let url = `${this.configService.config.rorApi}/v1/acl/${id}`;
    acl.version = 2;
    return this.httpClient.put<AclV2>(url, acl);
  }

  delete(apikeyId: string): Observable<boolean> {
    let url = `${this.configService.config.rorApi}/v1/acl/${apikeyId}`;
    return this.httpClient.delete<boolean>(url);
  }

  check(scope: AclScopes, subject: AclScopes, access: AclAccess): Observable<boolean> {
    let url = `${this.configService.config.rorApi}/v1/acl/${scope}/${subject}/${access}`;
    return this.httpClient.head<boolean>(url).pipe(
      catchError((error: any) => {
        if (error?.status === 401 || error?.status === 403) {
          return of(false);
        }
        throw error;
      }),
      map((data: any) => {
        if (data === true || data === false) {
          return data;
        }
        if (data === undefined || data === null) {
          return true;
        }
      }),
    );
  }
}
