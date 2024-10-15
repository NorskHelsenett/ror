import { ChangeDetectionStrategy, ChangeDetectorRef, Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { ClusterService } from '../../services';

@Component({
  selector: 'app-cluster-metadata-page',
  templateUrl: './cluster-metadata-page.component.html',
  styleUrl: './cluster-metadata-page.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ClusterMetadataPageComponent implements OnInit {
  @Input() cluster: any;
  @Output() refreshRequested = new EventEmitter<void>();
  edit = false;

  tags: string[] = [];

  constructor(
    private changeDetector: ChangeDetectorRef,
    private clusterService: ClusterService,
  ) {}

  ngOnInit(): void {
    this.tags = this.clusterService.fillTags(this.cluster.metadata?.serviceTags || this.cluster.metadata?.project?.projectMetadata?.serviceTags);
  }

  toggleEdit() {
    this.edit = !this.edit;
  }

  onUpdateOk(event: boolean): void {
    if (event) {
      this.toggleEdit();
      this.refreshRequested.emit();
    } else {
      console.log('Update failed');
    }
  }
}
