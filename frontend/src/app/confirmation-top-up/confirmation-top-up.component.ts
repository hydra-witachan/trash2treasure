import { Component } from '@angular/core';
import { MySharedService } from '../shared/my-shared-service.service';

@Component({
  selector: 'app-confirmation-top-up',
  templateUrl: './confirmation-top-up.component.html',
  styleUrls: ['./confirmation-top-up.component.css']
})
export class ConfirmationTopUpComponent {
  topUpPoint: any;
  formattedPoint: any;
  feeTopUp: any;

  constructor(private sharedService: MySharedService) {}

  ngOnInit() {
    this.topUpPoint = this.sharedService.getPointTopUp();
    this.formattedPoint = this.sharedService.getPointTopUpFormated();
    this.feeTopUp = this.sharedService.getFeePointTopUp();
  }
}
