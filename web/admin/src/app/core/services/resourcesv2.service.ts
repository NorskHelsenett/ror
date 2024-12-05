import { inject, Injectable } from '@angular/core';
import { ConfigService } from './config.service';
import { HttpClient } from '@angular/common/http';
import { map, Observable } from 'rxjs';
import { ResourceSet, ResourceQuery } from '@rork8s/ror-resources/models';

@Injectable({
  providedIn: 'root',
})
export class Resourcesv2Service {
  private configService = inject(ConfigService);
  private httpClient = inject(HttpClient);

  constructor() {}

  getResources(resourceQuery: ResourceQuery): Observable<ResourceSet> {
    const queryBase64 = btoa(JSON.stringify(resourceQuery));
    const url = `${this.configService.config.rorApi}/v2/resources?query=${queryBase64}`;
    return this.httpClient.get<ResourceSet>(url).pipe(
      map((resources) => {
        if (!resources) {
          return null;
        }
        return resources;
      }),
    );
  }

  getResourcesById(uid: string): Observable<ResourceSet> {
    const url = `${this.configService.config.rorApi}/v2/resources/uid/${uid}`;
    return this.httpClient.get<ResourceSet>(url).pipe(
      map((resources) => {
        if (!resources) {
          return null;
        }
        return resources;
      }),
    );
  }
}
