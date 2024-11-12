import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UserprofileApikeysComponent } from './userprofile-apikeys.component';

describe('UserprofileApikeysComponent', () => {
  let component: UserprofileApikeysComponent;
  let fixture: ComponentFixture<UserprofileApikeysComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [UserprofileApikeysComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(UserprofileApikeysComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
