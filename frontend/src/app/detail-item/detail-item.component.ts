import { Component } from '@angular/core';

@Component({
  selector: 'app-detail-item',
  templateUrl: './detail-item.component.html',
  styleUrls: ['./detail-item.component.css']
})
export class DetailItemComponent {
  itemName = 'Skincare';
  itemDesc = 'I need skincare plastic for recycle to make my own research';
  neededAmount = 100;
  fullfiledAmount = 10;
  pointPerItem = 20;
}
