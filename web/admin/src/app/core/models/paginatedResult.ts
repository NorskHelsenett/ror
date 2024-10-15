export interface PaginationResult<T> {
  data: T[];
  dataCount: number;
  totalCount: number;
  offset: number;
}
