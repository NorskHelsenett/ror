import { Routes, RouterModule } from '@angular/router';
import * as resourcesPages from './pages';

const routes: Routes = [
  {
    path: '',
    component: resourcesPages.ResourcesComponent,
  },
  {
    path: ':apiVersion/:kind/:scope/:subject/:uid',
    component: resourcesPages.ResourceDetailsComponent,
  },
  { path: '**', redirectTo: 'error/404' },
];

export const ResourcesModulesRoutingModule = RouterModule.forChild(routes);
