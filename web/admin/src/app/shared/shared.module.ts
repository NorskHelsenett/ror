import { TranslateModule } from '@ngx-translate/core';
import { NgModule } from '@angular/core';
import { CommonModule, NgOptimizedImage } from '@angular/common';
import { sharedPipes } from './pipes';
import { sharedComponents } from './components';
import { sharedIcons } from './icons';
import { sharedServices } from './services';

import { ChipModule } from 'primeng/chip';
import { TagModule } from 'primeng/tag';
import { TooltipModule } from 'primeng/tooltip';
import { MenuModule } from 'primeng/menu';
import { ButtonModule } from 'primeng/button';

@NgModule({
  declarations: [...sharedPipes, ...sharedComponents, ...sharedIcons],
  providers: [...sharedServices],
  imports: [CommonModule, TranslateModule, ChipModule, TagModule, TooltipModule, NgOptimizedImage, MenuModule, ButtonModule],
  exports: [...sharedPipes, ...sharedComponents, ...sharedIcons],
})
export class SharedModule {}
