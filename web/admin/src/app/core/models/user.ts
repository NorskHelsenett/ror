export interface User {
  email: string;
  name: string;
  groups: Array<string>;
  isAdmin: boolean;
}
