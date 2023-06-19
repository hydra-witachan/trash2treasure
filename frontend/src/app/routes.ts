import { Routes } from '@angular/router';
import { AuthComponent } from './auth/auth.component';
import { TopUpComponent } from './top-up/top-up.component';
import { ConfirmationTopUpComponent } from './confirmation-top-up/confirmation-top-up.component';
import { TransactionSuccessComponent } from './transaction-success/transaction-success.component';
import { HomeCollectorComponent } from './home-collector/home-collector.component';
import { HomeDonatorComponent } from './home-donator/home-donator.component';
import { UploadItemComponent } from './upload-item/upload-item.component';
import { DonateComponent } from './donate/donate.component';
import { ListItemComponent } from './list-item/list-item.component';
import { DetailItemComponent } from './detail-item/detail-item.component';
import { SignInComponent } from './sign-in/sign-in.component';
import { SignUpComponent } from './sign-up/sign-up.component';
import { RedeemPointsComponent } from './redeem-points/redeem-points.component';

let role: string | null;

export function SetupRoutes(): Routes {
  role = localStorage.getItem('role');

  const routes: Routes = [
    { path: 'auth', component: AuthComponent },
    { path: 'top-up', component: TopUpComponent },
    { path: 'top-up/confirmation', component: ConfirmationTopUpComponent },
    { path: 'auth/sign-in', component: SignInComponent },
    { path: 'auth/sign-up', component: SignUpComponent },
    {
      path: 'top-up/confirmation/success',
      component: TransactionSuccessComponent,
    },
  ]

  if (role === 'collector') {
    routes.push({ path: '', component: HomeCollectorComponent });
    routes.push({ path: 'home', component: HomeCollectorComponent });
    routes.push({ path: 'upload', component: UploadItemComponent });
  } else if (role === 'donator') {
    routes.push({ path: '', component: HomeDonatorComponent });
    routes.push({ path: 'home', component: HomeDonatorComponent });
    routes.push({ path: 'donate', component: DonateComponent });
    routes.push({ path: 'category/:subCategory/items', component: ListItemComponent });
    routes.push({ path: 'items/:id', component: DetailItemComponent });
    routes.push({ path: 'redeem-points', component: RedeemPointsComponent });
  } else {
    // a guest
    routes.push({ path: '', component: AuthComponent });
  }

  return routes;
}

export default SetupRoutes();
