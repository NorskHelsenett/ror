import { Component, Input, OnInit, QueryList, ViewChildren } from '@angular/core';
import { Table } from 'primeng/table';
import { ConfigService } from '../../../core/services/config.service';

@Component({
  selector: 'app-policy-namespace',
  templateUrl: './policy-namespace.component.html',
  styleUrls: ['./policy-namespace.component.scss'],
})
export class PolicyNamespaceComponent implements OnInit {
  @ViewChildren('namespaceTable')
  tables: QueryList<Table>;

  @Input() namespace: any;
  rowsPerPage = this.configService.config.rowsPerPage;

  resultFilter: string[] = ['failed', 'error', 'passed', 'warning', 'skipped'];
  resultFilterValue: string[];

  constructor(private configService: ConfigService) {}

  ngOnInit(): void {
    this.resultFilterValue = ['failed'];
    this.triggerFilter();
  }

  triggerFilter(): void {
    if (this.tables && this.tables?.length > 0) {
      this.tables.forEach((table: Table) => {
        table.filter(['fail', 'error'], 'result', 'in');
      });
    }
  }
}
