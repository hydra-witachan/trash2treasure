import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-detail-item-collector',
  templateUrl: './detail-item-collector.component.html',
  styleUrls: ['./detail-item-collector.component.css'],
})
export class DetailItemCollectorComponent {
  itemName = 'Skincare';
  itemDesc = 'I need skincare plastic for recycle to make my own research';
  neededAmount = 100;
  fullfiledAmount = 10;
  pointPerItem = 20;
  item: any = [];

  constructor(
    private route: ActivatedRoute,
    private http: HttpClient,
    private router: Router
  ) {}

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

  cancel() {
    this.router.navigate([`/home`]);
  }

  goToProfile() {
    this.router.navigate([`/profile/${this.item.authorId}`])
  }

}
