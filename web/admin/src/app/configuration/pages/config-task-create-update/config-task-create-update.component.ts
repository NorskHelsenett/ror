import { ChangeDetectionStrategy, ChangeDetectorRef, Component, OnDestroy, OnInit } from '@angular/core';
import { FormBuilder } from '@angular/forms';
import { ActivatedRoute } from '@angular/router';
import { TranslateService, LangChangeEvent } from '@ngx-translate/core';
import { catchError, Subscription, tap } from 'rxjs';
import { TasksService } from '../../../core/services/tasks.service';
import { Task } from '../../../core/models/task';

@Component({
  selector: 'app-config-task-create-update',
  templateUrl: './config-task-create-update.component.html',
  styleUrls: ['./config-task-create-update.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ConfigTaskCreateUpdateComponent implements OnInit, OnDestroy {
  id: string;
  task: Task | undefined;
  taskFetchError: any;

  constructor(
    private changeDetector: ChangeDetectorRef,
    private fb: FormBuilder,
    private route: ActivatedRoute,
    private translateService: TranslateService,
    private taskService: TasksService,
  ) {}

  private subscriptions = new Subscription();

  ngOnInit(): void {
    this.setupForm();

    this.subscriptions.add(
      this.route.params.subscribe((param) => {
        this.id = param?.['id'];
        if (this.id !== '' && this.id !== null && this.id !== undefined) {
          this.fetch();
          this.changeDetector.detectChanges();
        }
      }),
    );

    this.subscriptions.add(
      this.translateService.onLangChange
        .pipe(
          tap((event: LangChangeEvent) => {
            event.lang;
          }),
        )
        .subscribe(),
    );

    this.changeDetector.detectChanges();
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }

  fetch(): void {
    this.taskFetchError = undefined;
    this.subscriptions.add(
      this.taskService
        .getById(this.id)
        .pipe(
          catchError((error: any) => {
            this.taskFetchError = error;
            this.changeDetector.detectChanges();
            throw error;
          }),
          tap((task: Task) => {
            this.task = task;
            this.fillForm();
            this.changeDetector.detectChanges();
          }),
        )
        .subscribe(),
    );
  }

  private setupForm(): void {}

  private fillForm(): void {}
}
