import { ClusterFormService } from '../../services/cluster-form.service';
import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-summary',
  templateUrl: './summary.component.html',
  styleUrls: ['./summary.component.scss'],
})
export class SummaryComponent {
  @Input() value: any;
  @Input() relativeTo: string = '.';

  constructor(private clusterFormService: ClusterFormService) {}

  getNodePoolSum(): number {
    return this.clusterFormService.getNodePoolSum();
  }
}
