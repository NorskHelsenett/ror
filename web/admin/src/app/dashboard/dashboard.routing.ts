import { Routes, RouterModule } from '@angular/router';
import { DashboardComponent } from './dashboard.component';

import * as clustersPages from '../clusters/pages';

export const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    component: DashboardComponent,
  },
  {
    path: 'cluster/create',
    component: clustersPages.ClusterNewComponent,
  },
  {
    path: 'cluster/:id',
    component: clustersPages.ClusterDetailsComponent,
  },
  {
    path: 'cluster/:id/ingresses/:ingressid',
    component: clustersPages.IngressDetailsComponent,
  },
  { path: '**', redirectTo: 'error/404' },
];

export const DashboardRoutingModule = RouterModule.forChild(routes);
