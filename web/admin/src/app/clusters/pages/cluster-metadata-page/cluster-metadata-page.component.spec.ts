import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClusterMetadataPageComponent } from './cluster-metadata-page.component';

describe('ClusterMetadataComponent', () => {
  let component: ClusterMetadataPageComponent;
  let fixture: ComponentFixture<ClusterMetadataPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ClusterMetadataPageComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ClusterMetadataPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
