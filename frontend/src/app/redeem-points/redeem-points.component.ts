import { DecimalPipe } from '@angular/common';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Component } from '@angular/core';
import { Router } from '@angular/router';
import jwt_decode from 'jwt-decode';

@Component({
  selector: 'app-redeem-points',
  templateUrl: './redeem-points.component.html',
  styleUrls: ['./redeem-points.component.css'],
})
export class RedeemPointsComponent {
  userPoints = 0;
  selectedRedeemIdx = -1;
  formattedUserPoints = '';
  accessToken: string | null = '';

  redeemMap: PointsExchange[] = [
    {
      pointsToExchange: 20_000,
      idrMoney: 10_000,
    },
    {
      pointsToExchange: 50_000,
      idrMoney: 30_000,
    },
    {
      pointsToExchange: 100_000,
      idrMoney: 70_000,
    },
    {
      pointsToExchange: 200_000,
      idrMoney: 130_000,
    },
  ];

  readonly selectedColor = 'blue';

  constructor(private decimalPipe: DecimalPipe, private http: HttpClient, private router: Router) {
    this.userPoints = 0;
    this.formattedUserPoints =
      this.decimalPipe
        .transform(this.userPoints, '1.0-3')
        ?.replaceAll(',', '.') ?? '0';
    const token = localStorage.getItem('accessToken');
    if(token) {
      this.accessToken = localStorage.getItem('accessToken');
    }
  }

  ngOnInit() {
    const token = this.accessToken;
    let decode: any;
    if (token) {
      decode = jwt_decode(token);
      const { id } = decode;

      const url = `http://localhost:5000/users/${id}`;
      const headers = new HttpHeaders({
        Authorization: `Bearer ${token}`,
      });

      // console.log(url);
      // console.log(headers);

      this.http.get(url, { headers }).subscribe(
        (response: any) => {
          const { points } = response;
          this.formattedUserPoints = this.formatNumber(points);
          this.userPoints = points;
          // console.log(response);
          // console.log(this.userPoints);
        },
        (error) => {
          // Handle any errors that occurred during the request
          console.error(error);
        }
      );
    }
  }

  formatNumber(num: number): string {
    return (
      this.decimalPipe.transform(num, '1.0-3')?.replaceAll(',', '.') ?? '0'
    );
  }

  onRedeemSelect(idx: number) {
    if (this.selectedRedeemIdx === idx) {
      this.selectedRedeemIdx = -1;
    }
    this.selectedRedeemIdx = idx;
  }

  submit() {
    // console.log(this);
    if(this.selectedRedeemIdx !== -1) { // Didn't choice
      const token = this.accessToken;
      if(token) {
        const url = "http://localhost:5000/users/redeem";
        const body = {
          pointsToExchange: this.redeemMap[this.selectedRedeemIdx].pointsToExchange,
        };
        const headers = new HttpHeaders({
          Authorization: `Bearer ${token}`,
        });

        // console.log(url);
        // console.log(body);

        this.http.post<any>(url, body, { headers })
        .subscribe((response: any) => {
          this.userPoints -= this.redeemMap[this.selectedRedeemIdx].pointsToExchange;
          this.formattedUserPoints = this.formatNumber(this.userPoints);
          this.router.navigate(['/redeem/success']);
        }, error => {
          // Handle any errors that occurred during the request
          console.error(error);
        });
      }
    }
  }
}

type PointsExchange = {
  pointsToExchange: number;
  idrMoney: number;
};
