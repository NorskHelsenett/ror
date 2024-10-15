import { Observable } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Datacenter } from '../models/datacenter';
import { ConfigService } from './config.service';

@Injectable({
  providedIn: 'root',
})
export class DatacenterService {
  constructor(
    private httpClient: HttpClient,
    private configService: ConfigService,
  ) {}

  get(): Observable<Array<Datacenter>> {
    const url = `${this.configService.config.rorApi}/v1/datacenters`;
    return this.httpClient.get<Array<Datacenter>>(url);
  }

  getByName(name: string): Observable<Datacenter> {
    const url = `${this.configService.config.rorApi}/v1/datacenters/${name}`;
    return this.httpClient.get<Datacenter>(url);
  }

  getById(id: string): Observable<Datacenter> {
    const url = `${this.configService.config.rorApi}/v1/datacenters/id/${id}`;
    return this.httpClient.get<Datacenter>(url);
  }

  create(datacenter: Datacenter): Observable<Datacenter> {
    const url = `${this.configService.config.rorApi}/v1/datacenters`;
    return this.httpClient.post<Datacenter>(url, datacenter);
  }

  update(datacenter: Datacenter): Observable<Datacenter> {
    const url = `${this.configService.config.rorApi}/v1/datacenters/${datacenter.id}`;
    return this.httpClient.put<Datacenter>(url, datacenter);
  }
}
