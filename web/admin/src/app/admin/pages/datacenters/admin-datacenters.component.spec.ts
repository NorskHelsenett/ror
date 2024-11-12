import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AdminDatacentersComponent } from './admin-datacenters.component';

describe('AdminDatacentersComponent', () => {
  let component: AdminDatacentersComponent;
  let fixture: ComponentFixture<AdminDatacentersComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [AdminDatacentersComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(AdminDatacentersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
