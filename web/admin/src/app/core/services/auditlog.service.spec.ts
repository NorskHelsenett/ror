import { TestBed } from '@angular/core/testing';

import { AuditlogService } from './auditlog.service';

describe('AuditlogService', () => {
  let service: AuditlogService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(AuditlogService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
