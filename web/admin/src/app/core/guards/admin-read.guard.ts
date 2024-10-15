import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, Route, Router, RouterStateSnapshot, UrlSegment, UrlTree } from '@angular/router';
import { Observable, tap } from 'rxjs';
import { AclService } from '../services/acl.service';
import { AclAccess, AclScopes } from '../models/acl-scopes';

@Injectable({
  providedIn: 'root',
})
export class AdminReadGuard {
  constructor(
    private router: Router,
    private aclService: AclService,
  ) {}

  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot,
  ): Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {
    return this.aclService.check(AclScopes.ROR, AclScopes.Global, AclAccess.Read).pipe(
      tap((result: boolean) => {
        if (result === false) {
          this.router.navigateByUrl('/error/401');
          return false;
        } else {
          return true;
        }
      }),
    );
  }

  canLoad(route: Route, segments: UrlSegment[]): boolean | UrlTree | Observable<boolean | UrlTree> | Promise<boolean | UrlTree> {
    return this.aclService.check(AclScopes.ROR, AclScopes.Global, AclAccess.Read).pipe(
      tap((result: boolean) => {
        if (result === false) {
          this.router.navigateByUrl('/error/401');
          return false;
        } else {
          return true;
        }
      }),
    );
  }
}
