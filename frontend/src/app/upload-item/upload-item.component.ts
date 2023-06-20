import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-upload-item',
  templateUrl: './upload-item.component.html',
  styleUrls: ['./upload-item.component.css']
})
export class UploadItemComponent {

  selectOptionData: string[] = ["Organic", "Paper", "Plastic"];
  selectedCategoryOption: string = ""; // To store the selected option
  selecterSubCategoryOption: string = "";
  organicData = ['Banana', 'Leaf', 'Grass', 'Tea'];
  paperData = ['A4', 'Book', 'Duplex', 'Origami Paper'];
  plasticData = ['Spoon', 'Skincare', 'Spray', 'Origami Paper'];

  constructor(private router: Router) {

  }
  


  submit() {
    this.router.navigate(['/home']);
  }
}
