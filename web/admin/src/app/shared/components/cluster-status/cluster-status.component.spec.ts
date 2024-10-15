import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterStatusComponent } from './cluster-status.component';

describe('ClusterStatusComponent', () => {
  let component: ClusterStatusComponent;
  let fixture: ComponentFixture<ClusterStatusComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ClusterStatusComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ClusterStatusComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
