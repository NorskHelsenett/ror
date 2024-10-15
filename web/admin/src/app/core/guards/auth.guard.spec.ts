import { AuthConfig, LoginOptions, OAuthEvent, OAuthService } from 'angular-oauth2-oidc';
import { RouterTestingModule } from '@angular/router/testing';
import { TestBed } from '@angular/core/testing';

import { AuthGuard } from './auth.guard';
import { Observable } from 'rxjs';

describe('AuthGuard', () => {
  let guard: AuthGuard;

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
    guard = TestBed.inject(AuthGuard);
  });

  it('should be created', () => {
    expect(guard).toBeTruthy();
  });
});
