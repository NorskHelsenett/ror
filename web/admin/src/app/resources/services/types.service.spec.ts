import { TestBed } from '@angular/core/testing';

import { TypesService } from './types.service';

describe('TypesService', () => {
  let service: TypesService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TypesService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
