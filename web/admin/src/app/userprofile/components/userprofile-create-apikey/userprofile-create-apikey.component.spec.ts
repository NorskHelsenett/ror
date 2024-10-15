import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UserprofileCreateApikeyComponent } from './userprofile-create-apikey.component';

describe('UserprofileCreateApikeyComponent', () => {
  let component: UserprofileCreateApikeyComponent;
  let fixture: ComponentFixture<UserprofileCreateApikeyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [UserprofileCreateApikeyComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(UserprofileCreateApikeyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
