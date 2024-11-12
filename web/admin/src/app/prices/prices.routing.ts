import { PricesComponent } from './prices.component';
import { Routes, RouterModule } from '@angular/router';

const routes: Routes = [
  {
    path: '',
    component: PricesComponent,
  },
];

export const PricesRoutingModule = RouterModule.forChild(routes);
