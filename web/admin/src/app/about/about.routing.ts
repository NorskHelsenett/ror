import { AboutComponent } from './about.component';
import { Routes, RouterModule } from '@angular/router';

const routes: Routes = [
  {
    path: '',
    component: AboutComponent,
  },
];

export const AboutRoutingModule = RouterModule.forChild(routes);
