import { TestBed } from '@angular/core/testing';

import { ResourcesV2QueryService } from './resources-v2-query.service';

describe('ResourcesV2QueryService', () => {
  let service: ResourcesV2QueryService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ResourcesV2QueryService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
