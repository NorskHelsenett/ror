import { ChangeDetectionStrategy, ChangeDetectorRef, Component, OnInit } from '@angular/core';
import { TranslateService } from '@ngx-translate/core';
import { ConfirmationService, MessageService } from 'primeng/api';
import { Observable, Subscription, catchError, finalize, share, tap } from 'rxjs';
import { User } from '../../../core/models/user';
import { Filter } from '../../../core/models/apiFilter';
import { ApiKey } from '../../../core/models/apikey';
import { PaginationResult } from '../../../core/models/paginatedResult';
import { ApikeysService } from '../../../core/services/apikeys.service';
import { ConfigService } from '../../../core/services/config.service';
import { UserService } from '../../../core/services/user.service';
import { SignalService } from '../../../create/create-cluster/services/signal.service';
import { FilterService } from '../../../core/services/filter.service';

@Component({
  selector: 'app-apikeys',
  templateUrl: './apikeys.component.html',
  styleUrls: ['./apikeys.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ApikeysComponent implements OnInit {
  user$: Observable<User> | undefined;
  refreshList: boolean;

  apikeys$: Observable<PaginationResult<ApiKey>> | undefined;
  fetchError: any;

  lastLazyLoad: any;
  loading: boolean;

  filter: Filter = {
    limit: 0,
    skip: 0,
    filters: [
      {
        field: 'type',
        matchMode: 'equals',
        value: 'Cluster',
      },
    ],
  };
  lastFilter: Filter;
  rows = this.configService.config.rows;
  rowsPerPage = this.configService.config.rowsPerPage;

  types: string[] = ['Cluster', 'User', 'Service'];
  selectedType: string = 'Cluster';

  createIsHidden = true;
  clusterCreated$: Observable<any> | undefined;

  private subscriptions = new Subscription();

  constructor(
    private changeDetector: ChangeDetectorRef,
    private userService: UserService,
    private filterService: FilterService,
    private apikeyService: ApikeysService,
    private confirmationService: ConfirmationService,
    private messageService: MessageService,
    private translateService: TranslateService,
    private configService: ConfigService,
    private signalService: SignalService,
  ) {}

  ngOnInit(): void {
    this.setupObservables();
    this.setupEventListening();
  }

  setupObservables(): void {
    this.user$ = this.userService?.getUser();
    if (!this.userService.user.value) {
      this.subscriptions.add(
        this.userService
          .getUser()
          .pipe(share())
          .subscribe((_) => {
            this.changeDetector.detectChanges();
          }),
      );
    }
  }

  fetch(event: any): void {
    if (event) {
      this.filter = this.filterService.mapFilter(event);
    }
    this.loading = true;
    this.lastFilter = this.filter;
    this.fetchError = undefined;
    this.apikeys$ = this.apikeyService.getByFilter(this.filter).pipe(
      share(),
      catchError((error: any) => {
        this.fetchError = error;
        throw error;
      }),
      finalize(() => {
        this.loading = false;
      }),
    );
  }

  delete(apikey: ApiKey): void {
    this.confirmationService.confirm({
      header: this.translateService.instant('pages.apikeys.delete.title'),
      message: this.translateService.instant('pages.apikeys.delete.details', { name: apikey?.displayName }),
      accept: () => {
        this.subscriptions.add(
          this.apikeyService
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
              this.fetch(this.lastLazyLoad);
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

  private setupEventListening(): void {
    this.clusterCreated$ = this.signalService.clusterCreated$.pipe(
      tap((data: any) => {
        this.fetch(this.lastLazyLoad);
        this.changeDetector.detectChanges();
      }),
    );
  }
}
