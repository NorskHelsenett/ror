import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AdminPricesCreateComponent } from './admin-prices-create.component';

describe('PricesNewComponent', () => {
  let component: AdminPricesCreateComponent;
  let fixture: ComponentFixture<AdminPricesCreateComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [AdminPricesCreateComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(AdminPricesCreateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
