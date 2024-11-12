import { ChangeDetectionStrategy, ChangeDetectorRef, Component, OnDestroy, OnInit } from '@angular/core';
import { Observable, Subscription, catchError, tap } from 'rxjs';
import { OrderService } from '../../../core/services/order.service';

@Component({
  selector: 'app-order-list',
  templateUrl: './order-list.component.html',
  styleUrl: './order-list.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class OrderListComponent implements OnInit, OnDestroy {
  orders$: Observable<any> | undefined;
  ordersFetchError: any;

  private subscriptions = new Subscription();

  constructor(
    private changeDetector: ChangeDetectorRef,
    private orderService: OrderService,
  ) {}

  ngOnInit(): void {
    this.fetchOrders();
  }

  fetchOrders(): void {
    this.ordersFetchError = undefined;
    this.orders$ = this.orderService.getOrders().pipe(
      tap(
        () => this.changeDetector.detectChanges(),
        catchError((error) => {
          this.ordersFetchError = error;
          this.changeDetector.detectChanges();
          return [];
        }),
      ),
    );
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }
}
