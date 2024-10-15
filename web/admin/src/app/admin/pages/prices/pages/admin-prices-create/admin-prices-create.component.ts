import { DatePipe } from '@angular/common';
import { ChangeDetectorRef, Component, OnDestroy, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { catchError, Subscription, tap } from 'rxjs';
import { Price } from '../../../../../core/models/price';
import { PriceService } from '../../../../../core/services/price.service';

@Component({
  selector: 'app-admin-prices-create',
  templateUrl: './admin-prices-create.component.html',
  styleUrls: ['./admin-prices-create.component.scss'],
})
export class AdminPricesCreateComponent implements OnInit, OnDestroy {
  priceForm: FormGroup | undefined;
  id: string;
  submitted = false;
  createError: any;
  updateError: any;
  priceCreateModel: Price | undefined;

  price: Price | undefined;
  priceFetchError: any;

  datePipe: DatePipe = new DatePipe('no');

  private subscriptions = new Subscription();

  constructor(
    private changeDetector: ChangeDetectorRef,
    private fb: FormBuilder,
    private priceService: PriceService,
    private router: Router,
    private route: ActivatedRoute,
  ) {}

  ngOnInit(): void {
    this.setupForm();

    this.subscriptions.add(
      this.route.params
        .pipe(
          tap((data: any) => {
            this.id = data?.id;
            if (this.id !== '' && this.id !== null && this.id !== undefined) {
              this.fetchPrice();
            }
          }),
        )
        .subscribe(),
    );

    this.changeDetector.detectChanges();
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }

  setupForm(): void {
    this.priceForm = this.fb.group({
      provider: ['', [Validators.required, Validators.minLength(1)]],
      machineClass: ['', [Validators.required, Validators.minLength(1)]],
      cpu: [, [Validators.required, Validators.min(1)]],
      memoryBytes: [Validators.required, Validators.min(1)],
      memory: [, [Validators.required, Validators.min(1)]],
      price: [, [Validators.required, Validators.min(1)]],
      from: [new Date(), [Validators.required]],
      to: [, []],
    });
  }

  fetchPrice(): void {
    this.subscriptions.add(
      this.priceService
        .getById(this.id)
        .pipe(
          tap((price: Price) => {
            this.price = price;
            if (price.from !== null && price.from !== undefined) {
              this.price.from = new Date(price.from);
            }
            if (price.to !== null && price.to !== undefined) {
              this.price.to = new Date(price.to);
            }
            this.fillForm();
          }),
          catchError((error) => {
            this.priceFetchError = error;
            this.changeDetector.detectChanges();
            throw error;
          }),
        )
        .subscribe(),
    );
  }

  fillForm(): void {
    this.priceForm.patchValue(this.price);
    this.changeDetector.detectChanges();
  }

  createPrice(): void {
    this.createError = undefined;
    this.priceCreateModel = this.priceForm.value as Price;
    if (!this.priceForm.valid) {
      this.priceForm.markAllAsTouched();
      this.changeDetector.detectChanges();
      return;
    }
    this.submitted = true;
    this.priceCreateModel = this.priceForm.value as Price;
    this.subscriptions.add(
      this.priceService
        .create(this.priceCreateModel)
        .pipe(
          tap((price: Price) => {
            if (!price) {
              this.createError = 'Could not create price';
            } else {
              this.router.navigate(['../'], { relativeTo: this.route });
            }
            this.changeDetector.detectChanges();
          }),
          catchError((error) => {
            this.createError = error;
            this.changeDetector.detectChanges();
            throw error;
          }),
        )
        .subscribe(),
    );
    this.changeDetector.detectChanges();
  }

  updatePrice(): void {
    this.createError = undefined;
    this.priceCreateModel = this.priceForm.value as Price;
    if (!this.priceForm.valid) {
      this.priceForm.markAllAsTouched();
      this.changeDetector.detectChanges();
      return;
    }
    this.submitted = true;
    this.priceCreateModel = this.priceForm.value as Price;
    this.priceCreateModel.id = this.price.id;
    this.subscriptions.add(
      this.priceService
        .update(this.priceCreateModel)
        .pipe(
          tap((price: Price) => {
            if (!price) {
              this.createError = 'Could not update price';
            } else {
              this.router.navigate(['../../'], { relativeTo: this.route });
            }
            this.changeDetector.detectChanges();
          }),
          catchError((error) => {
            this.createError = error;
            this.changeDetector.detectChanges();
            throw error;
          }),
        )
        .subscribe(),
    );
    this.changeDetector.detectChanges();
  }

  reset(): void {
    this.createError = undefined;
    this.priceCreateModel = null;
    this.submitted = false;
    this.priceForm.reset();
    this.changeDetector.detectChanges();
  }
}
