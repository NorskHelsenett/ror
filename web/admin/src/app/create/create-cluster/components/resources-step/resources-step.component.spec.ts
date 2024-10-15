import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ResourcesStepComponent } from './resources-step.component';

describe('ResourcesStepComponent', () => {
  let component: ResourcesStepComponent;
  let fixture: ComponentFixture<ResourcesStepComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ResourcesStepComponent],
    });
    fixture = TestBed.createComponent(ResourcesStepComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
