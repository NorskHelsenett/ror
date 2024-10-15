import { Injectable, NgZone } from '@angular/core';
import { OAuthService } from 'angular-oauth2-oidc';
import { EventSourcePolyfill } from 'event-source-polyfill';
import { Observable, Subscriber } from 'rxjs';
import { ConfigService } from '../../../core/services/config.service';

@Injectable({
  providedIn: 'root',
})
export class SseService {
  isConnected = false;

  private events: Observable<any> | undefined;
  private eventSource: EventSourcePolyfill | undefined;
  private timeout = 5000;
  private error: any;
  private observer: Subscriber<any> | undefined;
  private reconnectFrecuencySeconds = 1;

  private url = `${this.configService?.config?.rorApi}${this.configService?.config?.sse?.postfixUrl}?ngsw-bypass=true`;

  constructor(
    private ngZone: NgZone,
    private configService: ConfigService,
    private oauthService: OAuthService,
  ) {}

  getEvents(): Observable<any> {
    return new Observable((observer) => {
      this.observer = observer;
      this.getEventSource(this.url);
      this.eventSource.onopen = (event: any) => {
        this.ngZone.runGuarded(() => {
          this.error = undefined;
          this.isConnected = true;
        });
      };
      this.eventSource.onerror = (error: any) => {
        console.log('error', error);
        this.ngZone.runGuarded(() => {
          this.error = error;
          this.isConnected = false;
          if (this.eventSource?.readyState === 2) {
            this.close();
            this.reconnect();
          }
        });
      };
      this.eventSource.onmessage = (event: any) => {
        console.log('event', event);
        this.ngZone.runGuarded(() => {
          this.observer.next(event);
        });
      };
    });
  }

  private getEventSource(url: string): void {
    if (this.eventSource) {
      this.eventSource.close();
    }
    this.eventSource = new EventSourcePolyfill(url, {
      headers: {
        Authorization: this.oauthService.authorizationHeader(),
      },
      heartbeatTimeout: this.configService?.config?.sse?.timeout,
    });

    return this.eventSource;
  }

  close(): void {
    this.eventSource?.close();
    this.eventSource?.close();
    this.eventSource = undefined;
    this.events = undefined;
    this.observer = undefined;
    this.isConnected = false;
  }

  private reconnect(): void {
    setTimeout(() => {
      this.tryToReconnect();
    }, this.timeout);
  }

  private tryToReconnect(): void {
    this.getEvents();
    this.reconnectFrecuencySeconds *= 2;
    if (this.reconnectFrecuencySeconds >= 64) {
      this.reconnectFrecuencySeconds = 64;
    }
  }
}
