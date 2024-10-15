import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PolicyPolicyComponent } from './policy-policy.component';

describe('PolicyPolicyComponent', () => {
  let component: PolicyPolicyComponent;
  let fixture: ComponentFixture<PolicyPolicyComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [PolicyPolicyComponent],
    });
    fixture = TestBed.createComponent(PolicyPolicyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
