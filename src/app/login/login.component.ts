import { Component, OnInit } from '@angular/core';
import { AuthenticationService } from '../authentication.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  title = 'Login'
  model: any = {};
  loading = false;
  error = '';

  newUser = false;

  constructor(private router: Router,
              private authenticationService: AuthenticationService) {
  }

  ngOnInit() {
    this.authenticationService.logout();
  }

  login() {
    this.loading = true;
    this.authenticationService.login(this.model.email, this.model.password)
      .subscribe(result =>  {
          if (result === true) {
          this.router.navigate(['/']);
        } else {
          this.error = 'Email or password is incorrect';
          this.loading = false;
        }
      }, error => {
        if (error.message) {
          this.loading = false;
          this.error = error.message;
        }
      });
  }

  showRegister() {
    this.error = null;
    this.newUser = true;
    this.title = 'Register new user';
  }

  register() {
    this.loading = true;
    this.authenticationService.register(this.model.email, this.model.password, this.model.firstName, this.model.lastName)
      .subscribe(result => {
        if (result === true) {
          this.router.navigate(['/']);
        }
      }, error => {
        if (error.message) {
          this.loading = false;
          this.error = error.message;
        }
      });
  }
}
