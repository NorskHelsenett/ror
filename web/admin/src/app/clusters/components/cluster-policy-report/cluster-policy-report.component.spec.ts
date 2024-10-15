import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterPolicyReportComponent } from './cluster-policy-report.component';

describe('ClusterPolicyReportComponent', () => {
  let component: ClusterPolicyReportComponent;
  let fixture: ComponentFixture<ClusterPolicyReportComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ClusterPolicyReportComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ClusterPolicyReportComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
