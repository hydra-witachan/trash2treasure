import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-upload-item',
  templateUrl: './upload-item.component.html',
  styleUrls: ['./upload-item.component.css']
})
export class UploadItemComponent {

  constructor(private router: Router) {

  }


  submit() {
    this.router.navigate(['/home']);
  }
}
