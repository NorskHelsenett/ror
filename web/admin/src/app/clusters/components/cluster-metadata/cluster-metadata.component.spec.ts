import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterMetadataComponent } from './cluster-metadata.component';

describe('ClusterMetadataComponent', () => {
  let component: ClusterMetadataComponent;
  let fixture: ComponentFixture<ClusterMetadataComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ClusterMetadataComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ClusterMetadataComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
