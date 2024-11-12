import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ResourcesV2ListComponent } from './resources-v2-list.component';

describe('ResourcesV2ListComponent', () => {
  let component: ResourcesV2ListComponent;
  let fixture: ComponentFixture<ResourcesV2ListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ResourcesV2ListComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ResourcesV2ListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
