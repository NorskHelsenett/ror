import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuard } from './core/guards/auth.guard';

export const routes: Routes = [
  {
    path: 'error',
    loadChildren: () => import('./error/error.module').then((m) => m.ErrorModule),
  },
  {
    path: 'auth/callback',
    loadChildren: () => import('./auth-callback/auth-callback.module').then((m) => m.AuthCallbackModule),
  },
  {
    path: '',
    canActivate: [AuthGuard],
    loadChildren: () => import('./layout/layout.module').then((m) => m.LayoutModule),
  },
  { path: '**', redirectTo: 'error/404' },
];

@NgModule({
  imports: [
    RouterModule.forRoot(routes, {
      anchorScrolling: 'enabled',
      scrollPositionRestoration: 'top',
      onSameUrlNavigation: 'ignore',
    }),
  ],
  exports: [RouterModule],
})
export class AppRoutingModule {}
