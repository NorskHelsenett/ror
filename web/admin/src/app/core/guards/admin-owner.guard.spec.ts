import { TestBed } from '@angular/core/testing';

import { AdminOwnerGuard } from './admin-owner.guard';

describe('AdminOwnerGuard', () => {
  let guard: AdminOwnerGuard;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    guard = TestBed.inject(AdminOwnerGuard);
  });

  it('should be created', () => {
    expect(guard).toBeTruthy();
  });
});
