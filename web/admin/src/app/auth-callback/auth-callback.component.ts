import { ChangeDetectionStrategy, Component, OnInit, ChangeDetectorRef } from '@angular/core';
import { Router } from '@angular/router';
import { OAuthService } from 'angular-oauth2-oidc';
import { AuthService } from '../core/services/auth.service';
import { filter } from 'rxjs';

@Component({
  selector: 'app-auth-callback',
  templateUrl: './auth-callback.component.html',
  styleUrls: ['./auth-callback.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AuthCallbackComponent implements OnInit {
  private relativePath = '../../';
  constructor(
    private changeDetector: ChangeDetectorRef,
    private oauthService: OAuthService,
    private authService: AuthService,
    private router: Router,
  ) {}

  ngOnInit(): void {
    this.oauthService.loadDiscoveryDocumentAndLogin();
    this.oauthService.setupAutomaticSilentRefresh();

    // Automatically load user profile
    this.oauthService.events.pipe(filter((e) => e?.type === 'token_received')).subscribe((_) => {
      this.oauthService.loadUserProfile();
      this.changeDetector.detectChanges();
      this.redirect();
      return;
    });

    this.changeDetector.detectChanges();
    this.redirect();
  }

  redirect(): void {
    setTimeout(() => {
      this.router.navigate([this.relativePath]);
      this.changeDetector.detectChanges();
    }, 0);
  }
}
