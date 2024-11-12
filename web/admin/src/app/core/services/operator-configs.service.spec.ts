import { TestBed } from '@angular/core/testing';

import { OperatorConfigsService } from './operator-configs.service';

describe('OperatorConfigsService', () => {
  let service: OperatorConfigsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(OperatorConfigsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
