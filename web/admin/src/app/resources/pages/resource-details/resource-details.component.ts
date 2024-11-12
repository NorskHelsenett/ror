import { CommonModule } from '@angular/common';
import { Observable, Subscription, catchError } from 'rxjs';
import { ResourcesService } from './../../../core/services/resources.service';
import { Component, OnDestroy, OnInit, ChangeDetectionStrategy, ChangeDetectorRef } from '@angular/core';
import { ActivatedRoute, Router, RouterModule } from '@angular/router';
import { TranslateModule } from '@ngx-translate/core';
import { SharedModule } from '../../../shared/shared.module';
import { ErrorComponent } from '../../../shared/components/error/error.component';

@Component({
  selector: 'app-resource-details',
  standalone: true,
  imports: [CommonModule, RouterModule, TranslateModule, SharedModule, ErrorComponent],
  templateUrl: './resource-details.component.html',
  styleUrl: './resource-details.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ResourceDetailsComponent implements OnInit, OnDestroy {
  uid: string = undefined;
  kind: string = undefined;
  apiVersion: string = undefined;
  scope: string = undefined;
  subject: string = undefined;
  resource$: Observable<any> | undefined;
  resourceFetchError: any;

  private subscriptions = new Subscription();

  constructor(
    private changeDetector: ChangeDetectorRef,
    private route: ActivatedRoute,
    private router: Router,
    private resourcesService: ResourcesService,
  ) {}

  ngOnInit(): void {
    this.subscriptions.add(
      this.route.params.subscribe((data: any) => {
        this.uid = data?.uid;
        this.kind = data?.kind;
        this.apiVersion = data?.apiVersion;
        this.scope = data?.scope;
        this.subject = data?.subject;

        this.fetchResource();
        this.changeDetector.detectChanges();
      }),
    );
  }

  ngOnDestroy(): void {
    return;
  }

  fetchResource(): void {
    this.resourceFetchError = undefined;
    this.resource$ = this.resourcesService.getResource(this.uid, this.scope, this.subject, this.kind, this.apiVersion).pipe(
      catchError((error) => {
        this.resourceFetchError = error;
        this.changeDetector.detectChanges();
        throw error;
      }),
    );
  }
}
