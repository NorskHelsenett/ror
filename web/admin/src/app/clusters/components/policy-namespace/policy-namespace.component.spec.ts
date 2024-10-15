import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PolicyNamespaceComponent } from './policy-namespace.component';

describe('PolicyNamespaceComponent', () => {
  let component: PolicyNamespaceComponent;
  let fixture: ComponentFixture<PolicyNamespaceComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [PolicyNamespaceComponent],
    });
    fixture = TestBed.createComponent(PolicyNamespaceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
