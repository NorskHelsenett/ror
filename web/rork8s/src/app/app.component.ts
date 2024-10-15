import { ChangeDetectionStrategy, ChangeDetectorRef, Component, HostListener, OnDestroy, OnInit } from '@angular/core';
import { NgOptimizedImage } from '@angular/common';
import { RouterOutlet } from '@angular/router';
import { environment } from '../environments/environment';
import { CommonModule } from '@angular/common';
import { Config } from './core/models/config';
import { ConfigService } from './core/services/config.service';
import { ThemeService } from './core/services/theme.service';
import { Subscription } from 'rxjs';
import { TranslateModule, TranslateService } from '@ngx-translate/core';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, CommonModule, NgOptimizedImage, TranslateModule],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AppComponent implements OnInit, OnDestroy {
  date: Date = new Date();
  environment = environment;
  config: Config = this.configService.config;

  isDark: boolean = this.themeService.isDark.value;
  lang = 'en';
  hasScrolled = false;

  private subscriptions = new Subscription();
  private isLocalStorageAvailable = typeof localStorage !== 'undefined';

  constructor(
    private changeDetector: ChangeDetectorRef,
    private configService: ConfigService,
    private themeService: ThemeService,
    private translateService: TranslateService,
  ) {}

  @HostListener('window:scroll', ['$event'])
  onWindowsScroll(event: Event) {
    const scrollTop = document.documentElement.scrollTop || document.body.scrollTop;
    if (scrollTop >= 20) {
      this.hasScrolled = true;
      this.changeDetector.detectChanges();
    } else {
      this.hasScrolled = false;
      this.changeDetector.detectChanges();
    }
  }

  ngOnInit(): void {
    this.subscriptions.add(
      this.themeService.isDark.subscribe((value) => {
        this.isDark = value;
        this.changeDetector.detectChanges();
      }),
    );

    this.translateService.addLangs(['en', 'no']);
    this.translateService.setDefaultLang('no');

    let lang = 'en';
    if (this.isLocalStorageAvailable) {
      const userLang = localStorage.getItem('language');
      if (userLang && userLang.length > 0) {
        lang = userLang;
      } else {
        const browserLang = this.translateService.getBrowserLang();
        lang = browserLang?.match(/en|no/) ? browserLang : 'en';
      }
    }

    this.translateService.use(lang);
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }

  toggleDark(): void {
    this.isDark = !this.isDark;
    this.themeService.setDark(this.isDark);
  }

  useLanguage(lang: string): void {
    if (!lang || (lang !== 'en' && lang !== 'no')) {
      return;
    }

    this.translateService.use(lang);
    this.lang = lang;
    if (this.isLocalStorageAvailable) {
      localStorage.setItem('language', lang);
    }
    this.changeDetector.detectChanges();
  }
}
