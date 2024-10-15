import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { TranslateModule } from '@ngx-translate/core';
import { SharedModule } from '../shared/shared.module';
import { OrdersRoutingModule } from './orders.routing';
import { orderssComponents } from './components';
import { orderssPages } from './pages';
import { TableModule } from 'primeng/table';
import { DropdownModule } from 'primeng/dropdown';
import { FormsModule } from '@angular/forms';

@NgModule({
  declarations: [...orderssPages, ...orderssComponents],
  imports: [CommonModule, OrdersRoutingModule, TranslateModule, SharedModule, TableModule, DropdownModule, FormsModule],
})
export class OrdersModule {}
