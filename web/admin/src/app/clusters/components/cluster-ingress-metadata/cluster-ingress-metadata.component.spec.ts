import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterIngressMetadataComponent } from './cluster-ingress-metadata.component';

describe('ClusterIngressMetadataComponent', () => {
  let component: ClusterIngressMetadataComponent;
  let fixture: ComponentFixture<ClusterIngressMetadataComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ClusterIngressMetadataComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ClusterIngressMetadataComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
