import { TestBed } from '@angular/core/testing';

import { AdminReadGuard } from './admin-read.guard';

describe('AdminReadGuard', () => {
  let guard: AdminReadGuard;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    guard = TestBed.inject(AdminReadGuard);
  });

  it('should be created', () => {
    expect(guard).toBeTruthy();
  });
});
