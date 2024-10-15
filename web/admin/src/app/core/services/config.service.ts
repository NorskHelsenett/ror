import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, catchError, map, of } from 'rxjs';
import { Config } from '../models/config';
import { environment } from '../../../environments/environment';

export const configFactory = (config: ConfigService) => {
  return () => config.loadConfig();
};

@Injectable({
  providedIn: 'root',
})
export class ConfigService {
  config: Config;

  constructor(private httpClient: HttpClient) {
    this.config = this.defaultConfig();
  }

  loadConfig(): Observable<boolean> {
    return this.httpClient.get(environment.configPath).pipe(
      map((response: any) => {
        // do something to reflect into local model
        this.config = this.createConfig(response);
        return true;
      }),
      catchError((error: any) => {
        this.config = this.defaultConfig();
        console.error('Error loading config', error);
        return of(false);
      }),
    );
  }

  private createConfig(json: any): Config {
    const config = { ...(<Config>json) };
    config.auth.redirectUri = window.location.origin + config.auth.redirectUri;
    config.auth.postLogoutRedirectUri = window.location.origin;
    return config;
  }

  private defaultConfig(): Config {
    const config: Config = {
      auth: {
        issuer: 'http://localhost:5556/dex',
        clientId: 'ror.sky.test.nhn.no',
        redirectUri: window.location.origin + '/auth/callback',
        scope: 'profile email groups',
        responseType: 'id_token token',
        logoutUrl: window.location.origin,
        postLogoutRedirectUri: window.location.origin,
        requireHttps: false,
        strictDiscoveryDocumentValidation: false,
      },
      regex: {
        forms: `^[@()\\/:?\\r\\n.,a-zA-Z æøåÆØÅ0-9_-]+$`,
      },
      rorApi: 'https://ror.sky.test.nhn.no',
      rows: 25,
      rowsPerPage: [10, 25, 50, 75, 100],
      sse: {
        postfixUrl: '/v1/events/listen',
        timeout: 30000,
        method: 'GET',
      },
    };

    if (environment.production) {
      config.auth.clientId = 'ror.sky.test.nhn.no';
      config.auth.issuer = 'https://auth.sky.nhn.no/dex';
      config.rorApi = 'https://ror.sky.nhn.no';
      config.auth.requireHttps = true;
      config.auth.strictDiscoveryDocumentValidation = true;
    }

    return config;
  }
}
