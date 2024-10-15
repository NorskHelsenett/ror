import { Routes, RouterModule } from '@angular/router';
import { ClustersComponent } from './clusters.component';

import * as clustersPages from './pages';

const routes: Routes = [
  {
    path: '',
    component: ClustersComponent,
  },
  {
    path: 'create',
    component: clustersPages.ClusterNewComponent,
  },
  {
    path: ':id',
    component: clustersPages.ClusterDetailsComponent,
  },
  {
    path: ':id/ingresses/:ingressid',
    component: clustersPages.IngressDetailsComponent,
  },
];

export const ClustersRoutingModule = RouterModule.forChild(routes);
