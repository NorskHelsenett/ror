import { Component, Input, OnInit } from '@angular/core';
import { TranslateService } from '@ngx-translate/core';
import { ClipboardService } from 'ngx-clipboard';
import { MessageService } from 'primeng/api';

@Component({
  selector: 'app-userprofile-details',
  templateUrl: './userprofile-details.component.html',
  styleUrls: ['./userprofile-details.component.scss'],
})
export class UserprofileDetailsComponent implements OnInit {
  @Input()
  claims: any;

  @Input()
  exp: Date;

  @Input()
  iat: Date;

  @Input()
  accessToken: any;

  domain: string;

  constructor(
    private clipboardService: ClipboardService,
    private messageService: MessageService,
    private translateService: TranslateService,
  ) {}

  ngOnInit(): void {
    if (!this.claims) {
      return;
    }
    const emailArray = this.claims?.email?.split('@');
    if (emailArray?.length > 1) {
      this.domain = emailArray[1];
    }
  }

  copyAccessToken(withBearer: boolean): void {
    if (withBearer) {
      this.clipboardService.copyFromContent(`Bearer ${this.accessToken}` || '');
    } else {
      this.clipboardService.copyFromContent(this.accessToken || '');
    }
    this.messageService.add({
      severity: 'success',
      summary: this.translateService.instant('common.copied'),
    });
  }
}
