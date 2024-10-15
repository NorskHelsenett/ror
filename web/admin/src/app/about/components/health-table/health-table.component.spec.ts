import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HealthTableComponent } from './health-table.component';

describe('HealthTableComponent', () => {
  let component: HealthTableComponent;
  let fixture: ComponentFixture<HealthTableComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [HealthTableComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(HealthTableComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
