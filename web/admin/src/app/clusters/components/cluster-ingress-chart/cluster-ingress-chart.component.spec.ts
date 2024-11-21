import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterIngressChartComponent } from './cluster-ingress-chart.component';

describe('ClusterIngressChartComponent', () => {
  let component: ClusterIngressChartComponent;
  let fixture: ComponentFixture<ClusterIngressChartComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ClusterIngressChartComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ClusterIngressChartComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
