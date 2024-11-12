import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class ThemeService {
  isDark = new BehaviorSubject<boolean>(true);

  private isLocalStorageAvailable = typeof localStorage !== 'undefined';

  constructor() {
    if (this.isLocalStorageAvailable) {
      const localIsDark = window.localStorage.getItem('isDark');
      if (localIsDark === null) {
        this.isDark.next(true);
        return;
      }

      const isDark = localIsDark == 'true';
      if (isDark === true) {
        this.isDark.next(true);
      } else {
        this.isDark.next(false);
      }
    }
  }

  setDark(setDark: boolean): void {
    this.isDark.next(setDark);
    if (this.isLocalStorageAvailable) {
      localStorage['isDark'] = setDark;
    }
  }
}
