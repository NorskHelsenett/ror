import { OAuthService } from 'angular-oauth2-oidc';
import { TestBed } from '@angular/core/testing';
import { provideHttpClientTesting } from '@angular/common/http/testing';
import { ClustersService } from './clusters.service';
import { provideHttpClient, withInterceptorsFromDi } from '@angular/common/http';

describe('ClustersService', () => {
  let service: ClustersService;

  let mockOAuthService = {};

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [],
      providers: [{ provide: OAuthService, useValue: mockOAuthService }, provideHttpClient(withInterceptorsFromDi()), provideHttpClientTesting()],
    });
    service = TestBed.inject(ClustersService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
