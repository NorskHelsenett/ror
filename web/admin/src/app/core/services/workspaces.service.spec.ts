import { TestBed } from '@angular/core/testing';

import { WorkspacesService } from './workspaces.service';

describe('WorkspacesService', () => {
  let service: WorkspacesService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(WorkspacesService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
