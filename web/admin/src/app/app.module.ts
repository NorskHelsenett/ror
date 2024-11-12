import { APP_INITIALIZER, NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HttpClient, HTTP_INTERCEPTORS, provideHttpClient, withInterceptorsFromDi } from '@angular/common/http';

import { MissingTranslationHandler, TranslateLoader, TranslateModule } from '@ngx-translate/core';
import { CustomMissingTranslationHandler, HttpLoaderFactory } from './core/i18n/i18nfunctions';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { ServiceWorkerModule } from '@angular/service-worker';
import { OAuthModule, OAuthStorage } from 'angular-oauth2-oidc';
import { AuthInterceptor } from './core/interceptors/auth-interceptor';
import { ClipboardModule } from 'ngx-clipboard';
import { ToastModule } from 'primeng/toast';

import { registerLocaleData } from '@angular/common';
import localeNo from '@angular/common/locales/no';
import localeNoExtra from '@angular/common/locales/extra/no';
import localeEn from '@angular/common/locales/en';
import localeEnExtra from '@angular/common/locales/extra/en';
import { CoreModule } from './core/core.modules';
import { SharedModule } from './shared/shared.module';
import { MessageService } from 'primeng/api';
import { ConfigService, configFactory } from './core/services/config.service';
import { provideHighlightOptions } from 'ngx-highlightjs';
import { InMemoryScrollingFeature, InMemoryScrollingOptions, withInMemoryScrolling } from '@angular/router';
import { environment } from '../environments/environment';

registerLocaleData(localeNo, localeNoExtra);
registerLocaleData(localeEn, localeEnExtra);

export function storageFactory(): OAuthStorage {
  return localStorage;
}

const scrollConfig: InMemoryScrollingOptions = {
  scrollPositionRestoration: 'top',
  anchorScrolling: 'enabled',
};

const inMemoryScrollingFeature: InMemoryScrollingFeature = withInMemoryScrolling(scrollConfig);

@NgModule({
  declarations: [AppComponent],
  bootstrap: [AppComponent],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    AppRoutingModule,
    OAuthModule.forRoot(),
    TranslateModule.forRoot({
      missingTranslationHandler: { provide: MissingTranslationHandler, useClass: CustomMissingTranslationHandler },
      useDefaultLang: false,
      loader: {
        provide: TranslateLoader,
        useFactory: HttpLoaderFactory,
        deps: [HttpClient],
      },
    }),
    ClipboardModule,
    CoreModule,
    SharedModule,
    ToastModule,
    ServiceWorkerModule.register('ngsw-worker.js', {
      enabled: environment.production,
      // Register the ServiceWorker as soon as the application is stable
      // or after 30 seconds (whichever comes first).
      registrationStrategy: 'registerWhenStable:30000',
    }),
  ],
  providers: [
    //provideRouter(routes, inMemoryScrollingFeature, withViewTransitions(), withDebugTracing()),
    {
      provide: HTTP_INTERCEPTORS,
      useClass: AuthInterceptor,
      multi: true,
    },
    {
      provide: APP_INITIALIZER,
      useFactory: configFactory,
      multi: true,
      deps: [ConfigService],
    },
    MessageService,
    provideHighlightOptions({
      lineNumbersOptions: {
        singleLine: false,
      },
      coreLibraryLoader: () => import('highlight.js/lib/core'),
      lineNumbersLoader: () => import('ngx-highlightjs/line-numbers'),
      themePath: 'assets/styles/highlight/tokyo-night-light.css',
      languages: {
        bash: () => import('highlight.js/lib/languages/bash'),
        json: () => import('highlight.js/lib/languages/json'),
      },
    }),
    { provide: OAuthStorage, useFactory: storageFactory },
    provideHttpClient(withInterceptorsFromDi()),
  ],
})
export class AppModule {}
