import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PolicyReportsComponent } from './policy-reports.component';

describe('PolicyReportsComponent', () => {
  let component: PolicyReportsComponent;
  let fixture: ComponentFixture<PolicyReportsComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [PolicyReportsComponent],
    });
    fixture = TestBed.createComponent(PolicyReportsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
