import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterRawComponent } from './cluster-raw.component';

describe('ClusterRawComponent', () => {
  let component: ClusterRawComponent;
  let fixture: ComponentFixture<ClusterRawComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ClusterRawComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ClusterRawComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
