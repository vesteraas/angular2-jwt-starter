import { Injectable } from '@angular/core';
import { AuthHttp } from 'angular2-jwt';
import { Response } from '@angular/http';
import { Observable } from 'rxjs';
import 'rxjs/add/operator/map'
import { User } from './models/user';

@Injectable()
export class UserService {

  constructor( private http: AuthHttp) {
  }

  getUsers(): Observable<User[]> {
    return this.http.get('/api/users')
      .map((response: Response) => response.json());
  }
}
