import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { ConfigService } from './config.service';
import { Observable } from 'rxjs';
import { ResourceQuery, ResourceSet } from '@rork8s/ror-resources/models';

@Injectable({
  providedIn: 'root',
})
export class ResourcesService {
  constructor(
    private httpClient: HttpClient,
    private configService: ConfigService,
  ) {}

  getResources<T>(query: ResourceQuery): Observable<ResourceSet> {
    const queryString = JSON.stringify(query);
    const b64Query = btoa(queryString);
    const url = `${this.configService.config.rorApi}/v2/resources?query=${b64Query}`;
    return this.httpClient.get<ResourceSet>(url);
  }

  getResource<T>(uid: string, ownerScope: string, ownerSubject: string, kind: string, apiVersion: string): Observable<T> {
    const url = `${this.configService.config.rorApi}/v2/resources/uid/${uid}?ownerScope=${ownerScope}&ownerSubject=${ownerSubject}&kind=${kind}&apiversion=${apiVersion}`;
    return this.httpClient.get<T>(url);
  }

  createResource<T>(resourceSet: ResourceSet): Observable<T> {
    const url = `${this.configService.config.rorApi}/v2/resources`;
    return this.httpClient.post<T>(url, resourceSet);
  }

  updateResource<T>(resourceSet: ResourceSet, uid: string): Observable<T> {
    const url = `${this.configService.config.rorApi}/v2/resources/uid/${uid}`;
    return this.httpClient.put<T>(url, resourceSet);
  }

  deleteResource<T>(uid: string): Observable<T> {
    const url = `${this.configService.config.rorApi}/v2/resources/uid/${uid}`;
    return this.httpClient.delete<T>(url);
  }
}
