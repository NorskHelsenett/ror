import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterIngressAnnotationsComponent } from './cluster-ingress-annotations.component';

describe('ClusterIngressAnnotationsComponent', () => {
  let component: ClusterIngressAnnotationsComponent;
  let fixture: ComponentFixture<ClusterIngressAnnotationsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ClusterIngressAnnotationsComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ClusterIngressAnnotationsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
