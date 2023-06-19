import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  ngOnInit() {
    const role = localStorage.getItem("role");
    if(!role) {
      localStorage.setItem("role", "guest");
    }
  }
}
