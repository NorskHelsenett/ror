import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AclCreateUpdateComponent } from './acl-create-update.component';

describe('AclCreateUpdateComponent', () => {
  let component: AclCreateUpdateComponent;
  let fixture: ComponentFixture<AclCreateUpdateComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [AclCreateUpdateComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(AclCreateUpdateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
