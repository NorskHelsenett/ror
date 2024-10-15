import { SharedModule } from '../shared/shared.module';
import { TranslateModule } from '@ngx-translate/core';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MetricsComponent } from './metrics.component';
import { MetricsRoutingModule } from './metrics.routing';
import { ChartModule } from 'primeng/chart';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CoreModule } from '../core/core.modules';
import { DialogModule } from 'primeng/dialog';

@NgModule({
  declarations: [MetricsComponent],
  imports: [
    CommonModule,
    MetricsRoutingModule,
    TranslateModule,
    CoreModule,
    SharedModule,
    FormsModule,
    ReactiveFormsModule,
    SharedModule,
    ChartModule,
    DialogModule,
  ],
})
export class MetricsModule {}
