import { Datacenter } from './datacenter';

export interface Workspace {
  id: string;
  name: string;
  projectId: string;
  datacenter: Datacenter;
}
