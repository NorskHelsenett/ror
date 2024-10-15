import { TestBed } from '@angular/core/testing';

import { ColumnFactoryService } from './column-factory.service';

describe('ColumnFactoryService', () => {
  let service: ColumnFactoryService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ColumnFactoryService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
