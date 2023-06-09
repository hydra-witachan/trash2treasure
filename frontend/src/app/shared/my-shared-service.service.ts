import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class MySharedService {
  private pointTopUp!: number;
  private feePointTopUp!: number;
  private pointTopUpFormated!: string;

  private role: string | null = 'guest';
  private subCategory!: string;

  constructor() {
    this.role = localStorage.getItem("role");
    if(!this.role) {
      this.role = 'guest';
    }
  }

  setRole(role: string) {
    this.role = role;
  }

  setPointTopUp(data: number) {
    this.pointTopUp = data;
  }

  setFeePointTopUp(data: number) {
    this.feePointTopUp = data;
  }

  setPointTopUpFormated(data: string) {
    this.pointTopUpFormated = data;
  }

  setSubCategory(subCategory: string) {
    this.subCategory = subCategory;
  }

  getRole(): string {
    return this.role!;
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

  getSubCategory(): string {
    return this.subCategory;
  }
}
