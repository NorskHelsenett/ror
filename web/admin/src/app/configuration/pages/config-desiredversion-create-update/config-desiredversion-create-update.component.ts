import { ChangeDetectionStrategy, ChangeDetectorRef, Component, OnDestroy, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { TranslateService } from '@ngx-translate/core';
import { MessageService } from 'primeng/api';
import { Subscription, catchError } from 'rxjs';
import { DesiredVersion } from '../../../core/models/desiredversion';
import { ConfigService } from '../../../core/services/config.service';
import { DesiredversionsService } from '../../../core/services/desiredversions.service';

@Component({
  selector: 'app-config-desiredversion-create-update',
  templateUrl: './config-desiredversion-create-update.component.html',
  styleUrls: ['./config-desiredversion-create-update.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ConfigDesiredversionCreateUpdateComponent implements OnInit, OnDestroy {
  formGroup: FormGroup;
  createError: boolean;
  updateError: boolean;
  desiredVersion: DesiredVersion;
  private rortextregex = this.configService.config.regex.forms;
  private subscriptions = new Subscription();

  constructor(
    private changeDetector: ChangeDetectorRef,
    private route: ActivatedRoute,
    private router: Router,
    private fb: FormBuilder,
    private desiredVersionsService: DesiredversionsService,
    private messageService: MessageService,
    private translateService: TranslateService,
    private configService: ConfigService,
  ) {
    this.desiredVersion = this?.router?.getCurrentNavigation()?.extras?.state as DesiredVersion;
  }

  ngOnInit(): void {
    this.setupForm();
    if (this.desiredVersion !== null) {
      this.patchForm();
    }
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }

  private setupForm(): void {
    this.formGroup = this.fb.group({
      key: ['', { validators: [Validators.required, Validators.minLength(1), Validators.pattern(this.rortextregex)] }],
      value: [null, { validators: [Validators.required, Validators.minLength(1), Validators.pattern(this.rortextregex)] }],
    });
  }

  patchForm(): void {
    this.formGroup.patchValue(this.desiredVersion);
    this.changeDetector.detectChanges();
  }

  create(): void {
    this.createError = undefined;

    const desiredVersion = this.formGroup.value as DesiredVersion;
    if (!this.formGroup.valid) {
      this.formGroup.markAllAsTouched();
      this.changeDetector.detectChanges();
      return;
    }

    this.subscriptions.add(
      this.desiredVersionsService
        .create(desiredVersion)
        .pipe(
          catchError((error: any) => {
            this.createError = error;
            this.messageService.add({
              severity: 'error',
              summary: this.translateService.instant('pages.configuration.desiredVersion.create.error'),
            });
            this.changeDetector.detectChanges();
            throw error;
          }),
        )
        .subscribe((data: DesiredVersion) => {
          this.messageService.add({
            severity: 'success',
            summary: this.translateService.instant('pages.configuration.desiredVersion.create.success', { key: data?.key }),
          });
          this.router.navigate(['../../'], { relativeTo: this.route });
          this.changeDetector.detectChanges();
        }),
    );
  }

  update(): void {
    this.updateError = undefined;
    const desiredVersion = this.formGroup.value as DesiredVersion;

    if (!this.formGroup.valid) {
      this.formGroup.markAllAsTouched();
      this.changeDetector.detectChanges();
      return;
    }

    this.subscriptions.add(
      this.desiredVersionsService
        .update(this.desiredVersion?.key, desiredVersion)
        .pipe(
          catchError((error: any) => {
            this.updateError = error;
            this.changeDetector.detectChanges();
            this.messageService.add({
              severity: 'error',
              summary: this.translateService.instant('pages.configuration.desiredVersion.update.error'),
            });
            throw error;
          }),
        )
        .subscribe((data: DesiredVersion) => {
          this.router.navigate(['../../'], { relativeTo: this.route });
          this.messageService.add({
            severity: 'success',
            summary: this.translateService.instant('pages.configuration.desiredVersion.update.success', { key: data?.key }),
          });
        }),
    );
  }

  reset(): void {
    this.formGroup.reset();
    this.changeDetector.detectChanges();
  }
}
