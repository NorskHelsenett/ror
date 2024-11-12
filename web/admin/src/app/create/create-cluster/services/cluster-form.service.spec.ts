import { TestBed } from '@angular/core/testing';

import { ClusterFormService } from './cluster-form.service';

describe('ClusterFormService', () => {
  let service: ClusterFormService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ClusterFormService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
