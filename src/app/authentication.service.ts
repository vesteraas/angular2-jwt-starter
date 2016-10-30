import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';
import { tokenNotExpired } from 'angular2-jwt';
import { Observable } from 'rxjs';
import 'rxjs/add/operator/map'

@Injectable()
export class AuthenticationService {
  public token: string;

  constructor(private http: Http) { }

  register(email, password, firstName, lastName): Observable<boolean> {
    return this.http.post('/api/register', JSON.stringify({ email: email, password: password, firstName: firstName, lastName: lastName }))
      .map((response: Response) => {
        let token = response.json() && response.json().token;
        if (token) {
          this.token = token;

          localStorage.setItem('id_token', token);

          return true;
        } else {
          return false;
        }
      }).catch((error: any) => {
        return Observable.throw(error.json());
      });
  }

  login(email, password): Observable<boolean> {
    return this.http.post('/api/authenticate', JSON.stringify({ email: email, password: password }))
      .map((response: Response) => {
        let token = response.json() && response.json().token;
        if (token) {
          this.token = token;

          localStorage.setItem('id_token', token);

          return true;
        } else {
          return false;
        }
      }).catch((error: any) => {
        return Observable.throw(error.json());
      });
  }

  logout() {
    localStorage.removeItem('id_token');
  }

  loggedIn() {
    return tokenNotExpired();
  }
}
