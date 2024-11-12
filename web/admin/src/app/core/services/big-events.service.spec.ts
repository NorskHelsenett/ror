import { TestBed } from '@angular/core/testing';

import { BigEventsService } from './big-events.service';

describe('BigEventsService', () => {
  let service: BigEventsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(BigEventsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
