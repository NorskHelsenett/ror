import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ApikeysComponent } from './apikeys.component';

describe('ApikeysComponent', () => {
  let component: ApikeysComponent;
  let fixture: ComponentFixture<ApikeysComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ApikeysComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ApikeysComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
