import { TestBed } from '@angular/core/testing';

import { DatacenterService } from './datacenter.service';

describe('DatacenterService', () => {
  let service: DatacenterService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(DatacenterService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
