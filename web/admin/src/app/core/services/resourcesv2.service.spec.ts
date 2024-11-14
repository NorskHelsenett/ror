import { TestBed } from '@angular/core/testing';

import { Resourcesv2Service } from './resourcesv2.service';

describe('Resourcesv2Service', () => {
  let service: Resourcesv2Service;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(Resourcesv2Service);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
