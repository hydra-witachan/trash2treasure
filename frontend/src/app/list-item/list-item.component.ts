import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { MySharedService } from '../shared/my-shared-service.service';

@Component({
  selector: 'app-list-item',
  templateUrl: './list-item.component.html',
  styleUrls: ['./list-item.component.css']
})
export class ListItemComponent {
  title = "";
  items: any[] = [];

  searchQuery = "";

  constructor(private route: ActivatedRoute, private http: HttpClient, private router: Router, public sharedService: MySharedService) { }

  ngOnInit() {
    const subCategory = this.route.snapshot.paramMap.get('subCategory');
    this.title = this.sharedService.getSubCategory();
    console.log(subCategory); // Do whatever you need with the subCategory value
    const url = `http://localhost:5000/items?sub_category=${subCategory}&search=`;
    console.log(url); // Use the URL for your API request
    this.http.get(url).subscribe((response: any) => {
      this.items = response;
    }, (error: any) => console.log(error));
  }
  displayItemDetails(item: any) {
    this.router.navigate([`/items/${item.id}`]);
  }

  onSearchInputChange() {
    const subCategory = this.route.snapshot.paramMap.get('subCategory');
    const url = `http://localhost:5000/items?sub_category=${subCategory}&search=${this.searchQuery}`;
    console.log(url); // Use the URL for your API request

    this.http.get(url).subscribe((response: any) => {
      this.items = response;
    }, (error: any) => console.log(error));
  }
}
