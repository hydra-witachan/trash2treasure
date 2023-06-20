import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-upload-item',
  templateUrl: './upload-item.component.html',
  styleUrls: ['./upload-item.component.css'],
})
export class UploadItemComponent {
  selectOptionData: string[] = ['Organic', 'Paper', 'Plastic'];
  selectedCategoryOption: string = ''; // To store the selected option
  selecterSubCategoryOption: string = '';
  organicData = ['Banana', 'Leaf', 'Grass', 'Tea'];
  paperData = ['A4', 'Book', 'Duplex', 'Origami Paper'];
  plasticData = ['Spoon', 'Skincare', 'Spray', 'Origami Paper'];
  encodedImage: any;
  itemName: string = '';
  description: string = '';
  pointsPerItem: any;
  neededAmount: any;
  accessToken: string = '';
  previewImageUrl: string = '../../assets/default-avatar.png';
  imageUploaded: boolean = false;
  loading: boolean = false; // Flag to indicate loading state

  constructor(private router: Router, private http: HttpClient) {
    const token = localStorage.getItem('accessToken');
    if(token) {
      this.accessToken = token;
    }
  }

  handleFileInput(event: any) {
    const file = event.target.files[0];
    if (file) {
      const reader = new FileReader();
      reader.readAsDataURL(file);
      reader.onload = () => {
        const base64Image = reader.result as string;
        this.previewImageUrl = reader.result as string;
        this.imageUploaded = true;
        this.encodedImage = base64Image.split(',')[1]; // Remove 'base64,' from the string
      };
    }
  }

  submit() {
    const apiUrl = 'http://localhost:5000/items';

    const formData = {
      itemName: this.itemName,
      description: this.description,
      pointsPerItem: parseInt(this.pointsPerItem),
      neededAmount: parseInt(this.neededAmount),
      subCategory: this.selecterSubCategoryOption,
      encodedImage: this.encodedImage,
    };
    
    const token = this.accessToken;
    console.log('object');
    console.log(token);
    console.log(formData);
    if (token) {
      const headers = new HttpHeaders({
        Authorization: `Bearer ${token}`,
      });
      console.log('post');
      this.loading = true; // Set loading flag to true
      this.http.post(apiUrl, formData, { headers }).subscribe(
        (response) => {
          console.log('Item uploaded successfully:', response);
          this.router.navigate(['/home']);
          this.loading = false; // Set loading flag to false after successful upload
        },
        (error) => {
          console.error('Error uploading item:', error);
          this.loading = false; // Set loading flag to false after successful upload
        }
      );
    }
  }
}
