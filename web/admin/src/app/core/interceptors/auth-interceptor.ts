import { OAuthService } from 'angular-oauth2-oidc';
import { Injectable } from '@angular/core';
import { HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from '@angular/common/http';
import { Observable } from 'rxjs';
import { ConfigService } from '../services/config.service';
import { AuthService } from '../services/auth.service';

@Injectable({
  providedIn: 'root',
})
export class AuthInterceptor implements HttpInterceptor {
  constructor(
    private oauthService: OAuthService,
    private authService: AuthService,
    private configService: ConfigService,
  ) {}

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    if (!req.url.includes(this.configService.config.rorApi)) {
      return next.handle(req);
    }

    if (this.authService.isTokenExpired()) {
      this.authService.logout();
    }

    req = req.clone({
      setHeaders: {
        'Content-Type': 'application/json; charset=utf-8',
        Accept: 'application/json',
        Authorization: this.oauthService.authorizationHeader(),
      },
    });

    return next.handle(req);
  }
}
