import { DecimalPipe } from '@angular/common';
import { Component } from '@angular/core';

@Component({
  selector: 'app-redeem-points',
  templateUrl: './redeem-points.component.html',
  styleUrls: ['./redeem-points.component.css']
})
export class RedeemPointsComponent {

  userPoints = 0;
  selectedRedeemIdx = -1;
  formattedUserPoints = '';

  redeemMap: PointsExchange[] = [
    {
      pointsToExchange: 20_000,
      idrMoney: 10_000
    },
    {
      pointsToExchange: 50_000,
      idrMoney: 25_000
    },
    {
      pointsToExchange: 100_000,
      idrMoney: 55_000
    },
    {
      pointsToExchange: 200_000,
      idrMoney: 115_000
    }
  ]

  readonly selectedColor = 'blue';

  constructor(private decimalPipe: DecimalPipe) {
    this.formattedUserPoints = this.decimalPipe.transform(this.userPoints, '1.0-3')?.replaceAll(',', '.') ?? '0';
  }

  formatNumber(num: number): string {
    return this.decimalPipe.transform(num, '1.0-3')?.replaceAll(',', '.') ?? '0';
  }

  onRedeemSelect(idx: number) {
    if (this.selectedRedeemIdx === idx) {
      this.selectedRedeemIdx = -1
    }
    this.selectedRedeemIdx = idx;
  }

}

type PointsExchange = {
  pointsToExchange: number;
  idrMoney: number;
}
