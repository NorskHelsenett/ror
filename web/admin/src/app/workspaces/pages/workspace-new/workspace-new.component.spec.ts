import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WorkspaceNewComponent } from './workspace-new.component';

describe('WorkspaceNewComponent', () => {
  let component: WorkspaceNewComponent;
  let fixture: ComponentFixture<WorkspaceNewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [WorkspaceNewComponent],
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(WorkspaceNewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
