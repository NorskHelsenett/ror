import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HealthCheck } from '../models/healthcheck';
import { ConfigService } from './config.service';

@Injectable({
  providedIn: 'root',
})
export class AboutService {
  constructor(
    private httpclient: HttpClient,
    private configService: ConfigService,
  ) {}

  getHealth(): Observable<HealthCheck> {
    const url = `${this.configService.config.rorApi}/health`;
    return this.httpclient.get<HealthCheck>(url);
  }
}
