import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AdminAuditlogsComponent } from './admin-auditlogs.component';

describe('AdminAuditlogsComponent', () => {
  let component: AdminAuditlogsComponent;
  let fixture: ComponentFixture<AdminAuditlogsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [AdminAuditlogsComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(AdminAuditlogsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
