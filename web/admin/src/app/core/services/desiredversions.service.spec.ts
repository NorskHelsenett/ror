import { TestBed } from '@angular/core/testing';

import { DesiredversionsService } from './desiredversions.service';

describe('DesiredversionsService', () => {
  let service: DesiredversionsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(DesiredversionsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
