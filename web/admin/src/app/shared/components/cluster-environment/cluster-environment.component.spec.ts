import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterEnvironmentComponent } from './cluster-environment.component';

describe('ClusterEnvironmentComponent', () => {
  let component: ClusterEnvironmentComponent;
  let fixture: ComponentFixture<ClusterEnvironmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ClusterEnvironmentComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ClusterEnvironmentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
