import { TestBed } from '@angular/core/testing';

import { ApikeysService } from './apikeys.service';

describe('ApikeysService', () => {
  let service: ApikeysService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ApikeysService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
