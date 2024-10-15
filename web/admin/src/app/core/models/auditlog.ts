import { User } from './user';

export interface AuditlogMetadata {
  msg: string;
  timestamp: Date;
  category: string;
  action: string;
  user: User;
}

export interface AuditLog {
  id: string;
  metadata: AuditlogMetadata;
  data: any;
}
