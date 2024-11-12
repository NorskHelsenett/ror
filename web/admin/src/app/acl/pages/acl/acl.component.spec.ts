import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AclComponent } from './acl.component';

describe('AclComponent', () => {
  let component: AclComponent;
  let fixture: ComponentFixture<AclComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [AclComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(AclComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
