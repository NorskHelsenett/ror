import { RouterTestingModule } from '@angular/router/testing';
import { OAuthService } from 'angular-oauth2-oidc';
import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UserprofileComponent } from './userprofile.component';

describe('UserprofileComponent', () => {
  let component: UserprofileComponent;
  let fixture: ComponentFixture<UserprofileComponent>;

  let mockOAuthService = {
    getIdToken() {
      return '';
    },
    getAccessToken() {
      return '';
    },
    getIdentityClaims() {
      return null;
    },
    authorizationHeader() {
      return '';
    },
  };

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [UserprofileComponent],
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
    fixture = TestBed.createComponent(UserprofileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
