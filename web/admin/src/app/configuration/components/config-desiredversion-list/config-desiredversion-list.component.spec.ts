import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ConfigDesiredversionListComponent } from './config-desiredversion-list.component';

describe('ConfigDesiredversionListComponent', () => {
  let component: ConfigDesiredversionListComponent;
  let fixture: ComponentFixture<ConfigDesiredversionListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ConfigDesiredversionListComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ConfigDesiredversionListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
