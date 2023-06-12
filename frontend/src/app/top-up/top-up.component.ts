import { Component } from '@angular/core';
import { DecimalPipe } from '@angular/common';

@Component({
  selector: 'app-top-up',
  templateUrl: './top-up.component.html',
  styleUrls: ['./top-up.component.css']
})
export class TopUpComponent {
  point = 50000;
  formattedPoint: string | null;
  topUp = [
    {topUpPoint: 10000, feeTopUp: 10000},
    {topUpPoint: 25000, feeTopUp: 25000},
    {topUpPoint: 50000, feeTopUp: 50000},
    {topUpPoint: 100000, feeTopUp: 100000},
    {topUpPoint: 150000, feeTopUp: 150000},
    {topUpPoint: 200000, feeTopUp: 200000},
    {topUpPoint: 300000, feeTopUp: 300000},
    {topUpPoint: 400000, feeTopUp: 400000},
    {topUpPoint: 500000, feeTopUp: 500000},
  ]

  constructor(private decimalPipe: DecimalPipe) {
    this.formattedPoint = this.decimalPipe.transform(this.point, '1.0-3');
    if(this.formattedPoint !== null) {
      this.formattedPoint =  this.formattedPoint.replaceAll(',', '.')
    }
  }

  updatePoint(value: number) {
    console.log(value);
    this.point = value;
    this.formattedPoint = this.decimalPipe.transform(this.point, '1.0-3');
    if(this.formattedPoint !== null) {
      this.formattedPoint =  this.formattedPoint.replaceAll(',', '.')
    }
  }
}