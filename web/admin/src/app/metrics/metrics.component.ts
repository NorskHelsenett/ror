import { ChangeDetectionStrategy, Component, OnInit, OnDestroy, ChangeDetectorRef } from '@angular/core';
import { Subscription, Observable, tap, catchError, interval } from 'rxjs';
import { MetricsCustom, MetricsCustomItem } from '../core/models/metricsCustom';
import { ChartOptions } from 'chart.js';
import { Dialog } from 'primeng/dialog';
import { MetricsService } from '../core/services/metrics.service';
import { ThemeService } from '../core/services/theme.service';

@Component({
  selector: 'app-metrics',
  templateUrl: './metrics.component.html',
  styleUrls: ['./metrics.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class MetricsComponent implements OnInit, OnDestroy {
  chartOptions: ChartOptions = {
    plugins: {
      legend: {
        display: false,
      },
    },
    animation: false,
    responsive: false,
  };
  chartData: any | undefined;
  isDark = false;

  backgroundColors = ['#00467A', '#372770', '#6B1E27', '#E85800'];
  lightbackgroundColors = ['#90DDFA', '#C0A9FF', '#D48282', '#FFC46B'];

  customMetrics$: Observable<any> | undefined;
  data$: Observable<any> | undefined;
  metricsError: any;
  customMetricsData: MetricsCustom | undefined;

  display: boolean = false;

  private subscriptions = new Subscription();

  constructor(
    private changeDetector: ChangeDetectorRef,
    private themeService: ThemeService,
    private metricsService: MetricsService,
  ) {}

  ngOnInit(): void {
    this.subscriptions.add(
      this.themeService.isDark.subscribe((value) => {
        this.isDark = value;
        this.changeDetector.detectChanges();
      }),
    );
    this.fetchData();
    this.fetchCustomMetrics();
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }

  showDialog(dialog: Dialog) {
    if (dialog) {
      dialog?.maximize();
    }
    this.display = true;
  }

  fetchCustomMetrics(): void {
    this.subscriptions.add(
      interval(2000)
        .pipe(
          tap(() => {
            this.fetchData();
          }),
        )
        .subscribe(),
    );
  }

  fetchData(): void {
    this.subscriptions.add(
      this.metricsService
        .getForClusterByProperty('versions.nhntooling.version')
        .pipe(
          tap((data: MetricsCustom) => {
            this.customMetricsData = data;
            this.setChartData();
            this.changeDetector.detectChanges();
          }),
          catchError((error) => {
            this.metricsError = error;
            return error;
          }),
        )
        .subscribe(),
    );
  }

  setChartData(): void {
    let labels: string[] = [];
    let data: number[] = [];

    this.customMetricsData?.data?.forEach((element: MetricsCustomItem) => {
      labels.push(element?.text);
      data.push(element?.value);
    });

    this.chartData = {
      responsive: true,
      labels: labels,
      datasets: [
        {
          data: data,
          backgroundColor: this.isDark ? this.backgroundColors : this.lightbackgroundColors,
          hoverBackgroundColor: this.isDark ? this.lightbackgroundColors : this.backgroundColors,
        },
      ],
    };
  }
}
