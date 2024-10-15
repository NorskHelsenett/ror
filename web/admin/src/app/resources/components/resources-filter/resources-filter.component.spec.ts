import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ResourcesFilterComponent } from './resources-filter.component';

describe('ResourcesFilterComponent', () => {
  let component: ResourcesFilterComponent;
  let fixture: ComponentFixture<ResourcesFilterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ResourcesFilterComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ResourcesFilterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
