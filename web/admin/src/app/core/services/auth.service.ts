import { Injectable } from '@angular/core';
import { AuthConfig, OAuthService } from 'angular-oauth2-oidc';
import { Subject } from 'rxjs';
import { UserService } from './user.service';
import { ConfigService } from './config.service';
import { jwtDecode } from 'jwt-decode';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  authenticationEventObservable: Subject<boolean> = new Subject<boolean>();
  authConfig: AuthConfig = {
    issuer: this.configService.config.auth.issuer,
    redirectUri: this.configService.config.auth.redirectUri,
    clientId: this.configService.config.auth.clientId,
    responseType: this.configService.config.auth.responseType,
    scope: this.configService.config.auth.scope,
    showDebugInformation: false,
    timeoutFactor: 0.75,
    postLogoutRedirectUri: this.configService.config.auth.postLogoutRedirectUri,
    logoutUrl: this.configService.config.auth.logoutUrl,
    requireHttps: this.configService.config.auth.requireHttps,
    strictDiscoveryDocumentValidation: this.configService.config.auth.strictDiscoveryDocumentValidation,
  };

  constructor(
    private oauthService: OAuthService,
    private userService: UserService,
    private configService: ConfigService,
  ) {
    this.oauthService.configure(this.authConfig);
  }

  getToken(): string | null {
    return localStorage.getItem('access_token');
  }

  logout() {
    this.userService.user.next(null);
    this.oauthService.logOut();
    window.location.reload();
  }

  isAuthenticated(): boolean {
    return this.oauthService.hasValidAccessToken() && this.oauthService.hasValidIdToken();
  }

  login() {
    this.oauthService
      .loadDiscoveryDocumentAndLogin()
      .then((result: boolean) => {
        this.authenticationEventObservable.next(result);
      })
      .catch((error: any) => {
        this.logout();
      });

    // Optional
    this.oauthService.setupAutomaticSilentRefresh();
  }

  refresh() {
    this.oauthService.refreshToken();
  }

  isTokenExpired(): boolean {
    const token = this.getToken();
    if (!token) {
      return true;
    }
    const decodedToken = jwtDecode(token);
    if (!decodedToken) {
      return true;
    }

    const now = new Date();
    const exp = new Date(0);
    exp.setUTCSeconds(decodedToken.exp);
    return now > exp;
  }
}
