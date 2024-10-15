import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProjectsCreateComponent } from './projects-create.component';

describe('ProjectsCreateComponent', () => {
  let component: ProjectsCreateComponent;
  let fixture: ComponentFixture<ProjectsCreateComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ProjectsCreateComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(ProjectsCreateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
