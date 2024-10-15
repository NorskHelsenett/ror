import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Workspace } from '../models/workspace';
import { ConfigService } from './config.service';

@Injectable({
  providedIn: 'root',
})
export class WorkspacesService {
  constructor(
    private httpClient: HttpClient,
    private configService: ConfigService,
  ) {}

  get(): Observable<Workspace[]> {
    const url = `${this.configService.config.rorApi}/v1/workspaces`;
    return this.httpClient.get<Workspace[]>(url);
  }

  getByName(name: string): Observable<Workspace> {
    const url = `${this.configService.config.rorApi}/v1/workspaces/${name}`;
    return this.httpClient.get<Workspace>(url);
  }

  getById(id: string): Observable<Workspace> {
    const url = `${this.configService.config.rorApi}/v1/workspaces/id/${id}`;
    return this.httpClient.get<Workspace>(url);
  }

  update(input: Workspace): Observable<Workspace> {
    const url = `${this.configService.config.rorApi}/v1/workspaces/${input.id}`;
    return this.httpClient.put<Workspace>(url, input);
  }
}
