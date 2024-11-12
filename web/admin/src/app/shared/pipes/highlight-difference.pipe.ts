import { Pipe, PipeTransform, SecurityContext } from '@angular/core';
import { DomSanitizer } from '@angular/platform-browser';

@Pipe({
  name: 'highlightDifference',
})
export class HighlightDifferencePipe implements PipeTransform {
  constructor(private sanitizer: DomSanitizer) {}
  transform(changedObject: any, originalObject: any, color: number): unknown {
    let colorString: string;
    switch (color) {
      case 0:
        colorString = 'bg-green-400';
        break;
      case 1:
        colorString = 'bg-red-400';
        break;
      default:
        colorString = 'bg-yellow-400';
    }
    if (!changedObject) {
      return this.sanitizer.sanitize(SecurityContext.HTML, '<pre><mark class="' + colorString + '">null</mark></pre>');
    }
    if (!originalObject) {
      return this.sanitizer.sanitize(
        SecurityContext.HTML,
        '<pre><mark class="' + colorString + '">' + JSON.stringify(changedObject, null, 2) + '</mark></pre>',
      );
    }
    try {
      let highlighted = this.highlight(changedObject, originalObject, colorString);
      let result = '<pre>' + JSON.stringify(highlighted, null, 2) + '</pre>';
      const regex = new RegExp('"(?=<mark class=\\\\"' + colorString + '\\\\">[0-9]+<)|(?<=>[0-9]+</mark>)"|\\\\(?=")', 'gi');
      result = result.replace(regex, '');
      return this.sanitizer.sanitize(SecurityContext.HTML, result);
    } catch {
      return this.sanitizer.sanitize(SecurityContext.HTML, JSON.stringify(changedObject, null, '\t'));
    }
  }

  highlight(changedObject: any, originalObject: any, colorString: string): any {
    let newValue: any;
    let oldValue: any;
    Object.keys(changedObject).forEach((key) => {
      newValue = changedObject[key];
      oldValue = originalObject[key];
      if (newValue !== oldValue) {
        if (typeof newValue === 'string' || newValue instanceof String) {
          newValue = newValue as string;
          changedObject[key] = '<mark class="' + colorString + '">' + newValue + '</mark>';
        } else if (typeof newValue === 'number' || newValue instanceof Number) {
          newValue = newValue as number;
          changedObject[key] = '<mark class="' + colorString + '">' + newValue + '</mark>';
        } else {
          changedObject[key] = this.highlight(changedObject[key], originalObject[key], colorString);
        }
      }
    });
    return changedObject;
  }
}
