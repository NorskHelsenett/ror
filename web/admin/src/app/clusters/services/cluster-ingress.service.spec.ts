import { TestBed } from '@angular/core/testing';

import { ClusterIngressService } from './cluster-ingress.service';

describe('ClusterIngressService', () => {
  let service: ClusterIngressService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ClusterIngressService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
