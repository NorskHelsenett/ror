import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterNotificationsComponent } from './cluster-notifications.component';

describe('ClusterNotificationsComponent', () => {
  let component: ClusterNotificationsComponent;
  let fixture: ComponentFixture<ClusterNotificationsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ClusterNotificationsComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ClusterNotificationsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
