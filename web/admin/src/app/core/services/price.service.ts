import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Price } from '../models/price';
import { ConfigService } from './config.service';

@Injectable({
  providedIn: 'root',
})
export class PriceService {
  constructor(
    private httpClient: HttpClient,
    private configService: ConfigService,
  ) {}

  getAll(): Observable<Price[]> {
    let url = `${this.configService.config.rorApi}/v1/prices`;
    return this.httpClient.get<Price[]>(url);
  }

  getById(id: string): Observable<Price> {
    let url = `${this.configService.config.rorApi}/v1/prices/${id}`;
    return this.httpClient.get<Price>(url);
  }

  create(price: Price): Observable<Price> {
    let url = `${this.configService.config.rorApi}/v1/prices`;
    return this.httpClient.post<Price>(url, price);
  }

  update(price: Price): Observable<Price> {
    let url = `${this.configService.config.rorApi}/v1/prices/${price.id}`;
    return this.httpClient.put<Price>(url, price);
  }
}
