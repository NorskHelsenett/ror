import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IconSortAscComponent } from './icon-sort-asc.component';

describe('IconSortAscComponent', () => {
  let component: IconSortAscComponent;
  let fixture: ComponentFixture<IconSortAscComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [IconSortAscComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(IconSortAscComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
