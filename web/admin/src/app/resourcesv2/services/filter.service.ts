import { Injectable } from '@angular/core';

import { ResourceQueryFilter } from '../../core/models/resource-query';
import { ColumnDefinition } from '../../resources/models/columnDefinition';

@Injectable({
  providedIn: 'root',
})
export class FilterService {
  getFilters(event: any, columnDefinitions: ColumnDefinition[]): any {
    if (!event || !event.filters || event.filters.length === 0) {
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
