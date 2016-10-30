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
    console.log(email, password);
    return this.http.post('/api/register', JSON.stringify({ email: email, password: password, firstName: firstName, lastName: lastName }))
      .map((response: Response) => {
        console.log(response);
        // login successful if there's a jwt token in the response
        let token = response.json() && response.json().token;
        if (token) {
          this.token = token;

          // store username and jwt token in local storage to keep user logged in between page refreshes
          localStorage.setItem('id_token', token);

          // return true to indicate successful login
          return true;
        } else {
          // return false to indicate failed login
          return false;
        }
      }, (err: any) => {
        console.log(err);
      });
  }

  login(email, password): Observable<boolean> {
    console.log(email, password);
    return this.http.post('/api/authenticate', JSON.stringify({ email: email, password: password }))
      .map((response: Response) => {
        console.log(response);
        // login successful if there's a jwt token in the response
        let token = response.json() && response.json().token;
        if (token) {
          this.token = token;

          // store username and jwt token in local storage to keep user logged in between page refreshes
          localStorage.setItem('id_token', token);

          // return true to indicate successful login
          return true;
        } else {
          // return false to indicate failed login
          return false;
        }
      }, (err: any) => {
        console.log(err);
      });
  }

  logout() {
    localStorage.removeItem('id_token');
  }

  loggedIn() {
    return tokenNotExpired();
  }
}
