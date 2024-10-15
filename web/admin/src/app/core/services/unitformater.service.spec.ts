import { TestBed } from '@angular/core/testing';

import { UnitformaterService } from './unitformater.service';

describe('UnitformaterService', () => {
  let service: UnitformaterService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(UnitformaterService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
