import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ConfigDesiredversionCreateUpdateComponent } from './config-desiredversion-create-update.component';

describe('ConfigDesiredversionCreateUpdateComponent', () => {
  let component: ConfigDesiredversionCreateUpdateComponent;
  let fixture: ComponentFixture<ConfigDesiredversionCreateUpdateComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ConfigDesiredversionCreateUpdateComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ConfigDesiredversionCreateUpdateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
