import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IngressDetailsComponent } from './ingress-details.component';

describe('IngressDetailsComponent', () => {
  let component: IngressDetailsComponent;
  let fixture: ComponentFixture<IngressDetailsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [IngressDetailsComponent],
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(IngressDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
