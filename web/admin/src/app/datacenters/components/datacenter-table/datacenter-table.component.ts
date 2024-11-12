import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-datacenter-table',
  templateUrl: './datacenter-table.component.html',
  styleUrls: ['./datacenter-table.component.scss'],
})
export class DatacenterTableComponent implements OnInit {
  @Input() datacenters: any[] = [];

  constructor() {}

  ngOnInit(): void {
    return;
  }
}
