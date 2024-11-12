import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { CreateClusterComponent } from './create-cluster.component';
import * as createClusterComponents from './components';
const routes: Routes = [
  {
    path: '',
    component: CreateClusterComponent,
    children: [
      { path: '', redirectTo: 'new', pathMatch: 'full' },
      {
        path: 'new',
        component: createClusterComponents.LocationStepComponent,
      },
      {
        path: 'resources',
        component: createClusterComponents.ResourcesStepComponent,
      },
      {
        path: 'metadata',
        component: createClusterComponents.MetadataStepComponent,
      },
      {
        path: 'summary',
        component: createClusterComponents.SummaryStepComponent,
      },
    ],
  },
  { path: '**', redirectTo: 'error/404' },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class CreateClusterRoutingModule {}
