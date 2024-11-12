import { TranslateModule } from '@ngx-translate/core';
import { NgModule } from '@angular/core';
import { CommonModule, NgOptimizedImage } from '@angular/common';
import { LayoutComponent } from './layout.component';
import { LayoutRoutingModule, routes } from './layout.routing';

import { ToastModule } from 'primeng/toast';

import { ScrollTopModule } from 'primeng/scrolltop';
import { SharedModule } from '../shared/shared.module';
import { SantaComponent } from '../shared/components/santa/santa.component';
import { DesemberGiftComponent } from '../shared/components/desember-gift/desember-gift.component';
import {
  InMemoryScrollingFeature,
  InMemoryScrollingOptions,
  RouterModule,
  RouterOutlet,
  provideRouter,
  withDebugTracing,
  withInMemoryScrolling,
  withViewTransitions,
} from '@angular/router';

const scrollConfig: InMemoryScrollingOptions = {
  scrollPositionRestoration: 'top',
  anchorScrolling: 'enabled',
};

const inMemoryScrollingFeature: InMemoryScrollingFeature = withInMemoryScrolling(scrollConfig);

@NgModule({
  declarations: [LayoutComponent],
  imports: [
    CommonModule,
    LayoutRoutingModule,
    TranslateModule,
    ScrollTopModule,
    ToastModule,
    SharedModule,
    NgOptimizedImage,
    SantaComponent,
    DesemberGiftComponent,
  ],
  //providers: [provideRouter(routes, inMemoryScrollingFeature, withViewTransitions())],
})
export class LayoutModule {}
