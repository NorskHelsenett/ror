import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ApikeyCreateComponent } from './apikey-create.component';

describe('ApikeyCreateComponent', () => {
  let component: ApikeyCreateComponent;
  let fixture: ComponentFixture<ApikeyCreateComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ApikeyCreateComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ApikeyCreateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
