import { UnitformaterService } from '../../core/services/unitformater.service';
import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'formatBytes',
})
export class FormatBytesPipe implements PipeTransform {
  constructor(private unitformaterService: UnitformaterService) {}

  transform(bytes: number, decimals: number, binaryUnits: boolean): string {
    return this.unitformaterService.formatBytes(bytes, decimals, binaryUnits);
  }
}
