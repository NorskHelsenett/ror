import { Injectable } from '@angular/core';
import { ResourceQuery } from '@rork8s/ror-resources/models';

@Injectable({
  providedIn: 'root',
})
export class ResourcesV2QueryService {
  private query: ResourceQuery | undefined;

  getQuery(): ResourceQuery {
    return this.query ?? {};
  }

  setQuery(query: ResourceQuery): void {
    this.query = query;
  }

  updateQuery(query: any): void {
    var obj = { ...this.query, ...query };
    this.query = obj;
  }

  clearQuery(): void {
    this.query = undefined;
  }
}
