import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-list-item',
  templateUrl: './list-item.component.html',
  styleUrls: ['./list-item.component.css']
})
export class ListItemComponent {
  subCategory = "Skincare";
  items: any[] = [];

  constructor(private route: ActivatedRoute, private http: HttpClient, private router: Router) { }

  ngOnInit() {
    const subCategory = this.route.snapshot.paramMap.get('subCategory');
    console.log(subCategory); // Do whatever you need with the subCategory value
    const url = `http://localhost:5000/items?sub_category=${subCategory}&search=`;
    console.log(url); // Use the URL for your API request
    this.http.get(url).subscribe((response: any) => {
      this.items = response;
      console.log(this.items);
    }, (error: any) => console.log(error));
  }
  displayItemDetails(item: any) {
    this.router.navigate([`/items/${item.id}`]);
  }
}
