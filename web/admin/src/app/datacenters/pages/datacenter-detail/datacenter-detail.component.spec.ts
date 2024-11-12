import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DatacenterDetailComponent } from './datacenter-detail.component';

describe('DatacenterDetailComponent', () => {
  let component: DatacenterDetailComponent;
  let fixture: ComponentFixture<DatacenterDetailComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [DatacenterDetailComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(DatacenterDetailComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
