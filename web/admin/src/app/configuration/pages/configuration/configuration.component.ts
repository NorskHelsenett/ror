import { ChangeDetectionStrategy, ChangeDetectorRef, Component, OnDestroy, OnInit } from '@angular/core';
import { catchError, Observable, share, Subscription } from 'rxjs';
import { DesiredVersion } from '../../../core/models/desiredversion';
import { OperatorConfig } from '../../../core/models/operatorconfig';
import { User } from '../../../core/models/user';
import { DesiredversionsService } from '../../../core/services/desiredversions.service';
import { OperatorConfigsService } from '../../../core/services/operator-configs.service';
import { TasksService } from '../../../core/services/tasks.service';
import { UserService } from '../../../core/services/user.service';
import { Task } from '../../../core/models/task';

@Component({
  selector: 'app-configuration',
  templateUrl: './configuration.component.html',
  styleUrls: ['./configuration.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ConfigurationComponent implements OnInit, OnDestroy {
  user$: Observable<User> | undefined;

  operatorConfigs$: Observable<OperatorConfig[]> | undefined;
  operatorConfigFetchError: any;

  tasks$: Observable<Task[]> | undefined;
  tasksFetchError: any;

  desiredVersions$: Observable<DesiredVersion[]> | undefined;
  desiredVersionsFetchError: any;

  private subscriptions: Subscription = new Subscription();

  constructor(
    private changeDetector: ChangeDetectorRef,
    private userService: UserService,
    private taskService: TasksService,
    private desiredVersionService: DesiredversionsService,
    private operatorConfigService: OperatorConfigsService,
  ) {}

  ngOnInit(): void {
    this.setupObservables();
    this.fetchData();
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }

  setupObservables(): void {
    this.user$ = this.userService?.getUser();
    if (!this.userService.user.value) {
      this.subscriptions.add(
        this.userService
          .getUser()
          .pipe(share())
          .subscribe((_) => {
            this.changeDetector.detectChanges();
          }),
      );
    }
    this.fetchOperatorConfigs();
    this.fetchTasks();
    this.fetchDesiredVersions();
  }

  fetchOperatorConfigs(): void {
    this.operatorConfigFetchError = undefined;
    this.operatorConfigs$ = this.operatorConfigService.getAll().pipe(
      catchError((error: any) => {
        this.operatorConfigFetchError = error;
        this.changeDetector.detectChanges();
        throw error;
      }),
    );
  }

  fetchTasks(): void {
    this.tasksFetchError = undefined;
    this.tasks$ = this.taskService.getAll().pipe(
      catchError((error: any) => {
        this.tasksFetchError = error;
        this.changeDetector.detectChanges();
        throw error;
      }),
    );
  }

  fetchDesiredVersions(): void {
    this.desiredVersionsFetchError = undefined;
    this.desiredVersions$ = this.desiredVersionService.getAll().pipe(
      catchError((error: any) => {
        this.desiredVersionsFetchError = error;
        this.changeDetector.detectChanges();
        throw error;
      }),
    );
  }

  fetchData(): void {
    this.fetchOperatorConfigs();
    this.fetchTasks();
    this.fetchDesiredVersions();
  }
}
