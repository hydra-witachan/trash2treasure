import { Component } from '@angular/core';
import { MySharedService } from '../shared/my-shared-service.service';
import { DecimalPipe } from '@angular/common';
import { Router } from '@angular/router';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Component({
  selector: 'app-confirmation-top-up',
  templateUrl: './confirmation-top-up.component.html',
  styleUrls: ['./confirmation-top-up.component.css'],
})
export class ConfirmationTopUpComponent {
  topUpPoint: any;
  formattedPoint: any;
  feeTopUp: any;
  formattedFeeTopUp: any;

  constructor(
    private sharedService: MySharedService,
    private decimalPipe: DecimalPipe,
    private router: Router,
    private http: HttpClient
  ) {}

  ngOnInit() {
    this.topUpPoint = this.sharedService.getPointTopUp();
    this.formattedPoint = this.sharedService.getPointTopUpFormated();
    this.feeTopUp = this.sharedService.getFeePointTopUp();
    this.formattedFeeTopUp = this.decimalPipe.transform(this.feeTopUp, '1.0-3');
    if (this.formattedFeeTopUp !== null) {
      this.formattedFeeTopUp = this.formattedFeeTopUp.replaceAll(',', '.');
    }
  }

  submit() {
    const token = localStorage.getItem('accessToken');
    const body = {
      fee: this.feeTopUp,
      points: this.topUpPoint,
      method: 'gopay',
    };
    const headers = new HttpHeaders({
      Authorization: `Bearer ${token}`,
    });
    const url = 'http://localhost:5000/users/topup';

    console.log(token);
    console.log(body);
    console.log(url);


    if(token) {
      this.http.patch<any>(url, body, { headers })
      .subscribe((response: any) => {
        this.router.navigate(['/top-up/confirmation/success']);
      }, error => {
        // Handle any errors that occurred during the request
        console.error(error);
      });
    }
  }
}
