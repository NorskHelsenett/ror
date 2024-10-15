import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SummaryStepComponent } from './summary-step.component';

describe('SummaryStepComponent', () => {
  let component: SummaryStepComponent;
  let fixture: ComponentFixture<SummaryStepComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [SummaryStepComponent],
    });
    fixture = TestBed.createComponent(SummaryStepComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
