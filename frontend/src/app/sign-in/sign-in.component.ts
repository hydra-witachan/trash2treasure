import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-sign-in',
  templateUrl: './sign-in.component.html',
  styleUrls: ['./sign-in.component.css']
})
export class SignInComponent {
  email: string = '';
  password: string = '';

  constructor(private router: Router, private http: HttpClient) {}

  submit() {
    console.log('Email:', this.email);
    console.log('Password:', this.password);

    const url = 'http://localhost:5000/users/login';
    const body = {
      email: this.email,
      password: this.password
    };

    this.http.post<any>(url, body)
    .subscribe(response => {
      // Handle the response from the server
      console.log(response);
    }, error => {
      // Handle any errors that occurred during the request
      console.error(error);
    });
  }
}
