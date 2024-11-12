import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterNewComponent } from './cluster-new.component';

describe('ClusterNewComponent', () => {
  let component: ClusterNewComponent;
  let fixture: ComponentFixture<ClusterNewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ClusterNewComponent],
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ClusterNewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
