import { Routes, RouterModule } from '@angular/router';
import { ReleaseNotesComponent } from './release-notes.component';

const routes: Routes = [
  {
    path: '',
    component: ReleaseNotesComponent,
  },
];

export const ReleaseNotesRoutingModule = RouterModule.forChild(routes);
