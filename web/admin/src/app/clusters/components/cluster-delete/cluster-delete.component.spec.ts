import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterDeleteComponent } from './cluster-delete.component';

describe('ClusterDeleteComponent', () => {
  let component: ClusterDeleteComponent;
  let fixture: ComponentFixture<ClusterDeleteComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ClusterDeleteComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ClusterDeleteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
