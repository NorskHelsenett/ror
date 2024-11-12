import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UserprofileRawComponent } from './userprofile-raw.component';

describe('UserprofileRawComponent', () => {
  let component: UserprofileRawComponent;
  let fixture: ComponentFixture<UserprofileRawComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [UserprofileRawComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(UserprofileRawComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
