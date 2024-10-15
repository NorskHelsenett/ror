import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AdminCreateGuard } from '../core/guards/admin-create.guard';

const routes: Routes = [
  {
    path: 'cluster',
    canActivate: [AdminCreateGuard],
    loadChildren: () => import('./create-cluster/create-cluster.module').then((m) => m.CreateClusterModule),
  },
  { path: '**', redirectTo: 'error/404' },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class CreateRoutingModule {}
