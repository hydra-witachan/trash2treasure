import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-home-collector',
  templateUrl: './home-collector.component.html',
  styleUrls: ['./home-collector.component.css']
})
export class HomeCollectorComponent {
  constructor(private router: Router) {
  }

  upload() {
    this.router.navigate(['/upload']);
  }
}