import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterIngressDetailsComponent } from './cluster-ingress-details.component';

describe('ClusterIngressDetailsComponent', () => {
  let component: ClusterIngressDetailsComponent;
  let fixture: ComponentFixture<ClusterIngressDetailsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ClusterIngressDetailsComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ClusterIngressDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
