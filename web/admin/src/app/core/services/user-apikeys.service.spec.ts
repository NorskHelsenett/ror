import { TestBed } from '@angular/core/testing';

import { UserApikeysService } from './user-apikeys.service';

describe('UserApikeysService', () => {
  let service: UserApikeysService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(UserApikeysService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
