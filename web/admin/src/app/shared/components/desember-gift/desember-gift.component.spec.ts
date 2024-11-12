import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DesemberGiftComponent } from './desember-gift.component';

describe('DesemberGiftComponent', () => {
  let component: DesemberGiftComponent;
  let fixture: ComponentFixture<DesemberGiftComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DesemberGiftComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(DesemberGiftComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
