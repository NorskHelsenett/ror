import { ChangeDetectionStrategy, ChangeDetectorRef, Component, OnDestroy, OnInit } from '@angular/core';
import { TranslateService } from '@ngx-translate/core';
import { ConfirmationService, MessageService } from 'primeng/api';
import { Observable, Subscription, catchError, finalize, share } from 'rxjs';
import { AclV2 } from '../../../core/models/aclv2';
import { Filter } from '../../../core/models/apiFilter';
import { PaginationResult } from '../../../core/models/paginatedResult';
import { User } from '../../../core/models/user';
import { AclService } from '../../../core/services/acl.service';
import { ConfigService } from '../../../core/services/config.service';
import { UserService } from '../../../core/services/user.service';
import { FilterService } from '../../../core/services/filter.service';

@Component({
  selector: 'app-acl',
  templateUrl: './acl.component.html',
  styleUrls: ['./acl.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AclComponent implements OnInit, OnDestroy {
  user$: Observable<User> | undefined;
  acl$: Observable<PaginationResult<AclV2>> | undefined;
  fetchError: any;

  lastLazyLoad: any;
  loading: boolean;

  filter: Filter = {
    limit: 0,
    skip: 0,
  };
  lastFilter: Filter;
  rows = this.configService.config.rows;
  rowsPerPage = this.configService.config.rowsPerPage;

  accessTypes: any[];

  private subscriptions = new Subscription();

  constructor(
    private changeDetector: ChangeDetectorRef,
    private userService: UserService,
    private aclService: AclService,
    private filterService: FilterService,
    private confirmationService: ConfirmationService,
    private messageService: MessageService,
    private translateService: TranslateService,
    private configService: ConfigService,
  ) {}

  ngOnInit(): void {
    this.setupTypes();
    this.setupObservables();

    this.subscriptions.add(this.translateService.onLangChange.subscribe(() => this.setupTypes()));
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

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }

  fetch(event?: any): void {
    if (event) {
      this.filter = this.filterService.mapFilter(event);
    }
    this.loading = true;
    this.lastFilter = this.filter;
    this.fetchError = undefined;
    this.acl$ = this.aclService.getByFilter(this.filter).pipe(
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

  delete(acl: AclV2): void {
    this.confirmationService.confirm({
      header: this.translateService.instant('pages.acl.delete.title'),
      message: this.translateService.instant('pages.acl.delete.details', { name: acl?.group }),
      accept: () => {
        this.subscriptions.add(
          this.aclService
            .delete(acl?.id)
            .pipe(
              catchError((error: any) => {
                this.messageService.add({
                  severity: 'error',
                  summary: this.translateService.instant('pages.acl.delete.errortitle'),
                  detail: this.translateService.instant('pages.acl.delete.errordetails'),
                });
                this.changeDetector.detectChanges();
                throw error;
              }),
            )
            .subscribe(() => {
              this.fetch(this.lastLazyLoad);
              this.messageService.add({
                severity: 'success',
                summary: this.translateService.instant('pages.acl.delete.success'),
              });
              this.changeDetector.detectChanges();
            }),
        );
      },
    });
  }

  setupTypes(): void {
    this.accessTypes = [
      {
        name: this.translateService.instant('shared.trueOrFalse.true'),
        value: true,
      },
      {
        name: this.translateService.instant('shared.trueOrFalse.false'),
        value: false,
      },
    ];
  }
}
