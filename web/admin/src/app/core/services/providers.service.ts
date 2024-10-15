import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { ConfigService } from './config.service';
import { Observable } from 'rxjs';
import { Provider, ProviderKubernetesVersion } from '../models/provider';
import { ClusterProvider } from '../../clusters/models/clusterProvider';

@Injectable({
  providedIn: 'root',
})
export class ProvidersService {
  constructor(
    private httpClient: HttpClient,
    private configService: ConfigService,
  ) {}

  get(): Observable<Array<Provider>> {
    const url = `${this.configService.config.rorApi}/v1/providers`;
    return this.httpClient.get<Array<Provider>>(url);
  }

  getKubernetesVersionByProviderType(type: ClusterProvider): Observable<Array<ProviderKubernetesVersion>> {
    const url = `${this.configService.config.rorApi}/v1/providers/${type}/kubernetes/versions`;
    return this.httpClient.get<Array<ProviderKubernetesVersion>>(url);
  }
}
