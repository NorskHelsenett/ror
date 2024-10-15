import { Pipe, PipeTransform } from '@angular/core';

import dayjs from 'dayjs';
import customParseFormat from 'dayjs/plugin/customParseFormat';
import 'dayjs/locale/nb';
import 'dayjs/locale/en';

@Pipe({
  name: 'time',
})
export class TimePipe implements PipeTransform {
  constructor() {
    dayjs.extend(customParseFormat);
  }

  transform(date: Date, format: string = '', locale: string = 'nb'): unknown {
    let d = dayjs(date);
    if (!d.isValid()) {
      return '';
    }

    const year = d.year();
    if (year === 0 || year === 1) {
      return '';
    }
    return d.locale(locale).format(format);
  }
}
