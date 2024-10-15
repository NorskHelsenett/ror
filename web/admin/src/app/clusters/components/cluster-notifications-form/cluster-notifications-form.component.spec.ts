import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterNotificationsFormComponent } from './cluster-notifications-form.component';

describe('ClusterNotificationsFormComponent', () => {
  let component: ClusterNotificationsFormComponent;
  let fixture: ComponentFixture<ClusterNotificationsFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ClusterNotificationsFormComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ClusterNotificationsFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
