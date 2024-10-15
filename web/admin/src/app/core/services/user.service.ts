import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject, map, Observable } from 'rxjs';
import { User } from '../models/user';
import { ConfigService } from './config.service';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  user: BehaviorSubject<User> = new BehaviorSubject<User>(null);

  constructor(
    private httpClient: HttpClient,
    private configService: ConfigService,
  ) {}

  getUser(): Observable<User> {
    const url = `${this.configService.config.rorApi}/v1/users/self`;
    return this.httpClient.get<User>(url).pipe(
      map((user: User) => {
        if (!user) {
          this.user.next(null);
          return null;
        }

        this.user.next(user);
        return user;
      }),
    );
  }
}
