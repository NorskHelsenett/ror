import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterIngressTableComponent } from './cluster-ingress-table.component';

describe('ClusterIngressTableComponent', () => {
  let component: ClusterIngressTableComponent;
  let fixture: ComponentFixture<ClusterIngressTableComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ClusterIngressTableComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ClusterIngressTableComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
