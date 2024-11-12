import { Resourcesv2Service } from './../../../core/services/resourcesv2.service';
import { AsyncPipe, JsonPipe, CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, ChangeDetectorRef, Component, inject, Input } from '@angular/core';
import { TranslateModule } from '@ngx-translate/core';
import { HighlightModule } from 'ngx-highlightjs';
import { Observable } from 'rxjs';
import { SharedModule } from '../../../shared/shared.module';

@Component({
  selector: 'app-resource-v2-details',
  standalone: true,
  imports: [CommonModule, HighlightModule, TranslateModule, JsonPipe, SharedModule, AsyncPipe],
  templateUrl: './resource-v2-details.component.html',
  styleUrl: './resource-v2-details.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ResourceV2DetailsComponent {
  @Input() set resource(resource: any) {
    this._resource = resource;
    this.fetchResource();
    this.changeDetector.detectChanges();
  }
  resourceFetchError: any;
  resource$: Observable<any> | undefined;

  private resourcesv2Service = inject(Resourcesv2Service);
  private changeDetector = inject(ChangeDetectorRef);

  get resource() {
    return this._resource;
  }

  private _resource: any;

  private fetchResource() {
    this.resource$ = this.resourcesv2Service.getResourcesById(this.resource?.metadata?.uid);
  }
}
