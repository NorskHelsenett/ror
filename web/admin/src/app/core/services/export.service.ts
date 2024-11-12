import { Injectable } from '@angular/core';
import { TranslateService } from '@ngx-translate/core';
import * as FileSaver from 'file-saver';
import { MessageService } from 'primeng/api';
import * as ExcelJS from 'exceljs';
import { saveAs } from 'file-saver';

@Injectable({
  providedIn: 'root',
})
export class ExportService {
  constructor(
    private messageService: MessageService,
    private translateService: TranslateService,
  ) {}

  exportToCsv(data: any[], filename: string): void {
    if (!data || !data[0] || !filename || filename === '') {
      this.messageService.add({
        severity: 'error',
        summary: this.translateService.instant('shared.export.errorSummary'),
        detail: this.translateService.instant('shared.export.errorDetail'),
      });
      return;
    }
    const replacer = (key: any, value: any) => (value === null ? '' : value);
    const header = Object.keys(data[0]);
    let csv = data.map((row) => header.map((fieldName) => JSON.stringify(row[fieldName], replacer)).join(','));
    csv.unshift(header.join(','));
    let csvArray = csv.join('\r\n');

    const blob = new Blob([csvArray], { type: 'text/csv' });
    FileSaver.saveAs(blob, filename);
  }

  exportAsExcelFile(data: any[], filename: string): void {
    if (!data || !data[0] || !filename || filename === '') {
      this.messageService.add({
        severity: 'error',
        summary: this.translateService.instant('shared.export.errorSummary'),
        detail: this.translateService.instant('shared.export.errorDetail'),
      });
      return;
    }

    const workbook = new ExcelJS.Workbook();
    const worksheet = workbook.addWorksheet('Sheet 1');

    const headers = Object.keys(data[0]);
    worksheet.addRow(headers);

    data.forEach((item) => {
      const row = [];
      headers.forEach((header) => {
        row.push(item[header]);
      });
      worksheet.addRow(row);
    });

    workbook.xlsx.writeBuffer().then((buffer) => {
      const blob = new Blob([buffer], { type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' });
      saveAs(blob, `${filename}.xlsx`);
    });
  }
}
