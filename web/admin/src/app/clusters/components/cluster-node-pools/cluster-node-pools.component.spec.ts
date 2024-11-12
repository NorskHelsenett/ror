import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterNodePoolsComponent } from './cluster-node-pools.component';

describe('ClusterNodePoolsComponent', () => {
  let component: ClusterNodePoolsComponent;
  let fixture: ComponentFixture<ClusterNodePoolsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ClusterNodePoolsComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ClusterNodePoolsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
