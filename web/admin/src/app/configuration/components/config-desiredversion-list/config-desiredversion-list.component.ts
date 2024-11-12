import { ChangeDetectionStrategy, ChangeDetectorRef, Component, Input, OnDestroy } from '@angular/core';
import { TranslateService } from '@ngx-translate/core';
import { ConfirmationService, MessageService } from 'primeng/api';
import { catchError, Subscription, tap } from 'rxjs';
import { DesiredVersion } from '../../../core/models/desiredversion';
import { ConfigService } from '../../../core/services/config.service';
import { DesiredversionsService } from '../../../core/services/desiredversions.service';

@Component({
  selector: 'app-config-desiredversion-list',
  templateUrl: './config-desiredversion-list.component.html',
  styleUrls: ['./config-desiredversion-list.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ConfigDesiredversionListComponent implements OnDestroy {
  @Input()
  desiredVersion: DesiredVersion[] | undefined;
  desiredVersionsFetchError: any;

  rowsPerPage = this.configService.config.rowsPerPage;

  private subscriptions = new Subscription();

  constructor(
    private changeDetector: ChangeDetectorRef,
    private desiredVersionsService: DesiredversionsService,
    private confirmationService: ConfirmationService,
    private messageService: MessageService,
    private translateService: TranslateService,
    private configService: ConfigService,
  ) {}

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }

  isStringType(obj: any): boolean {
    return typeof obj === 'string';
  }

  fetch(): void {
    this.desiredVersionsFetchError = undefined;
    this.subscriptions.add(
      this.desiredVersionsService
        .getAll()
        .pipe(
          tap((data: any) => {
            this.desiredVersion = data;
            this.changeDetector.detectChanges();
          }),
          catchError((error: any) => {
            this.desiredVersionsFetchError = error;
            this.changeDetector.detectChanges();
            throw error;
          }),
        )
        .subscribe(),
    );
    this.changeDetector.detectChanges();
  }

  delete(desiredVersion: DesiredVersion): void {
    this.confirmationService.confirm({
      header: this.translateService.instant('pages.configuration.desiredVersion.delete.title'),
      message: this.translateService.instant('pages.configuration.desiredVersion.delete.details', { key: desiredVersion?.key }),
      accept: () => {
        this.subscriptions.add(
          this.desiredVersionsService
            .delete(desiredVersion?.key)
            .pipe(
              catchError((error: any) => {
                this.messageService.add({
                  severity: 'error',
                  summary: this.translateService.instant('pages.configuration.desiredVersion.delete.errortitle'),
                  detail: this.translateService.instant('pages.configuration.desiredVersion.delete.errordetails'),
                });
                this.changeDetector.detectChanges();
                throw error;
              }),
            )
            .subscribe(() => {
              this.fetch();
              this.messageService.add({
                severity: 'success',
                summary: this.translateService.instant('pages.configuration.desiredVersion.delete.success'),
              });
              this.changeDetector.detectChanges();
            }),
        );
      },
    });
    this.changeDetector.detectChanges();
  }
}
