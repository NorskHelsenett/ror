import { AuthConfig, LoginOptions, OAuthEvent, OAuthService } from 'angular-oauth2-oidc';
import { RouterTestingModule } from '@angular/router/testing';
import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AuthCallbackComponent } from './auth-callback.component';
import { Observable } from 'rxjs';

describe('AuthCallbackComponent', () => {
  let component: AuthCallbackComponent;
  let fixture: ComponentFixture<AuthCallbackComponent>;

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

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [AuthCallbackComponent],
      imports: [RouterTestingModule],
      providers: [
        {
          provide: OAuthService,
          useValue: mockOAuthService,
        },
      ],
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AuthCallbackComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
