import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { SetupRoutes } from '../routes'
import jwt_decode from 'jwt-decode';
import { MySharedService } from '../shared/my-shared-service.service';


@Component({
  selector: 'app-sign-in',
  templateUrl: './sign-in.component.html',
  styleUrls: ['./sign-in.component.css']
})
export class SignInComponent {
  email: string = '';
  password: string = '';

  constructor(private router: Router, private http: HttpClient, private shared: MySharedService) {}

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
      const token = response.accessToken;
      const decodedToken: any = jwt_decode(token);

      const { role } = decodedToken;
      localStorage.setItem("role", role);

      this.shared.setRole(role);
      this.router.resetConfig(SetupRoutes());
      this.router.navigate(['/home']);

      console.log(response);
      console.log(decodedToken);
    }, error => {
      // Handle any errors that occurred during the request
      console.error(error);
    });
  }
}
