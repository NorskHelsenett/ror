import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-cluster-metrics',
  templateUrl: './cluster-metrics.component.html',
  styleUrls: ['./cluster-metrics.component.scss'],
})
export class ClusterMetricsComponent implements OnInit {
  @Input() metrics: any = undefined;

  constructor() {}

  ngOnInit(): void {
    return;
  }
}
