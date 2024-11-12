import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PricesComponent } from './prices.component';

describe('PricesComponent', () => {
  let component: PricesComponent;
  let fixture: ComponentFixture<PricesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [PricesComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(PricesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
