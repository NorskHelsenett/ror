import { Component, Input, OnInit } from '@angular/core';
import { ResourcesService } from '../../../core/services/resources.service';
import { v4 } from 'uuid';
import { catchError, map, mergeMap, Observable, tap } from 'rxjs';
import { Resource, ResourceQuery, ResourceRoute, ResourceSet } from '../../../core/models/resources-v2';

@Component({
  selector: 'app-cluster-notifications',
  templateUrl: './cluster-notifications.component.html',
  styleUrl: './cluster-notifications.component.scss',
})
export class ClusterNotificationsComponent implements OnInit {
  @Input({ required: true }) cluster!: any;
  resources$: Observable<Resource[]>;
  routeError: any;
  showPopup: boolean = false;

  resource: Resource = undefined;
  query: ResourceQuery;

  edit: boolean = false;
  readonly apiVersion: string = 'general.ror.internal/v1alpha1';
  readonly kind: string = 'Route';
  readonly scope: string = 'cluster';

  constructor(private resourcesService: ResourcesService) {}

  ngOnInit(): void {
    this.showPopup = false;
    this.query = {
      versionkind: {
        Kind: this.kind,
        Group: '',
        Version: this.apiVersion,
      },
      ownerrefs: [
        {
          scope: this.scope,
          subject: this.cluster?.clusterId,
        },
      ],
    };
    this.resources$ = this.getResources();
  }

  getResources(): Observable<Resource[]> {
    return this.resourcesService.getResources<ResourceSet>(this.query).pipe(
      tap((resourceSet) => (this.resource = resourceSet?.resources[0])),
      map((resourceSet) => resourceSet?.resources ?? []),
      catchError((err, caught) => {
        this.routeError = err;
        return caught;
      }),
    );
  }

  deleteRoute(): void {
    this.resources$ = this.resourcesService.deleteResource(this.resource.metadata.uid).pipe(
      mergeMap(() => {
        return this.getResources();
      }),
      catchError((err, caught) => {
        this.routeError = err;
        return caught;
      }),
    );
    this.showPopup = !this.showPopup;
  }

  getReceivers(route: ResourceRoute): any[] {
    let keys: string[] = Object.keys(route.spec.receivers);
    let receivers = keys.map((key) => {
      return {
        type: key,
        spec: route.spec.receivers,
      };
    });
    return receivers;
  }

  getSlackReceivers(route: ResourceRoute): any[] {
    return route?.spec?.receivers?.slack;
  }

  getKeys(route: ResourceRoute): string[] {
    return Object.keys(route?.spec?.receivers);
  }

  toggleEdit(): void {
    this.edit = !this.edit;
  }

  handleCreate(): void {
    if (!this.resource) {
      this.resource = {
        kind: this.kind,
        apiVersion: this.apiVersion,
        rormeta: {
          ownerref: {
            scope: this.scope,
            subject: this.cluster?.clusterId,
          },
          action: 'Add',
        },
        metadata: {
          uid: v4(),
        },
      };
    }
    this.toggleEdit();
  }

  handleUpdate(): void {
    this.toggleEdit();
  }

  handleCancel(): void {
    this.toggleEdit();
  }

  update(resource: Resource): void {
    this.resource = resource;
    this.toggleEdit();
  }

  popup(resource?: Resource): void {
    this.resource = resource;
    this.showPopup = !this.showPopup;
  }
}
