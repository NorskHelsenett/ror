import { ChangeDetectionStrategy, ChangeDetectorRef, Component, OnDestroy, OnInit } from '@angular/core';
import { OAuthService } from 'angular-oauth2-oidc';
import { Subscription, filter, tap } from 'rxjs';

import { TranslateService } from '@ngx-translate/core';

import { ThemeService } from './core/services/theme.service';
import { Title } from '@angular/platform-browser';
import { AuthService } from './core/services/auth.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AppComponent implements OnInit, OnDestroy {
  isDark = false;

  private subscriptions = new Subscription();

  constructor(
    private changeDetector: ChangeDetectorRef,
    private oauthService: OAuthService,
    private authService: AuthService,
    private themeService: ThemeService,
    private titleService: Title,
    private translateService: TranslateService,
  ) {
    this.oauthService.configure(this.authService.authConfig);
    this.oauthService.loadDiscoveryDocumentAndLogin();
    this.oauthService.setupAutomaticSilentRefresh();
    this.oauthService.events.pipe(filter((e) => e?.type === 'token_received')).subscribe((_) => {
      this.oauthService.loadUserProfile();
    });
  }

  ngOnInit(): void {
    this.subscriptions.add(
      this.themeService.isDark.subscribe((value) => {
        this.isDark = value;
        this.changeDetector.detectChanges();
      }),
    );

    this.subscriptions.add(
      this.translateService.onLangChange
        .pipe(
          tap((langSettings) => {
            this.changeTitle(langSettings.lang);
          }),
        )
        .subscribe(),
    );

    this.translateService.addLangs(['en', 'no']);
    this.translateService.setDefaultLang('no');

    let lang = 'en';
    const userLang = localStorage.getItem('language');
    if (userLang && userLang.length > 0) {
      lang = userLang;
    } else {
      const browserLang = this.translateService.getBrowserLang();
      lang = browserLang?.match(/en|no/) ? browserLang : 'en';
    }

    this.translateService.use(lang);

    this.changeDetector.detectChanges();
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }

  get userName(): string | null | object {
    const claims = this.oauthService.getIdentityClaims();
    if (!claims) {
      return null;
    }
    return claims;
  }

  get idToken(): string {
    return this.oauthService.getIdToken();
  }

  get accessToken(): string {
    return this.oauthService.getAccessToken();
  }

  private changeTitle(lang: string): void {
    const title = this.translateService.instant('app.title');
    this.titleService.setTitle(title);
  }
}
