import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AdminPricesComponent } from './admin-prices.component';

describe('AdminPricesComponent', () => {
  let component: AdminPricesComponent;
  let fixture: ComponentFixture<AdminPricesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [AdminPricesComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(AdminPricesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
