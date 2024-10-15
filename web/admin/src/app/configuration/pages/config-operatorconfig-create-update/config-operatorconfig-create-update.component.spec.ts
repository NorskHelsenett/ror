import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ConfigOperatorconfigCreateUpdateComponent } from './config-operatorconfig-create-update.component';

describe('ConfigOperatorconfigCreateUpdateComponent', () => {
  let component: ConfigOperatorconfigCreateUpdateComponent;
  let fixture: ComponentFixture<ConfigOperatorconfigCreateUpdateComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ConfigOperatorconfigCreateUpdateComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ConfigOperatorconfigCreateUpdateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
