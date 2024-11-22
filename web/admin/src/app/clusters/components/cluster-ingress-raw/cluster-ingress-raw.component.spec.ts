import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterIngressRawComponent } from './cluster-ingress-raw.component';

describe('ClusterIngressRawComponent', () => {
  let component: ClusterIngressRawComponent;
  let fixture: ComponentFixture<ClusterIngressRawComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ClusterIngressRawComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ClusterIngressRawComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
