import { TranslateModule } from '@ngx-translate/core';
import { NgModule } from '@angular/core';
import { CommonModule, NgOptimizedImage } from '@angular/common';
import { AboutComponent } from './about.component';
import { AboutRoutingModule } from './about.routing';
import { SharedModule } from '../shared/shared.module';
import { HealthTableComponent } from './components/health-table/health-table.component';

@NgModule({
  declarations: [AboutComponent, HealthTableComponent],
  imports: [CommonModule, AboutRoutingModule, TranslateModule, SharedModule, NgOptimizedImage],
})
export class AboutModule {}
