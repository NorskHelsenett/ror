import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterDetailsEditComponent } from './cluster-details-edit.component';

describe('ClusterDetailsEditComponent', () => {
  let component: ClusterDetailsEditComponent;
  let fixture: ComponentFixture<ClusterDetailsEditComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ClusterDetailsEditComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ClusterDetailsEditComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
