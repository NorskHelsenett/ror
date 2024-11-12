import { Component, Input } from '@angular/core';
import { TranslateService } from '@ngx-translate/core';
import { HighlightResult } from 'highlight.js';
import { ClipboardService } from 'ngx-clipboard';
import { MessageService } from 'primeng/api';

@Component({
  selector: 'app-cluster-raw',
  templateUrl: './cluster-raw.component.html',
  styleUrls: ['./cluster-raw.component.scss'],
})
export class ClusterRawComponent {
  @Input() cluster: any = undefined;

  response!: HighlightResult;

  constructor(
    private clipboardService: ClipboardService,
    private messageService: MessageService,
    private translateService: TranslateService,
  ) {}

  copy(): void {
    this.clipboardService.copyFromContent(JSON.stringify(this.cluster));
    this.messageService.add({
      severity: 'success',
      summary: this.translateService.instant('common.copied'),
    });
  }

  onHighlight(e: HighlightResult) {
    this.response = {
      language: e.language,
      relevance: e.relevance,
      illegal: e.illegal,
      secondBest: e.secondBest,
      _emitter: e._emitter,
      value: '{...}',
    };
  }
}
