import { Injectable } from '@angular/core';

@Injectable()
export class UtilsService {
  constructor() {}

  isEqual(x: any, y: any): boolean {
    const ok = Object.keys,
      tx = typeof x,
      ty = typeof y;
    return x && y && tx === 'object' && tx === ty ? ok(x).length === ok(y).length && ok(x).every((key) => this.isEqual(x[key], y[key])) : x === y;
  }
}
