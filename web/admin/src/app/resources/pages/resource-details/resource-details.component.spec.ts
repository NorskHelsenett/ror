import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ResourceDetailsComponent } from './resource-details.component';

describe('ResourceDetailsComponent', () => {
  let component: ResourceDetailsComponent;
  let fixture: ComponentFixture<ResourceDetailsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ResourceDetailsComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ResourceDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
