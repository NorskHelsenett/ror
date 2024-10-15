import { Component, Input, OnInit } from '@angular/core';
import { ConfigService } from '../../../core/services/config.service';

@Component({
  selector: 'app-cluster-acl',
  templateUrl: './cluster-acl.component.html',
  styleUrls: ['./cluster-acl.component.scss'],
})
export class ClusterACLComponent implements OnInit {
  @Input() cluster: any | undefined;

  rows = this.configService.config.rows;
  rowsPerPage = this.configService.config.rowsPerPage;
  accessGroups: any[];

  constructor(private configService: ConfigService) {}

  ngOnInit(): void {
    this.accessGroups = this.cluster?.acl?.accessGroups?.map((x: any) => {
      return { groupName: x };
    });
  }
}
