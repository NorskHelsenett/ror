import { ChangeDetectorRef, Component, OnInit } from '@angular/core';
import { catchError, finalize, Observable } from 'rxjs';
import { PriceService } from '../../../core/services/price.service';
import { ConfigService } from '../../../core/services/config.service';

@Component({
  selector: 'app-admin-prices',
  templateUrl: './admin-prices.component.html',
  styleUrls: ['./admin-prices.component.scss'],
})
export class AdminPricesComponent implements OnInit {
  prices$: Observable<any>;
  pricesError: any;
  loading: boolean;

  rows = this.configService.config.rows;
  rowsPerPage = this.configService.config.rowsPerPage;

  constructor(
    private changeDetector: ChangeDetectorRef,
    private pricesService: PriceService,
    private configService: ConfigService,
  ) {}

  ngOnInit(): void {
    this.fetch();
  }

  fetch(): void {
    this.loading = true;
    this.pricesError = undefined;
    this.prices$ = this.pricesService.getAll().pipe(
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
