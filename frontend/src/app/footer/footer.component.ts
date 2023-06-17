import { Component } from '@angular/core';

@Component({
  selector: 'app-footer',
  templateUrl: './footer.component.html',
  styleUrls: ['./footer.component.css']
})
export class FooterComponent {
  role: string | null;

  constructor() {
    this.role = localStorage.getItem('role');
  }


}
