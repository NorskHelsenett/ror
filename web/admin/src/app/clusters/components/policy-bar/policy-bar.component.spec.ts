import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PolicyBarComponent } from './policy-bar.component';

describe('PolicyBarComponent', () => {
  let component: PolicyBarComponent;
  let fixture: ComponentFixture<PolicyBarComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [PolicyBarComponent],
    });
    fixture = TestBed.createComponent(PolicyBarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
