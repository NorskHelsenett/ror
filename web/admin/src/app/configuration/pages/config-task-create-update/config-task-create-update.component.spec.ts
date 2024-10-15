import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ConfigTaskCreateUpdateComponent } from './config-task-create-update.component';

describe('ConfigTaskCreateUpdateComponent', () => {
  let component: ConfigTaskCreateUpdateComponent;
  let fixture: ComponentFixture<ConfigTaskCreateUpdateComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ConfigTaskCreateUpdateComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ConfigTaskCreateUpdateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
