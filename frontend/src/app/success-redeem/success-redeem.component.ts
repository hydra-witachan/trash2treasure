import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-success-redeem',
  templateUrl: './success-redeem.component.html',
  styleUrls: ['./success-redeem.component.css']
})
export class SuccessRedeemComponent {
  constructor(private router: Router) {}
  submit() {
    this.router.navigate(['/home']);
  }
}
