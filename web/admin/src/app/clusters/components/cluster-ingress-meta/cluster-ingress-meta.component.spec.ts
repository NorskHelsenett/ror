import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterIngressMetaComponent } from './cluster-ingress-meta.component';

describe('ClusterIngressMetaComponent', () => {
  let component: ClusterIngressMetaComponent;
  let fixture: ComponentFixture<ClusterIngressMetaComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ClusterIngressMetaComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ClusterIngressMetaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
