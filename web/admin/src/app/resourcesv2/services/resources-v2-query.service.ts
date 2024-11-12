import { Injectable } from '@angular/core';
import { ResourceQuery } from '../../core/models/resource-query';

@Injectable({
  providedIn: 'root',
})
export class ResourcesV2QueryService {
  private query: ResourceQuery | undefined;

  getQuery(): ResourceQuery {
    return this.query ?? new ResourceQuery();
  }

  setQuery(query: ResourceQuery): void {
    this.query = query;
  }

  updateQuery(query: any): void {
    var obj = { ...this.query, ...query };
    this.query = new ResourceQuery(obj);
  }

  clearQuery(): void {
    this.query = undefined;
  }
}
