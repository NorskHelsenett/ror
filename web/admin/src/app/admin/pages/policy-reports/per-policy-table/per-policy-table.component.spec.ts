import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PerPolicyTableComponent } from './per-policy-table.component';

describe('PerPolicyTableComponent', () => {
  let component: PerPolicyTableComponent;
  let fixture: ComponentFixture<PerPolicyTableComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [PerPolicyTableComponent],
    });
    fixture = TestBed.createComponent(PerPolicyTableComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
