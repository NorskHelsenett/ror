import { TestBed } from '@angular/core/testing';

import { AdminCreateGuard } from './admin-create.guard';

describe('AdminCreateGuard', () => {
  let guard: AdminCreateGuard;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    guard = TestBed.inject(AdminCreateGuard);
  });

  it('should be created', () => {
    expect(guard).toBeTruthy();
  });
});
