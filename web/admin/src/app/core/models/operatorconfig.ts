export interface OperatorConfig {
  id: string;
  apiVersion: string;
  kind: string;
  spec: OperatorSpec;
}

export interface OperatorSpec {
  imagePostfix: string;
  tasks: OperatorTask[];
}

export interface OperatorTask {
  index: number;
  name: string;
  version: string;
  runOnce: boolean;
}
