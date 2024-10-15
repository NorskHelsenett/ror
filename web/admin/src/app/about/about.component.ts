import { ChangeDetectionStrategy, ChangeDetectorRef, Component, OnInit } from '@angular/core';
import { catchError, Observable, share, tap } from 'rxjs';
import { AboutService } from '../core/services/apihealthcheck.service';
import { environment } from '../../environments/environment';

@Component({
  selector: 'app-about',
  templateUrl: './about.component.html',
  styleUrls: ['./about.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AboutComponent implements OnInit {
  appVersion = environment.appVersion;

  healthError: any;
  health$: Observable<any> | undefined;

  constructor(
    private aboutService: AboutService,
    private changeDetector: ChangeDetectorRef,
  ) {}

  ngOnInit(): void {
    this.healthcheck();
  }

  healthcheck(): void {
    this.healthError = undefined;
    this.health$ = this.aboutService.getHealth().pipe(
      share(),
      tap(() => {
        this.changeDetector.detectChanges();
      }),
      catchError((error) => {
        this.healthError = error;
        this.changeDetector.detectChanges();
        return error;
      }),
    );
  }
}
