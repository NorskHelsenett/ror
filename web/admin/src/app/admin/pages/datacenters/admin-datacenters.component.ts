import { DatacenterService } from '../../../core/services/datacenter.service';
import { catchError, finalize, Observable, share } from 'rxjs';
import { Component, OnInit, ChangeDetectorRef, ChangeDetectionStrategy } from '@angular/core';
import { AclScopes, AclAccess } from '../../../core/models/acl-scopes';
import { AclService } from '../../../core/services/acl.service';
import { ConfigService } from '../../../core/services/config.service';

@Component({
  selector: 'app-admin-datacenters',
  templateUrl: './admin-datacenters.component.html',
  styleUrls: ['./admin-datacenters.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AdminDatacentersComponent implements OnInit {
  datacenters$: Observable<any>;
  datacentersError: any;
  adminRead$: Observable<boolean> | undefined;
  adminReadFetchError: any;

  rows = this.configService.config.rows;
  rowsPerPage = this.configService.config.rowsPerPage;
  loading: boolean;

  constructor(
    private changeDetector: ChangeDetectorRef,
    private datacentersService: DatacenterService,
    private aclService: AclService,
    private configService: ConfigService,
  ) {}

  ngOnInit(): void {
    this.fetch();
    this.adminReadFetchError = undefined;
    this.adminRead$ = this.aclService.check(AclScopes.ROR, AclScopes.Global, AclAccess.Read).pipe(
      share(),
      catchError((error: any) => {
        this.adminReadFetchError = error;
        this.changeDetector.detectChanges();
        throw error;
      }),
    );
  }

  fetch(): void {
    this.loading = true;
    this.datacentersError = undefined;
    this.datacenters$ = this.datacentersService.get().pipe(
      share(),
      catchError((error) => {
        this.datacentersError = error;
        return error;
      }),
      finalize(() => {
        this.loading = false;
        this.changeDetector.detectChanges();
      }),
    );
  }
}
