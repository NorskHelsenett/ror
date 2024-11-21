import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterResourceTableComponent } from './cluster-resource-table.component';

describe('ClusterResourceTableComponent', () => {
  let component: ClusterResourceTableComponent;
  let fixture: ComponentFixture<ClusterResourceTableComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ClusterResourceTableComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ClusterResourceTableComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
