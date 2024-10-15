import { AuthConfig, LoginOptions, OAuthEvent, OAuthService } from 'angular-oauth2-oidc';
import { TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { AppComponent } from './app.component';
import { Observable } from 'rxjs';

describe('AppComponent', () => {
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
      imports: [RouterTestingModule],
      declarations: [AppComponent],
      providers: [
        {
          provide: OAuthService,
          useValue: mockOAuthService,
        },
      ],
    }).compileComponents();
  });

  it('should create the app', () => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.componentInstance;
    expect(app).toBeTruthy();
  });
});
