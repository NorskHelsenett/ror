import { Injectable, signal } from '@angular/core';
import { toObservable } from '@angular/core/rxjs-interop';
import { Observable } from 'rxjs';
import { v4 as uuidv4 } from 'uuid';
import { SseMessage } from '../../../core/models/sse';

@Injectable({
  providedIn: 'root',
})
export class SignalService {
  clusterCreated$: Observable<string> | undefined;
  private clusterCreated = signal<any>(null);

  clusterUpdated$: Observable<string> | undefined;
  private clusterUpdated = signal<any>(null);

  clusterOrderUpdated$: Observable<string> | undefined;
  private clusterOrderUpdated = signal<any>(null);

  constructor() {
    this.clusterCreated$ = toObservable(this.clusterCreated);
    this.clusterUpdated$ = toObservable(this.clusterCreated);
    this.clusterOrderUpdated$ = toObservable(this.clusterOrderUpdated);
  }

  handleEvent(event: any): void {
    if (!event) {
      return;
    }

    const messageString = event?.data;
    if (!messageString) {
      return;
    }

    const message = this.tryParseMessage(messageString);
    if (!message) {
      return;
    }

    switch (message?.event) {
      case 'cluster.created':
        this.clusterCreated.set(uuidv4());
        break;
      case 'cluster.updated':
        this.clusterUpdated.set(uuidv4());
        break;
      case 'clusterOrder.updated':
        this.clusterOrderUpdated.set(uuidv4());
        break;
      default:
        break;
    }
  }

  private tryParseMessage(messageString: string): SseMessage | undefined {
    try {
      return JSON.parse(messageString);
    } catch (error) {
      console.error('Error parsing message', error);
      return undefined;
    }
  }
}
