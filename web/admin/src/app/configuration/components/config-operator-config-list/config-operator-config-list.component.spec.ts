import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ConfigOperatorConfigListComponent } from './config-operator-config-list.component';

describe('ConfigOperatorConfigListComponent', () => {
  let component: ConfigOperatorConfigListComponent;
  let fixture: ComponentFixture<ConfigOperatorConfigListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ConfigOperatorConfigListComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ConfigOperatorConfigListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
