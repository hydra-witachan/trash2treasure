import { Component } from '@angular/core';
import { MySharedService } from '../shared/my-shared-service.service';

@Component({
  selector: 'app-footer',
  templateUrl: './footer.component.html',
  styleUrls: ['./footer.component.css']
})
export class FooterComponent {

  constructor(public shared: MySharedService) {
  }

}
