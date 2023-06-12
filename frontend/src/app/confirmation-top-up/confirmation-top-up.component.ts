import { Component } from '@angular/core';
import { MySharedService } from '../shared/my-shared-service.service';
import { DecimalPipe } from '@angular/common';

@Component({
  selector: 'app-confirmation-top-up',
  templateUrl: './confirmation-top-up.component.html',
  styleUrls: ['./confirmation-top-up.component.css']
})
export class ConfirmationTopUpComponent {
  topUpPoint: any;
  formattedPoint: any;
  feeTopUp: any;
  formattedFeeTopUp: any;

  constructor(private sharedService: MySharedService, private decimalPipe: DecimalPipe) {}

  ngOnInit() {
    this.topUpPoint = this.sharedService.getPointTopUp();
    this.formattedPoint = this.sharedService.getPointTopUpFormated();
    this.feeTopUp = this.sharedService.getFeePointTopUp();
    this.formattedFeeTopUp = this.decimalPipe.transform(this.feeTopUp, '1.0-3');
    if(this.formattedFeeTopUp !== null) {
      this.formattedFeeTopUp =  this.formattedFeeTopUp.replaceAll(',', '.')
    }
  }

  submit() {
    console.log("HELLO");
  }
}
