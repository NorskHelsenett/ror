import { Subscription } from 'rxjs';
import { PriceService } from '../../../core/services/price.service';
import { Component, Input, OnInit, OnDestroy, ChangeDetectionStrategy, ChangeDetectorRef } from '@angular/core';
import { Price } from '../../../core/models/price';
import { ConfigService } from '../../../core/services/config.service';

@Component({
  selector: 'app-cluster-node-pools',
  templateUrl: './cluster-node-pools.component.html',
  styleUrls: ['./cluster-node-pools.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ClusterNodePoolsComponent implements OnInit, OnDestroy {
  @Input() cluster: any = undefined;

  rows = this.configService.config.rows;
  rowsPerPage = this.configService.config.rowsPerPage;

  private prices: Price[];
  private subscriptions = new Subscription();

  constructor(
    private changeDetector: ChangeDetectorRef,
    private priceService: PriceService,
    private configService: ConfigService,
  ) {}

  ngOnInit(): void {
    this.subscriptions.add(
      this.priceService.getAll().subscribe((prices: Price[]) => {
        this.prices = prices;
        this.changeDetector.detectChanges();
      }),
    );
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }

  findMachineClasses(): void {
    for (let i = 0; i < this.cluster?.topology?.nodePools?.length; i++) {
      let nodePool = this.cluster?.topology?.nodePools[i];
      nodePool['machineClass'] = 'Unknown';
      if (!nodePool) {
        continue;
      }

      if (nodePool?.nodes?.length === 0) {
        continue;
      }

      const node = nodePool?.nodes[0];
      const price = this.prices.find((p: any) => {
        if (
          p?.memoryBytes === node?.metrics?.memory &&
          p?.provider === this.cluster?.workspace?.datacenter?.provider &&
          p?.cpu === node?.metrics?.cpu
        ) {
          return p;
        }
      });

      if (price) {
        nodePool['machineClass'] = price.machineClass;
      }
    }
  }
}
