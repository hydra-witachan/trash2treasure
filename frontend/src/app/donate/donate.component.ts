import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Component } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-donate',
  templateUrl: './donate.component.html',
  styleUrls: ['./donate.component.css'],
})
export class DonateComponent {
  accessToken: string | null = '';
  item: any = [];
  name: string | null = '';
  phoneNumber: string = '';
  quantity: number | string = '';
  address: string = '';

  constructor(
    private route: ActivatedRoute,
    private http: HttpClient,
    private router: Router
  ) {
    const token = localStorage.getItem('accessToken');
    if (token) {
      this.accessToken = localStorage.getItem('accessToken');
    }
  }

  ngOnInit() {
    const id = this.route.snapshot.paramMap.get('id');
    console.log(id); // Do whatever you need with the subCategory value
    const url = `http://localhost:5000/items/${id}`;
    console.log(url); // Use the URL for your API request

    this.http.get(url).subscribe(
      (response: any) => {
        this.item = response;
        console.log('DETAIL ITEM');
        console.log(this.item);
      },
      (error: any) => console.log(error)
    );
  }

  submit() {
    if (typeof this.quantity === 'string') {
      this.quantity = parseInt(this.quantity);
    }

    console.log(this.name);
    console.log(this.phoneNumber);
    console.log(this.address);
    console.log(this.quantity);
    console.log(typeof this.quantity);

    const token = this.accessToken;
    if (token) {
      const headers = new HttpHeaders({
        Authorization: `Bearer ${token}`,
      });

      const itemId: any = this.item.id;
      const url = `http://localhost:5000/items/donate/${itemId}`;
      console.log(url);

      const body = {
        name: 'Alvian',
        phoneNumber: this.phoneNumber,
        address: this.address,
        quantity: this.quantity,
      };

      console.log(url);
      console.log(body);

      this.http.post(url, body, { headers }).subscribe(
        (response) => {
          console.log(response);
          this.router.navigate(['/success/donate']);
        },
        (error) => {
          console.log(error);
        }
      );
    }
  }
}
