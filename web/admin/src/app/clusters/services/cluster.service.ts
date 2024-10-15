import { Injectable } from '@angular/core';

@Injectable()
export class ClusterService {
  fillTags(serviceTags: any[]): string[] {
    let tags: string[] = [];
    if (serviceTags) {
      const keys = Object.keys(serviceTags);
      keys.forEach((key: string) => {
        tags.push(key);
      });
    }
    return tags;
  }
}
