import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Task } from '../models/task';
import { ConfigService } from './config.service';

@Injectable({
  providedIn: 'root',
})
export class TasksService {
  constructor(
    private httpClient: HttpClient,
    private configService: ConfigService,
  ) {}

  getAll(): Observable<Task[]> {
    const url = `${this.configService.config.rorApi}/v1/tasks`;
    return this.httpClient.get<Task[]>(url);
  }

  getById(id: string): Observable<Task> {
    const url = `${this.configService.config.rorApi}/v1/tasks/${id}`;
    return this.httpClient.get<Task>(url);
  }
}
