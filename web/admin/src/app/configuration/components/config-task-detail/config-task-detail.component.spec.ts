import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ConfigTaskDetailComponent } from './config-task-detail.component';

describe('ConfigTaskDetailComponent', () => {
  let component: ConfigTaskDetailComponent;
  let fixture: ComponentFixture<ConfigTaskDetailComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ConfigTaskDetailComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ConfigTaskDetailComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
