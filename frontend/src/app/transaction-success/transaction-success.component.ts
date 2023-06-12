import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-transaction-success',
  templateUrl: './transaction-success.component.html',
  styleUrls: ['./transaction-success.component.css'],
})
export class TransactionSuccessComponent {
  constructor(private router: Router) {}

  submit() {
    this.router.navigate(['/home']);
  }
}
