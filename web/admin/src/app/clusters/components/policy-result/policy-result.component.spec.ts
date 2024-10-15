import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PolicyResultComponent } from './policy-result.component';

describe('PolicyResultComponent', () => {
  let component: PolicyResultComponent;
  let fixture: ComponentFixture<PolicyResultComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [PolicyResultComponent],
    });
    fixture = TestBed.createComponent(PolicyResultComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
