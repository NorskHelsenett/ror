import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterACLComponent } from './cluster-acl.component';

describe('ClusterACLComponent', () => {
  let component: ClusterACLComponent;
  let fixture: ComponentFixture<ClusterACLComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ClusterACLComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ClusterACLComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
