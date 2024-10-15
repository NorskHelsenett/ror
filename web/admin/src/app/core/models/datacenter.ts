import { Location } from './location';

export interface Datacenter {
  id: string;
  name: string;
  provider: string;
  apiEndpoint: string;
  location: Location;
}
