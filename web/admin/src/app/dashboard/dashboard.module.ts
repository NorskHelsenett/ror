import { TooltipModule } from 'primeng/tooltip';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { TranslateModule } from '@ngx-translate/core';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DashboardComponent } from './dashboard.component';
import { CoreModule } from '../core/core.modules';
import { SharedModule } from '../shared/shared.module';

import { ClustersModule } from '../clusters/clusters.module';
import { DashboardRoutingModule } from './dashboard.routing';
import { TableModule } from 'primeng/table';
import { ConfirmDialogModule } from 'primeng/confirmdialog';

@NgModule({
  declarations: [DashboardComponent],
  imports: [
    CommonModule,
    DashboardRoutingModule,
    TranslateModule,
    CoreModule,
    SharedModule,
    FormsModule,
    ReactiveFormsModule,
    TooltipModule,
    ClustersModule,
    TableModule,
    ConfirmDialogModule,
  ],
  //providers: [provideRouter(routes, withViewTransitions())],
})
export class DashboardModule {}
