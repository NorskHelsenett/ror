import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class UnitformaterService {
  formatBytes(bytes: number, decimals: number, binaryUnits: boolean): string {
    if (bytes == 0) {
      return '0 Bytes';
    }
    const unitMultiple = binaryUnits ? 1024 : 1000;
    const unitNames =
      unitMultiple === 1024 // 1000 bytes in 1 Kilobyte (KB) or 1024 bytes for the binary version (KiB)
        ? ['Bytes', 'KiB', 'MiB', 'GiB', 'TiB', 'PiB', 'EiB', 'ZiB', 'YiB']
        : ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];
    const unitChanges = Math.floor(Math.log(bytes) / Math.log(unitMultiple));
    return parseFloat((bytes / Math.pow(unitMultiple, unitChanges)).toFixed(decimals || 0)) + ' ' + unitNames[unitChanges];
  }

  formatBytesInNumber(bytes: number, binaryUnits: boolean): number {
    if (bytes == 0) {
      return 0;
    }
    const unitMultiple = binaryUnits ? 1024 : 1000;
    const unitChanges = Math.floor(Math.log(bytes) / Math.log(unitMultiple));
    return bytes / Math.pow(unitMultiple, unitChanges);
  }
}
