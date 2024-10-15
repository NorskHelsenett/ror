import { catchError, finalize, Observable, share } from 'rxjs';
import { ChangeDetectionStrategy, Component, OnInit, ChangeDetectorRef } from '@angular/core';
import { PriceService } from '../core/services/price.service';
import { ConfigService } from '../core/services/config.service';

@Component({
  selector: 'app-prices',
  templateUrl: './prices.component.html',
  styleUrls: ['./prices.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class PricesComponent implements OnInit {
  prices$: Observable<any>;
  pricesError: any;

  rows = this.configService.config.rows;
  rowsPerPage = this.configService.config.rowsPerPage;
  loading: boolean;

  constructor(
    private changeDetector: ChangeDetectorRef,
    private priceService: PriceService,
    private configService: ConfigService,
  ) {}

  ngOnInit(): void {
    this.fetch();
  }

  fetch(): void {
    this.loading = true;
    this.pricesError = undefined;
    this.prices$ = this.priceService.getAll().pipe(
      share(),
      catchError((error) => {
        this.pricesError = error;
        return error;
      }),
      finalize(() => {
        this.loading = false;
        this.changeDetector.detectChanges();
      }),
    );
  }
}
