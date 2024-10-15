import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ResourceTableComponent } from './resource-table.component';

describe('ResourceTableComponent', () => {
  let component: ResourceTableComponent;
  let fixture: ComponentFixture<ResourceTableComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ResourceTableComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ResourceTableComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
