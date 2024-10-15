import { Component, Input, OnInit } from '@angular/core';
import { OperatorConfig } from '../../../core/models/operatorconfig';
import { ConfigService } from '../../../core/services/config.service';

@Component({
  selector: 'app-config-operator-config-list',
  templateUrl: './config-operator-config-list.component.html',
  styleUrls: ['./config-operator-config-list.component.scss'],
})
export class ConfigOperatorConfigListComponent implements OnInit {
  @Input()
  operatorConfigs: OperatorConfig[] | undefined;

  rowsPerPage = this.configService.config.rowsPerPage;

  constructor(private configService: ConfigService) {}

  ngOnInit(): void {
    return;
  }
}
