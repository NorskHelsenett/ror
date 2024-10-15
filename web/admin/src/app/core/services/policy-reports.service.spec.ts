import { TestBed } from '@angular/core/testing';

import { PolicyReportsService } from './policy-reports.service';

describe('PolicyReportsService', () => {
  let service: PolicyReportsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PolicyReportsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
