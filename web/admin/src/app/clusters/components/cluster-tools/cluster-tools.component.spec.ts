import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterToolsComponent } from './cluster-tools.component';

describe('ClusterToolsComponent', () => {
  let component: ClusterToolsComponent;
  let fixture: ComponentFixture<ClusterToolsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ClusterToolsComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ClusterToolsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
