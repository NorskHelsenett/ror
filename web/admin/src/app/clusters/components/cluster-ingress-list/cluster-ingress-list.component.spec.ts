import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterIngressListComponent } from './cluster-ingress-list.component';

describe('ClusterIngressListComponent', () => {
  let component: ClusterIngressListComponent;
  let fixture: ComponentFixture<ClusterIngressListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ClusterIngressListComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ClusterIngressListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
