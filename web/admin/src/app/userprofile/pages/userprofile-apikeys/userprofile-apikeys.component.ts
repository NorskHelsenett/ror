import { ChangeDetectorRef, Component, Input, OnInit } from '@angular/core';
import { TranslateService } from '@ngx-translate/core';
import { ConfirmationService, LazyLoadEvent, MessageService } from 'primeng/api';
import { Observable, Subscription, catchError, finalize, share, tap } from 'rxjs';
import { FilterService } from '../../../core/services/filter.service';
import { Filter } from '../../../core/models/apiFilter';
import { ApiKey } from '../../../core/models/apikey';
import { PaginationResult } from '../../../core/models/paginatedResult';
import { ConfigService } from '../../../core/services/config.service';
import { UserApikeysService } from '../../../core/services/user-apikeys.service';

@Component({
  selector: 'app-userprofile-apikeys',
  templateUrl: './userprofile-apikeys.component.html',
  styleUrls: ['./userprofile-apikeys.component.scss'],
})
export class UserprofileApikeysComponent implements OnInit {
  @Input()
  upn: string;

  apikeys$: Observable<PaginationResult<ApiKey>> | undefined;
  apikeysFetchError: any;
  lastLazyLoad: LazyLoadEvent;
  filter: Filter;
  lastFilter: Filter;

  rows = this.configService.config.rows;
  rowsPerPage = this.configService.config.rowsPerPage;
  loading: boolean;

  createIsHidden = true;

  private subscriptions = new Subscription();

  constructor(
    private changeDetector: ChangeDetectorRef,
    private userApikeysService: UserApikeysService,
    private filterService: FilterService,
    private confirmationService: ConfirmationService,
    private translateService: TranslateService,
    private messageService: MessageService,
    private configService: ConfigService,
  ) {}

  ngOnInit(): void {
    this.changeDetector.detectChanges();
  }

  fetchApikeys(event: any): void {
    this.filter = this.filterService.mapFilter(event);
    this.apikeysFetchError = undefined;
    this.loading = true;
    this.lastFilter = this.filter;
    this.lastLazyLoad = event;

    this.apikeys$ = this.userApikeysService.getByFilter(this.filter).pipe(
      share(),
      tap(() => {
        this.loading = false;
        this.changeDetector.detectChanges();
      }),
      catchError((error: any) => {
        this.apikeysFetchError = error;

        throw error;
      }),
      finalize(() => {
        this.loading = false;
        this.changeDetector.detectChanges();
      }),
    );
  }

  deleteApiKey(apikey: ApiKey): void {
    this.confirmationService.confirm({
      header: this.translateService.instant('pages.apikeys.delete.title'),
      message: this.translateService.instant('pages.apikeys.delete.details', { name: apikey?.displayName }),
      accept: () => {
        this.subscriptions.add(
          this.userApikeysService
            .delete(apikey.id)
            .pipe(
              catchError((error: any) => {
                this.messageService.add({
                  severity: 'error',
                  summary: this.translateService.instant('pages.apikeys.delete.errortitle'),
                  detail: this.translateService.instant('pages.apikeys.delete.errordetails'),
                });
                this.changeDetector.detectChanges();
                throw error;
              }),
            )
            .subscribe(() => {
              this.fetchApikeys(this.lastFilter);
              this.messageService.add({
                severity: 'success',
                summary: this.translateService.instant('pages.apikeys.delete.success'),
              });
              this.changeDetector.detectChanges();
            }),
        );
      },
    });
  }

  toggleCreateVisibility(): void {
    this.createIsHidden = !this.createIsHidden;
  }
}
