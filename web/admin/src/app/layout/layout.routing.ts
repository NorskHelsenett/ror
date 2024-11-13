import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LayoutComponent } from './layout.component';

export const routes: Routes = [
  {
    path: '',
    component: LayoutComponent,
    children: [
      {
        path: '',
        loadChildren: () => import('../dashboard/dashboard.module').then((m) => m.DashboardModule),
      },
      {
        path: 'datacenters',
        loadChildren: () => import('../datacenters/datacenters.module').then((m) => m.DatacentersModule),
      },
      {
        path: 'workspaces',
        loadChildren: () => import('../workspaces/workspaces.module').then((m) => m.WorkspacesModule),
      },
      {
        path: 'clusters',
        loadChildren: () => import('../clusters/clusters.module').then((m) => m.ClustersModule),
      },
      {
        path: 'create',
        loadChildren: () => import('../create/create.module').then((m) => m.CreateModule),
      },
      {
        path: 'metrics',
        loadChildren: () => import('../metrics/metrics.module').then((m) => m.MetricsModule),
      },
      {
        path: 'orders',
        loadChildren: () => import('../orders/orders.module').then((m) => m.OrdersModule),
      },
      {
        path: 'prices',
        loadChildren: () => import('../prices/prices.module').then((m) => m.PricesModule),
      },
      {
        path: 'resources',
        loadChildren: () => import('../resources/resources.module').then((m) => m.ResourcesModule),
      },
      {
        path: 'resourcesv2',
        loadChildren: () => import('../resourcesv2/resourcesv2.module').then((m) => m.ResourcesV2Module),
      },
      {
        path: 'userprofile',
        loadChildren: () => import('../userprofile/userprofile.module').then((m) => m.UserprofileModule),
      },
      {
        path: 'about',
        loadChildren: () => import('../about/about.module').then((m) => m.AboutModule),
      },
      {
        path: 'admin',
        loadChildren: () => import('../admin/admin.module').then((m) => m.AdminModule),
      },
      {
        path: 'releasenotes',
        loadChildren: () => import('../release-notes/release-notes.module').then((m) => m.ReleaseNotesModule),
      },
      { path: '**', redirectTo: 'error/404' },
    ],
  },
  { path: '**', redirectTo: 'error/404' },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class LayoutRoutingModule {}
