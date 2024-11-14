import { Routes, RouterModule } from '@angular/router';
import { ResourcesComponent } from './pages/resources/resources.component';

const routes: Routes = [
  {
    path: '',
    component: ResourcesComponent,
  },
  { path: '**', redirectTo: 'error/404' },
];

export const ResourcesV2ModulesRoutingModule = RouterModule.forChild(routes);
