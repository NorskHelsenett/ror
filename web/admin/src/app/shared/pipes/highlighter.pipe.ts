import { Pipe, PipeTransform, SecurityContext } from '@angular/core';
import { DomSanitizer } from '@angular/platform-browser';

@Pipe({
  name: 'highlighter',
})
export class HighlighterPipe implements PipeTransform {
  constructor(private sanitizer: DomSanitizer) {}
  transform(value: string, args: string[]): unknown {
    if (!args || args.length === 0) {
      return this.sanitizer.sanitize(SecurityContext.HTML, value);
    }
    if (args.filter((item) => item !== undefined).map((item) => item.length > 0).length === 0) {
      return this.sanitizer.sanitize(SecurityContext.HTML, value);
    }
    const regex = new RegExp(args.join('|'), 'gi');
    value = value.replace(regex, '<mark class="bg-primary">$&</mark>');
    value = this.sanitizer.sanitize(SecurityContext.HTML, value);
    return value;
  }
}
