import { DOCUMENT } from '@angular/common';
import { Inject, Injectable } from '@angular/core';
import { HighlightLoader } from 'ngx-highlightjs';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class ThemeService {
  isDark = new BehaviorSubject<boolean>(true);

  constructor(
    @Inject(DOCUMENT) private document: Document,
    private hljsLoader: HighlightLoader,
  ) {
    const isDark = localStorage.getItem('isDark') == 'true';
    if (isDark === true) {
      this.isDark.next(true);
      this.switchTheme('dark');
    } else {
      this.isDark.next(false);
      this.switchTheme('light');
    }
  }

  setDark(setDark: boolean): void {
    this.isDark.next(setDark);
    localStorage['isDark'] = setDark;
    if (this.isDark?.getValue() === true) {
      this.switchTheme('dark');
    } else {
      this.switchTheme('light');
    }
  }

  switchTheme(theme: string) {
    let themeLink = this.document.getElementById('app-theme') as HTMLLinkElement;

    if (themeLink) {
      themeLink.href = `${theme}.css`;
    }

    this.hljsLoader.setTheme(theme === 'dark' ? 'assets/styles/highlight/tokyo-night-dark.css' : 'assets/styles/highlight/tokyo-night-light.css');
  }
}
