import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterIngressRulesComponent } from './cluster-ingress-rules.component';

describe('ClusterIngressRulesComponent', () => {
  let component: ClusterIngressRulesComponent;
  let fixture: ComponentFixture<ClusterIngressRulesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ClusterIngressRulesComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ClusterIngressRulesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
