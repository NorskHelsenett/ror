import { Component, Input, OnInit } from '@angular/core';
import { ConfigService } from '../../../core/services/config.service';
import { Task } from '../../../core/models/task';
@Component({
  selector: 'app-config-task-list',
  templateUrl: './config-task-list.component.html',
  styleUrls: ['./config-task-list.component.scss'],
})
export class ConfigTaskListComponent implements OnInit {
  @Input()
  tasks: Task[] | undefined;

  rowsPerPage = this.configService.config.rowsPerPage;

  constructor(private configService: ConfigService) {}

  ngOnInit(): void {
    return;
  }
}
