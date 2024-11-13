import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Resourcesv2FilterComponent } from './resourcesv2-filter.component';

describe('Resourcesv2FilterComponent', () => {
  let component: Resourcesv2FilterComponent;
  let fixture: ComponentFixture<Resourcesv2FilterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Resourcesv2FilterComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(Resourcesv2FilterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
