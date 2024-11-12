import { Subscription, tap, catchError } from 'rxjs';
import { DatacenterService } from '../../../../../core/services/datacenter.service';
import { ChangeDetectorRef, Component, OnInit, OnDestroy } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router, ActivatedRoute } from '@angular/router';
import { Datacenter } from '../../../../../core/models/datacenter';

@Component({
  selector: 'app-admin-datacenter-create',
  templateUrl: './admin-datacenter-create.component.html',
  styleUrls: ['./admin-datacenter-create.component.scss'],
})
export class AdminDatacenterCreateComponent implements OnInit, OnDestroy {
  datacenterForm: FormGroup | undefined;
  submitted = false;
  createError: any;
  updateError: any;
  datacenterCreateModel: Datacenter | undefined;
  datacenterName: string;

  datacenter: Datacenter | undefined;
  datacenterFetchError: any;

  private subscriptions = new Subscription();

  constructor(
    private changeDetector: ChangeDetectorRef,
    private fb: FormBuilder,
    private datacenterService: DatacenterService,
    private router: Router,
    private route: ActivatedRoute,
  ) {}

  ngOnInit(): void {
    this.setupForm();

    this.subscriptions.add(
      this.route.params
        .pipe(
          tap((data: any) => {
            if (data && data?.datacenterName?.length > 0) {
              this.datacenterName = data?.datacenterName;
              this.fetchDatacenter();
            }
          }),
        )
        .subscribe(),
    );

    this.changeDetector.detectChanges();
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }

  setupForm(): void {
    this.datacenterForm = this.fb.group({
      name: [null, [Validators.required, Validators.minLength(1)]],
      provider: [null, [Validators.required, Validators.min(2)]],
      location: this.fb.group({
        region: [null, [Validators.required, Validators.min(1)]],
        country: ['Norway', [Validators.required, Validators.min(1)]],
      }),
      apiEndpoint: [null, [Validators.min(2)]],
    });
  }

  fetchDatacenter(): void {
    this.subscriptions.add(
      this.datacenterService
        .getByName(this.datacenterName)
        .pipe(
          tap((datacenter: Datacenter) => {
            this.datacenter = datacenter;
            this.fillForm();
          }),
          catchError((error) => {
            this.datacenterFetchError = error;
            this.changeDetector.detectChanges();
            throw error;
          }),
        )
        .subscribe(),
    );
  }

  fillForm(): void {
    this.datacenterForm.patchValue(this.datacenter);
    this.changeDetector.detectChanges();
  }

  createDatacenter(): void {
    this.createError = undefined;
    this.datacenterCreateModel = this.datacenterForm.value as Datacenter;
    if (!this.datacenterForm.valid) {
      this.datacenterForm.markAllAsTouched();
      this.changeDetector.detectChanges();
      return;
    }
    this.submitted = true;
    this.datacenterCreateModel = this.datacenterForm.value as Datacenter;
    this.subscriptions.add(
      this.datacenterService
        .create(this.datacenterCreateModel)
        .pipe(
          tap((datacenter: Datacenter) => {
            if (!datacenter) {
              this.createError = 'Could not create datacenter';
            } else {
              this.router.navigate(['../'], { relativeTo: this.route });
            }
            this.changeDetector.detectChanges();
          }),
          catchError((error) => {
            this.createError = error;
            this.changeDetector.detectChanges();
            throw error;
          }),
        )
        .subscribe(),
    );
    this.changeDetector.detectChanges();
  }

  updateDatacenter(): void {
    this.createError = undefined;
    this.datacenterCreateModel = this.datacenterForm.value as Datacenter;
    if (!this.datacenterForm.valid) {
      this.datacenterForm.markAllAsTouched();
      this.changeDetector.detectChanges();
      return;
    }
    this.submitted = true;
    this.datacenterCreateModel = this.datacenterForm.value as Datacenter;
    this.datacenterCreateModel.id = this.datacenter.id;
    this.subscriptions.add(
      this.datacenterService
        .update(this.datacenterCreateModel)
        .pipe(
          tap((datacenter: Datacenter) => {
            if (!datacenter) {
              this.createError = 'Could not update datacenter';
            } else {
              this.router.navigate(['../../'], { relativeTo: this.route });
            }
            this.changeDetector.detectChanges();
          }),
          catchError((error) => {
            this.createError = error;
            this.changeDetector.detectChanges();
            throw error;
          }),
        )
        .subscribe(),
    );
    this.changeDetector.detectChanges();
  }

  reset(): void {
    this.createError = undefined;
    this.datacenterCreateModel = null;
    this.submitted = false;
    this.datacenterForm.reset();
    this.changeDetector.detectChanges();
  }
}
