import { Injectable } from '@angular/core';

import { ColumnDefinition } from '../../resources/models/columnDefinition';
import { LazyLoadEvent } from 'primeng/api';
import { ResourceQueryFilter, ResourceQueryOrder } from '../../core/models/resources-v2';

@Injectable({
  providedIn: 'root',
})
export class FilterService {
  getFilters(event: LazyLoadEvent, columnDefinitions: ColumnDefinition[]): any {
    if (!event || !event.filters) {
      return;
    }

    let filters: ResourceQueryFilter[] = [];
    Object.entries(event?.filters).forEach((entry: any) => {
      let value = entry[1]?.value;
      if (!value) {
        return;
      }

      let searchField = '';
      let field = String(entry[0])?.toLowerCase();
      if (field === 'apiversion') {
        field = 'apiVersion';
        searchField = 'typemeta.apiversion';
      } else if (field === 'kind') {
        searchField = 'typemeta.kind';
      } else {
        searchField = field;
      }
      let operator = this.getOperator(entry[1]?.matchMode);

      const columnDefinition = columnDefinitions.find((column) => column.field === field);

      if (!columnDefinition) {
        console.error('missing columnDefinition', field, searchField);
      }

      let type = 'string';
      if (columnDefinition?.type === 'numeric') {
        type = 'int';
      } else if (columnDefinition?.type === 'boolean') {
        type = 'bool';
      }
      filters.push({ field: searchField, value: value, type: type, operator: operator });
    });

    return filters;
  }

  getOrder(event: LazyLoadEvent, columnDefinitions: ColumnDefinition[]): ResourceQueryOrder[] {
    if (!event) {
      return [];
    }
    let order: ResourceQueryOrder[] = [];

    if (!event?.multiSortMeta) {
      return order;
    }

    let count = 1;
    event?.multiSortMeta.forEach((entry: any) => {
      console.log('entry', entry);
      let searchField = '';
      let field = String(entry?.field)?.toLowerCase();
      if (field === 'apiversion') {
        field = 'apiVersion';
        searchField = 'typemeta.apiversion';
      } else if (field === 'kind') {
        searchField = 'typemeta.kind';
      } else {
        searchField = field;
      }

      order.push({
        field: searchField,
        descending: entry?.order === 1 ? false : true,
        index: count,
      });
      count++;
    });

    return order;
  }

  private getOperator(matchMode: string) {
    let operator = 'regexp';

    if (matchMode === 'equals') {
      operator = 'eq';
    } else if (matchMode === 'notEquals') {
      operator = 'ne';
    }

    return operator;
  }
}
