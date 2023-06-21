import jwt_decode from 'jwt-decode';

import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Component } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { SetupRoutes } from '../routes';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent {

  accessToken!: string;
  items: any[] = [];
  user: UserProfile = {
    username: '',
    fullName: '',
    address: '',
    email: '',
    phoneNumber: '',
    points: 0,
    role: ''
  };

  constructor(
    private route: ActivatedRoute,
    private http: HttpClient,
    private router: Router
  ) {
    const token = localStorage.getItem('accessToken');
    if (token) {
      this.accessToken = token;
    }
  }

  ngOnInit() {
    const token = this.accessToken;
    if (token) {
      const decoded: any = jwt_decode(token);
      let targetUserId: string = decoded.id;

      if (!this.isCurrentUser()) {
        targetUserId = this.route.snapshot.paramMap.get('id')!
      }

      this.getCollectorItems(token, targetUserId);
      this.getProfile(token, targetUserId);
    }
  }

  getCollectorItems(token: string, targetUserId: string) {
    const url = `http://localhost:5000/items/collectors/${targetUserId}`;
    const headers = new HttpHeaders({
      Authorization: `Bearer ${token}`,
    });

    this.http.get(url, { headers }).subscribe(
      (response: any) => {
        this.items = response;
      },
      (error) => {
        console.error(error);
      }
    );
  }

  getProfile(token: string, targetUserId: string) {
    const url = `http://localhost:5000/users/${targetUserId}`;
    const headers = new HttpHeaders({
      Authorization: `Bearer ${token}`,
    });

    this.http.get(url, { headers }).subscribe(
      (response: any) => {
        this.user = response;
      },
      (error) => {
        console.error(error);
      }
    );
  }

  isCurrentUser() {
    return !this.route.snapshot.paramMap.has('id');
  }

  getCapitalizedRole() {
    if (!this.user.role) {
      return '';
    }
    return this.user.role[0].toUpperCase() + this.user.role.substring(1);
  }

  displayItemDetails(item: any) {
    this.router.navigate([`/items/${item.id}`]);
  }

  onLogout() {
    localStorage.removeItem('accessToken');
    localStorage.removeItem('role');

    this.router.resetConfig(SetupRoutes())
    this.router.navigate(['/auth/sign-in']);
  }

}

type UserProfile = {
  username: string,
  fullName: string,
  email: string,
  address: string,
  phoneNumber: string,
  role: string,
  points: number
}
