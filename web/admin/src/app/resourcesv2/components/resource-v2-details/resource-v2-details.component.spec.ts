import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ResourceV2DetailsComponent } from './resource-v2-details.component';

describe('ResourceV2DetailsComponent', () => {
  let component: ResourceV2DetailsComponent;
  let fixture: ComponentFixture<ResourceV2DetailsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ResourceV2DetailsComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ResourceV2DetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
