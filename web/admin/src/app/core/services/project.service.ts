import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Filter } from '../models/apiFilter';
import { ClusterInfo } from '../models/clusterInfo';
import { PaginationResult } from '../models/paginatedResult';
import { Project } from '../models/project';
import { ConfigService } from './config.service';

@Injectable({
  providedIn: 'root',
})
export class ProjectService {
  constructor(
    private httpClient: HttpClient,
    private configService: ConfigService,
  ) {}

  getByFilter(filter: Filter): Observable<PaginationResult<Project>> {
    let url = `${this.configService.config.rorApi}/v1/projects/filter`;
    return this.httpClient.post<PaginationResult<Project>>(url, filter);
  }

  getById(id: string): Observable<Project> {
    let url = `${this.configService.config.rorApi}/v1/projects/${id}`;
    return this.httpClient.get<Project>(url);
  }

  create(project: Project): Observable<Project> {
    let url = `${this.configService.config.rorApi}/v1/projects`;
    return this.httpClient.post<Project>(url, project);
  }

  update(project: Project): Observable<Project> {
    let url = `${this.configService.config.rorApi}/v1/projects/${project.id}`;
    return this.httpClient.put<Project>(url, project);
  }

  delete(project: Project): Observable<Project> {
    let url = `${this.configService.config.rorApi}/v1/projects/${project.id}`;
    return this.httpClient.delete<Project>(url);
  }

  clustersByProjectId(projectId: string): Observable<ClusterInfo[]> {
    let url = `${this.configService.config.rorApi}/v1/projects/${projectId}/clusters`;
    return this.httpClient.get<ClusterInfo[]>(url);
  }
}
