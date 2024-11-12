import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MetadataStepComponent } from './metadata-step.component';

describe('MetadataStepComponent', () => {
  let component: MetadataStepComponent;
  let fixture: ComponentFixture<MetadataStepComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [MetadataStepComponent],
    });
    fixture = TestBed.createComponent(MetadataStepComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
