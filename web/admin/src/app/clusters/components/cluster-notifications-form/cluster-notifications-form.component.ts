import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { FormArray, FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ResourcesService } from '../../../core/services/resources.service';
import { catchError, Observable, tap } from 'rxjs';
import { Resource, ResourceRoute, ResourceSet } from '../../../core/models/resources-v2';

@Component({
  selector: 'app-cluster-notifications-form',
  templateUrl: './cluster-notifications-form.component.html',
  styleUrl: './cluster-notifications-form.component.scss',
})
export class ClusterNotificationsFormComponent implements OnInit {
  @Input() resource: Resource;
  @Output() updateOk = new EventEmitter<boolean>();
  @Output() cancelUpdate = new EventEmitter<boolean>();

  result$: Observable<any>;

  routeForm: FormGroup;
  dropdownOptions: any[] = [
    {
      apiVersion: 'general.ror.internal/v1alpha1',
      kind: 'VulnerabilityEvent',
      label: 'Vulnerability',
    },
  ];

  readonly apiVersion: string = 'general.ror.internal/v1alpha1';
  readonly kind: string = 'Route';
  readonly scope: string = 'cluster';

  create: boolean = false;

  constructor(
    private fb: FormBuilder,
    private resourcesService: ResourcesService,
  ) {}

  ngOnInit(): void {
    this.setupForm();
    if (this.resource?.route != null) {
      this.fillForm();
    } else {
      this.create = true;
      this.addSlackReceiver('');
      this.routeForm.setValue({
        messageType: this.dropdownOptions[0],
        receivers: {
          slack: [
            {
              channelId: '',
            },
          ],
        },
      });
      this.resource.route = {
        spec: {
          messageType: this.dropdownOptions[0],
          receivers: {
            slack: [],
          },
        },
      };
    }
  }

  setupForm(): void {
    this.routeForm = this.fb.group({
      messageType: [
        {
          apiVersion: ['', Validators.required],
          kind: ['', Validators.required],
        },
        Validators.required,
      ],
      receivers: this.fb.group({
        slack: this.fb.array([]),
      }),
    });
  }

  get slackReceivers(): FormArray {
    return this.routeForm.get('receivers').get('slack') as FormArray;
  }

  addSlackReceiver(channelId: string): void {
    this.slackReceivers.push(this.fb.group({ channelId: [channelId, Validators.required] }));
  }

  removeSlackReceiver(index: number): void {
    this.slackReceivers.removeAt(index);
  }

  fillForm(): void {
    this.resource?.route?.spec?.receivers?.slack?.forEach((r) => {
      this.addSlackReceiver(r.channelId);
    });
    this.routeForm.setValue({
      messageType: this.resource.route.spec.messageType,
      receivers: {
        slack: this.resource.route.spec.receivers.slack,
      },
    });
  }

  onSubmit(): void {
    const route: ResourceRoute = {
      spec: this.routeForm.value,
    };
    this.resource.route = route;
    this.result$ = this.resourcesService.createResource(this.toResourceSet()).pipe(
      tap(() => {
        this.updateOk.emit(true);
      }),
      catchError((error) => {
        return error;
      }),
    );
  }

  toResourceSet(): ResourceSet {
    const resourceSet: ResourceSet = {
      resources: [this.resource],
    };
    return resourceSet;
  }

  cancel(): void {
    this.cancelUpdate.emit(true);
  }
}
