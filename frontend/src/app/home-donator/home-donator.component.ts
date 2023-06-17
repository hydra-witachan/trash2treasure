import { Component } from '@angular/core';
import { Router } from '@angular/router';

type items = {
  name: string,
  imgUrl: string
};

@Component({
  selector: 'app-home-donator',
  templateUrl: './home-donator.component.html',
  styleUrls: ['./home-donator.component.css']
})
export class HomeDonatorComponent {
  activeButton: string = '';
  organicData = [
    {name: "Banana", imgUrl: "https://static.vecteezy.com/system/resources/previews/008/848/350/original/fresh-yellow-banana-fruit-free-png.png"},
    {name: "Leaf", imgUrl: "https://static.vecteezy.com/system/resources/previews/009/664/930/original/green-leaf-on-transparent-background-file-free-png.png" },
    {name: "Grass", imgUrl: "../../assets/grass.png"},
    {name: "Tea", imgUrl: "https://static.vecteezy.com/system/resources/previews/010/856/649/original/a-cup-of-tea-with-leaves-tea-free-png.png"}
  ]

  paperData = [
    {name: "A4", imgUrl: "../../assets/a4.png"},
    {name: "Book", imgUrl: "https://static.vecteezy.com/system/resources/previews/009/377/766/original/3d-book-icon-with-transparent-background-free-png.png" },
    {name: "Duplex", imgUrl: "https://image1ws.indotrading.com/s3/productimages/webp/co249273/p1097364/w425-h425/693406ec-159f-4867-b8cd-f596ad0f690f.png"},
    {name: "Origami Paper", imgUrl: "https://cdn.pixabay.com/photo/2013/07/13/10/08/origami-156627_960_720.png"}
  ]

  plasticData = [
    {name: "Spoon", imgUrl: "../../assets/spoon.png"},
    {name: "Skincare", imgUrl: "https://cdn.pixabay.com/photo/2018/07/16/05/42/skincare-3541261_960_720.png" },
    {name: "Spray", imgUrl: "../../assets/spray.png"},
    {name: "Origami Paper", imgUrl: "../../assets/plastic_Bag.png"}
  ]

  glassData = [
    {name: "Spoon", imgUrl: "../../assets/spoon.png"},
    {name: "Skincare", imgUrl: "https://cdn.pixabay.com/photo/2018/07/16/05/42/skincare-3541261_960_720.png" },
    {name: "Spray", imgUrl: "../../assets/spray.png"},
    {name: "Plastic Bag", imgUrl: "../../assets/plastic_Bag.png"}
  ]
  
  itemData: items[];

  constructor(private router: Router) {
    this.itemData = [];
  }

  setActiveButton(button: string) {
    this.activeButton = button;

    if(button === 'organic') {
      this.itemData = this.organicData;
    } else if(button === 'paper') {
      this.itemData = this.paperData;
    } else if(button === 'plastic') {
      this.itemData = this.plasticData;
    } else if(button === 'glass') {
      // todo: next
    }
  }

  navigateToItem(item: string){
    item = item.toLowerCase();
    if(item.includes(" ")) {
      item = item.replace(" ", "-");
    }
    this.router.navigate([`/category/${item}/items`]);
  }
}
