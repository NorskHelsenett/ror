import { OAuthService } from 'angular-oauth2-oidc';
import { TestBed } from '@angular/core/testing';

import { AuthInterceptor } from './auth-interceptor';

describe('AuthInterceptorService', () => {
  let service: AuthInterceptor;

  let mockOAuthService = {};

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [
        {
          provide: OAuthService,
          useValue: mockOAuthService,
        },
      ],
    });
    service = TestBed.inject(AuthInterceptor);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
