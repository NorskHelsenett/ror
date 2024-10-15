export interface FilterMetadata {
  field: string;
  value: any;
  matchMode: string;
}

export interface SortMetadata {
  sortField: string;
  sortOrder: number;
}

export interface Filter {
  skip: number;
  limit: number;
  sort?: SortMetadata[];
  filters?: FilterMetadata[];
}
