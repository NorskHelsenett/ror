import { ChangeDetectionStrategy, Component, OnInit, ChangeDetectorRef, OnDestroy } from '@angular/core';
import { Location } from '@angular/common';
import { ActivatedRoute } from '@angular/router';
import { OAuthService } from 'angular-oauth2-oidc';
import { AclService } from '../core/services/acl.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-userprofile',
  templateUrl: './userprofile.component.html',
  styleUrls: ['./userprofile.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class UserprofileComponent implements OnInit, OnDestroy {
  idToken: string | undefined;
  accessToken: string | undefined;
  claims: any;
  authHeaders: any;
  expDate: Date = new Date(1970);
  iatDate: Date = new Date(1970);
  rawIsHidden = true;

  selectedTabIndex: number = 0;
  tabs: any[] = [
    {
      metadata: '',
      query: '',
    },
    {
      metadata: 'apikeys',
      query: 'tab=apikeys',
    },
  ];

  private subscriptions = new Subscription();

  constructor(
    private changeDetector: ChangeDetectorRef,
    private route: ActivatedRoute,
    private location: Location,
    private oauthService: OAuthService,
    private aclService: AclService,
  ) {}

  ngOnInit(): void {
    const tab = this.route.snapshot.queryParams['tab'];
    this.tabs.forEach((value: any, index: number) => {
      if (tab == value?.metadata) {
        this.selectedTabIndex = index;
      }
    });

    this.idToken = this.oauthService.getIdToken();
    this.accessToken = this.oauthService.getAccessToken();
    this.claims = this.oauthService.getIdentityClaims();
    if (this.claims != undefined) {
      this.expDate = new Date(this.claims?.exp * 1000);
      this.iatDate = new Date(this.claims?.iat * 1000);
      this.changeDetector.detectChanges();
    }
    this.authHeaders = this.oauthService.authorizationHeader();
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }

  switchTab(selectedIndex: number): void {
    try {
      const tab = this.tabs[selectedIndex];
      this.location.replaceState('userprofile', tab?.query);
    } catch {
      //ignoring
    }
  }
}
