import { Injectable } from '@angular/core';
import { LazyLoadEvent } from 'primeng/api';
import { Filter, FilterMetadata, SortMetadata } from '../models/apiFilter';

@Injectable({
  providedIn: 'root',
})
export class FilterService {
  mapFilter(event: LazyLoadEvent): Filter {
    if (!event) {
      return { skip: 0, limit: 25, filters: [], sort: [] };
    }

    let filters: FilterMetadata[] = [];
    Object.entries(event?.filters).forEach((entry) => {
      filters.push({ field: entry[0]?.toLowerCase(), value: entry[1]?.value, matchMode: entry[1]?.matchMode });
    });

    let sort: SortMetadata[] = [];
    if (event?.sortField && event?.sortOrder) {
      sort.push({ sortField: event?.sortField?.toLowerCase(), sortOrder: event.sortOrder });
    }
    return { skip: event.first, limit: event.rows, filters: filters, sort: sort };
  }
}
