import { ErrorComponent } from './error.component';
import { Routes, RouterModule } from '@angular/router';

import * as errorPages from './pages/index';

const routes: Routes = [
  {
    path: '',
    component: ErrorComponent,
    children: [
      {
        path: '',
        redirectTo: '404',
        pathMatch: 'full',
      },
      {
        path: '401',
        component: errorPages.UnauthorizedComponent,
      },
      {
        path: '403',
        component: errorPages.ForbiddenComponent,
      },
      {
        path: '404',
        component: errorPages.NotFoundComponent,
      },
      {
        path: '500',
        component: errorPages.ServerErrorComponent,
      },
    ],
  },
];

export const ErrorRoutingModule = RouterModule.forChild(routes);
