import { ChangeDetectionStrategy, ChangeDetectorRef, Component, Input, OnInit } from '@angular/core';
import { ConfigService } from '../../../core/services/config.service';

@Component({
  selector: 'app-cluster-ingress-list',
  templateUrl: './cluster-ingress-list.component.html',
  styleUrls: ['./cluster-ingress-list.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ClusterIngressListComponent implements OnInit {
  @Input() cluster: any = undefined;

  ingresses: any[] = [];
  rows = this.configService.config.rows;
  rowsPerPage = this.configService.config.rowsPerPage;

  constructor(
    private changeDetector: ChangeDetectorRef,
    private configService: ConfigService,
  ) {}

  ngOnInit(): void {
    if (!this.cluster) {
      this.changeDetector.detectChanges();
      return;
    }

    this.ingresses = this.cluster?.ingresses;
    this.changeDetector.detectChanges();
  }
}
