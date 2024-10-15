import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterComplianceReportComponent } from './cluster-compliance-report.component';

describe('ClusterComplianceReportComponent', () => {
  let component: ClusterComplianceReportComponent;
  let fixture: ComponentFixture<ClusterComplianceReportComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ClusterComplianceReportComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ClusterComplianceReportComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
