export interface Price {
  id: string;
  provider: string;
  machineClass: string;
  cpu: number;
  memoryBytes: number;
  memory: number;
  price: number;
  from: Date;
  to: Date;
}
