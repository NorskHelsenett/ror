import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AdminDatacenterCreateComponent } from './admin-datacenter-create.component';

describe('AdminDatacenterCreateComponent', () => {
  let component: AdminDatacenterCreateComponent;
  let fixture: ComponentFixture<AdminDatacenterCreateComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [AdminDatacenterCreateComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(AdminDatacenterCreateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
