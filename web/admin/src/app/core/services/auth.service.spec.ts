import { AuthConfig, LoginOptions, OAuthEvent, OAuthService } from 'angular-oauth2-oidc';
import { RouterTestingModule } from '@angular/router/testing';
import { TestBed } from '@angular/core/testing';

import { AuthService } from './auth.service';
import { Observable } from 'rxjs';

describe('AuthService', () => {
  let service: AuthService;

  let mockOAuthService = {
    configure(config: AuthConfig): void {
      return;
    },
    async loadDiscoveryDocumentAndLogin(
      options?: LoginOptions & {
        state?: string;
      },
    ): Promise<boolean> {
      return Promise.resolve(false);
    },
    events: new Observable<OAuthEvent>(),
  };

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [RouterTestingModule],
      providers: [
        {
          provide: OAuthService,
          useValue: mockOAuthService,
        },
      ],
    });
    service = TestBed.inject(AuthService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
