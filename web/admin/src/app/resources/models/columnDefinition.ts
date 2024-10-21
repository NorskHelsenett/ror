export interface ColumnDefinition {
  field: string;
  header: string;
  type: 'text' | 'numeric' | 'boolean' | 'date' | 'array' | 'object';
  enabled: boolean;
}
