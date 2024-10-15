import { TestBed } from '@angular/core/testing';

import { ComplianceReportsService } from './compliance-reports.service';

describe('ComplianceReportsService', () => {
  let service: ComplianceReportsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ComplianceReportsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
