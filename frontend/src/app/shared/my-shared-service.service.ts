import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class MySharedService {
  private pointTopUp!: number;
  private feePointTopUp!: number;
  private pointTopUpFormated!: string;

  setPointTopUp(data: number) {
    this.pointTopUp = data;
  }

  setFeePointTopUp(data: number) {
    this.feePointTopUp = data;
  }

  setPointTopUpFormated(data: string) {
    this.pointTopUpFormated = data;
  }

  getPointTopUp(): number {
    return this.pointTopUp;
  }

  getFeePointTopUp(): number {
    return this.feePointTopUp;
  }

  getPointTopUpFormated(): string {
    return this.pointTopUpFormated;
  }
}
