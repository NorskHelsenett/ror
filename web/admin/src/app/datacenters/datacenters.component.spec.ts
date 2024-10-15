import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DatacentersComponent } from './datacenters.component';

describe('DatacentersComponent', () => {
  let component: DatacentersComponent;
  let fixture: ComponentFixture<DatacentersComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [DatacentersComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(DatacentersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
