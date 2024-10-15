import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-price-list',
  templateUrl: './price-list.component.html',
  styleUrls: ['./price-list.component.scss'],
})
export class PriceListComponent implements OnInit {
  @Input() prices: any[];

  constructor() {}

  ngOnInit(): void {
    return;
  }
}
