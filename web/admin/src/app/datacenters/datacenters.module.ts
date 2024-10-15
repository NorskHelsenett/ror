import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { datacentersPages } from './pages';
import { TranslateModule } from '@ngx-translate/core';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DatacentersComponent } from './datacenters.component';
import { datacentersComponents } from './components';
import { SharedModule } from '../shared/shared.module';
import { RouterModule } from '@angular/router';

@NgModule({
  declarations: [DatacentersComponent, ...datacentersPages, ...datacentersComponents],
  imports: [CommonModule, TranslateModule, SharedModule, FormsModule, ReactiveFormsModule, RouterModule],
})
export class DatacentersModule {}
