import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterIngressCertmanagerComponent } from './cluster-ingress-certmanager.component';

describe('ClusterIngressCertmanagerComponent', () => {
  let component: ClusterIngressCertmanagerComponent;
  let fixture: ComponentFixture<ClusterIngressCertmanagerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ClusterIngressCertmanagerComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ClusterIngressCertmanagerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
