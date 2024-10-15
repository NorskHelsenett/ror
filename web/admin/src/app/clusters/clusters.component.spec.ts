import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClustersComponent } from './clusters.component';

describe('ClustersComponent', () => {
  let component: ClustersComponent;
  let fixture: ComponentFixture<ClustersComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ClustersComponent],
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ClustersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
