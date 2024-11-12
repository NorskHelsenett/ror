import { Injectable } from '@angular/core';
import { ClusterEnvironment } from '../models/clusterEnvironment';

@Injectable({
  providedIn: 'root',
})
export class EnvironmentService {
  getEnvironments(): any {
    return [
      {
        name: ClusterEnvironment[ClusterEnvironment.Development],
        value: 'dev',
      },
      {
        name: ClusterEnvironment[ClusterEnvironment.Testing],
        value: 'test',
      },
      {
        name: ClusterEnvironment[ClusterEnvironment.QA],
        value: 'qa',
      },
      {
        name: ClusterEnvironment[ClusterEnvironment.Production],
        value: 'prod',
      },
    ];
  }
}
