import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Component } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import jwt_decode from 'jwt-decode';

@Component({
  selector: 'app-home-collector',
  templateUrl: './home-collector.component.html',
  styleUrls: ['./home-collector.component.css'],
})
export class HomeCollectorComponent {
  items: any[] = [];
  accessToken: string = '';
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
      const decode: any = jwt_decode(token);
      const id = decode.id;
      const url = `http://localhost:5000/items/collectors/${id}`;
      const headers = new HttpHeaders({
        Authorization: `Bearer ${token}`,
      });
      
      this.http.get(url, { headers }).subscribe(
        (response: any) => {
          this.items = response;
          console.log(this.items);
        },
        (error) => {
          console.error(error);
        }
      );
    }
  }

  upload() {
    this.router.navigate(['/upload']);
  }
}
