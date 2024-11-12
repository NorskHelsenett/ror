import { Component, Input, OnInit } from '@angular/core';
import { HealthCheckService } from '../../../core/models/healthcheckservice';

@Component({
  selector: 'app-health-table',
  templateUrl: './health-table.component.html',
  styleUrls: ['./health-table.component.scss'],
})
export class HealthTableComponent implements OnInit {
  @Input() health: any[] = [];

  constructor() {}

  ngOnInit(): void {
    return;
  }

  getServices(): Array<HealthCheckService> {
    return Object.values(this.health['services']);
  }
}
