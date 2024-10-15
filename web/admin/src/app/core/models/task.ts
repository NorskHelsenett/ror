export interface Task {
  id: string;
  name: string;
  config: TaskSpec;
}

export interface TaskSpec {
  imageName: string;
  cmd: string;
  envVars: KeyValue[];
  backoffLimit: number;
  timeoutInSeconds: number;
  version: string;
  index: number;
  secret: any;
  scripts: Scripts;
}

export interface Scripts {
  scriptDirectory: string;
  fileNameAndData: FileNameAndData[];
}

export interface FileNameAndData {
  filename: string;
  data: string;
}

export interface TaskSecret {
  filename: string;
  path: string;
  data: any;
  gitSource: TaskGitSource;
  vaultSource: TaskVaultSource;
}

export interface TaskGitSource {
  token: string;
  user: string;
  repository: string;
  branch: string;
  projectId: number;
}

export interface TaskVaultSource {
  type: string;
  vaultPath: string;
}

export interface KeyValue {
  key: string;
  value: string;
}
