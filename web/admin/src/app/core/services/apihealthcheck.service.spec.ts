import { TestBed } from '@angular/core/testing';

import { AboutService } from './apihealthcheck.service';

describe('AboutService', () => {
  let service: AboutService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(AboutService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
