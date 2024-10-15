import { TableModule } from 'primeng/table';
import { TranslateModule } from '@ngx-translate/core';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PricesComponent } from './prices.component';
import { priceComponents } from './components';
import { PricesRoutingModule } from './prices.routing';
import { SharedModule } from '../shared/shared.module';
import { ReactiveFormsModule } from '@angular/forms';

@NgModule({
  declarations: [PricesComponent, ...priceComponents],
  imports: [CommonModule, TranslateModule, PricesRoutingModule, SharedModule, ReactiveFormsModule, TableModule],
})
export class PricesModule {}
