export interface ApiKey {
  id: string;
  identifier: string;
  displayName: string;
  type: string;
  readOnly: boolean;
  expires: Date;
  created: Date;
  lastUsed: Date;
}
