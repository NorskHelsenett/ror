import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ConfigTaskListComponent } from './config-task-list.component';

describe('ConfigTaskListComponent', () => {
  let component: ConfigTaskListComponent;
  let fixture: ComponentFixture<ConfigTaskListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ConfigTaskListComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ConfigTaskListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
