import { ChangeDetectionStrategy, ChangeDetectorRef, Component, OnDestroy, OnInit } from '@angular/core';
import { FormBuilder } from '@angular/forms';
import { ActivatedRoute } from '@angular/router';
import { LangChangeEvent, TranslateService } from '@ngx-translate/core';
import { catchError, Subscription, tap } from 'rxjs';
import { OperatorConfig } from '../../../core/models/operatorconfig';
import { OperatorConfigsService } from '../../../core/services/operator-configs.service';

@Component({
  selector: 'app-config-operatorconfig-create-update',
  templateUrl: './config-operatorconfig-create-update.component.html',
  styleUrls: ['./config-operatorconfig-create-update.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ConfigOperatorconfigCreateUpdateComponent implements OnInit, OnDestroy {
  id: string;
  operatorConfig: OperatorConfig | undefined;
  operatorConfigFetchError: any;

  constructor(
    private changeDetector: ChangeDetectorRef,
    private fb: FormBuilder,
    private route: ActivatedRoute,
    private translateService: TranslateService,
    private operatorConfigService: OperatorConfigsService,
  ) {}

  private subscriptions = new Subscription();

  ngOnInit(): void {
    this.setupForm();

    this.subscriptions.add(
      this.route.params.subscribe((param) => {
        this.id = param?.['id'];
        if (this.id !== '' && this.id !== null && this.id !== undefined) {
          this.fetch();
          this.fillForm();
          this.changeDetector.detectChanges();
        }
      }),
    );

    this.subscriptions.add(
      this.translateService.onLangChange
        .pipe(
          tap((event: LangChangeEvent) => {
            event.lang;
          }),
        )
        .subscribe(),
    );

    this.changeDetector.detectChanges();
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }

  fetch(): void {
    this.operatorConfigFetchError = undefined;
    this.subscriptions.add(
      this.operatorConfigService
        .getById(this.id)
        .pipe(
          catchError((error: any) => {
            this.operatorConfigFetchError = error;
            this.changeDetector.detectChanges();
            throw error;
          }),
          tap((operatorConfig: OperatorConfig) => {
            this.operatorConfig = operatorConfig;
            this.fillForm();
            this.changeDetector.detectChanges();
          }),
        )
        .subscribe(),
    );
  }

  private setupForm(): void {}

  private fillForm(): void {}
}
