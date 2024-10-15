import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateClusterComponent } from './create-cluster.component';

describe('CreateClusterComponent', () => {
  let component: CreateClusterComponent;
  let fixture: ComponentFixture<CreateClusterComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [CreateClusterComponent],
    });
    fixture = TestBed.createComponent(CreateClusterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
