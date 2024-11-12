import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DatacenterTableComponent } from './datacenter-table.component';

describe('DatacenterTableComponent', () => {
  let component: DatacenterTableComponent;
  let fixture: ComponentFixture<DatacenterTableComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [DatacenterTableComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(DatacenterTableComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
