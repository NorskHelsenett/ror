import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IconSortDescComponent } from './icon-sort-desc.component';

describe('IconSortDescComponent', () => {
  let component: IconSortDescComponent;
  let fixture: ComponentFixture<IconSortDescComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [IconSortDescComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(IconSortDescComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
