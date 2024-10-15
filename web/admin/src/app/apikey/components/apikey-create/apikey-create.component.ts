import { ChangeDetectorRef, Component, EventEmitter, OnDestroy, OnInit, Output } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TranslateService } from '@ngx-translate/core';
import { ClipboardService } from 'ngx-clipboard';
import { MessageService } from 'primeng/api';
import { Subscription, catchError } from 'rxjs';
import { ApiKey } from '../../../core/models/apikey';
import { ApikeysService } from '../../../core/services/apikeys.service';
import { ConfigService } from '../../../core/services/config.service';

@Component({
  selector: 'app-apikey-create',
  templateUrl: './apikey-create.component.html',
  styleUrls: ['./apikey-create.component.scss'],
})
export class ApikeyCreateComponent implements OnInit, OnDestroy {
  @Output() cancelRequested = new EventEmitter<void>();
  @Output() created = new EventEmitter<void>();

  apikeyForm: FormGroup;
  apikeyCreateError: any;
  minDate: Date;
  maxDate: Date;
  apikeyText: string;
  rorApiUrl: string = this.configService.config.rorApi;

  availableTypes: any[];

  private subscriptions = new Subscription();
  private rortextregex = this.configService.config.regex.forms;

  constructor(
    private changeDetector: ChangeDetectorRef,
    private fb: FormBuilder,
    private apikeysService: ApikeysService,
    private clipboardService: ClipboardService,
    private messageService: MessageService,
    private translateService: TranslateService,
    private configService: ConfigService,
  ) {}

  ngOnInit(): void {
    const today = new Date();
    const year = today.getFullYear();
    const month = today.getMonth();

    this.minDate = new Date();
    this.maxDate = new Date(year + 2, month, 0, 0, 0, 0);

    this.setupForm();
    this.changeDetector.detectChanges();

    this.setAvailableTypes();

    this.subscriptions.add(
      this.translateService.onLangChange.subscribe((event: any) => {
        this.setAvailableTypes();
        this.changeDetector.detectChanges();
      }),
    );
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }

  setupForm(): void {
    this.apikeyForm = this.fb.group({
      displayName: [null, { validators: [Validators.required, Validators.minLength(5), Validators.pattern(this.rortextregex)] }],
      identifier: [null, { validators: [Validators.required, Validators.minLength(5), Validators.pattern(this.rortextregex)] }],
      expires: [undefined, { validators: [] }],
      type: [undefined, { validators: [Validators.required] }],
    });
  }

  create(): void {
    this.apikeyCreateError = undefined;
    this.apikeyText = undefined;
    let apikey: ApiKey = this.apikeyForm.value as ApiKey;

    apikey.type = 'Service';

    if (apikey.expires && apikey?.expires !== new Date(0)) {
      apikey?.expires.setMinutes(0, 0, 0);
    }

    this.subscriptions.add(
      this.apikeysService
        .create(apikey)
        .pipe(
          catchError((error: any) => {
            this.apikeyCreateError = error;
            this.changeDetector.detectChanges();
            throw error;
          }),
        )
        .subscribe((result: string) => {
          this.apikeyText = result;
          this.created.emit();
          this.changeDetector.detectChanges();
        }),
    );
  }

  cancel(): void {
    this.reset();
    this.apikeyText = undefined;
    this.cancelRequested.emit();
    this.changeDetector.detectChanges();
  }

  reset(): void {
    this.apikeyCreateError = undefined;
    this.apikeyForm.reset();
    this.apikeyForm.patchValue({
      readOnly: true,
    });
    this.changeDetector.detectChanges();
  }

  startOver(): void {
    this.apikeyText = undefined;
    this.reset();
    this.changeDetector.detectChanges();
  }

  copyApiKey(): void {
    this.clipboardService.copy(this.apikeyText);
    this.messageService.add({ severity: 'success', summary: this.translateService.instant('pages.profile.apikeys.create.apikeyCopied') });
  }

  copyApiKeyHeader(): void {
    this.clipboardService.copy('X-API-KEY');
    this.messageService.add({ severity: 'success', summary: this.translateService.instant('pages.profile.apikeys.create.apikeyHeaderCopied') });
  }

  private setAvailableTypes(): void {
    this.availableTypes = [
      {
        name: this.translateService.instant('pages.apikeys.create.types.service'),
        value: 'Service',
      },
    ];
  }
}
