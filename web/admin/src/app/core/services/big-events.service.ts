import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class BigEventsService {
  private now = new Date();

  isDesember(): boolean {
    if (this.now.getMonth() === 11) {
      return true;
    }
    return false;
  }

  isRORBirthday(): boolean {
    let date = this.now.getDate();
    if (date === 17 && this.now.getMonth() === 2) {
      return true;
    }
    return false;
  }
}
