import { Routes, RouterModule } from '@angular/router';
import { WorkspacesComponent } from './workspaces.component';

import * as workspacesPages from './pages';

const routes: Routes = [
  {
    path: '',
    component: WorkspacesComponent,
  },
  {
    path: ':workspaceName',
    component: workspacesPages.WorkspaceDetailsComponent,
  },
  {
    path: ':workspaceName/edit',
    component: workspacesPages.WorkspaceNewComponent,
  },
];

export const WorkspacesRoutingModule = RouterModule.forChild(routes);
