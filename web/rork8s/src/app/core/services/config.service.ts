import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
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
    return config;
  }

  private defaultConfig(): Config {
    const config: Config = {};

    if (environment?.production) {
    }

    return config;
  }
}
