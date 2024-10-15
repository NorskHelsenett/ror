import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterConditionsComponent } from './cluster-conditions.component';

describe('ClusterConditionsComponent', () => {
  let component: ClusterConditionsComponent;
  let fixture: ComponentFixture<ClusterConditionsComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ClusterConditionsComponent],
    });
    fixture = TestBed.createComponent(ClusterConditionsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
