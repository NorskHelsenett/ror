export interface Ingress {
  health: string;
  name: string;
  namespace: string;
  class: string;
  hosts: string[];
  addresses: string[];
}
