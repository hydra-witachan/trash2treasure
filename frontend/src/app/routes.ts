import { Routes } from '@angular/router';
import { AuthComponent } from './auth/auth.component';
import { TopUpComponent } from './top-up/top-up.component';
import { ConfirmationTopUpComponent } from './confirmation-top-up/confirmation-top-up.component';
import { TransactionSuccessComponent } from './transaction-success/transaction-success.component';
import { HomeCollectorComponent } from './home-collector/home-collector.component';
import { HomeDonatorComponent } from './home-donator/home-donator.component';
import { HomeGuestComponent } from './home-guest/home-guest.component';
import { UploadItemComponent } from './upload-item/upload-item.component';
import { DonateComponent } from './donate/donate.component';

const role: string | null = localStorage.getItem('role');
const routes: Routes = [
  { path: 'auth', component: AuthComponent },
  { path: 'top-up', component: TopUpComponent },
  { path: 'top-up/confirmation', component: ConfirmationTopUpComponent },
  {
    path: 'top-up/confirmation/success',
    component: TransactionSuccessComponent,
  },
];

if (role === 'collector') {
  routes.push({ path: '', component: HomeCollectorComponent });
  routes.push({ path: 'home', component: HomeCollectorComponent });
  routes.push({ path: 'upload', component: UploadItemComponent });
} else if (role === 'donator') {
  routes.push({ path: '', component: HomeDonatorComponent });
  routes.push({ path: 'home', component: HomeDonatorComponent });
  routes.push({ path: 'donate', component: DonateComponent });
} else {
  // a guest
  routes.push({ path: '', component: HomeGuestComponent });
  routes.push({ path: 'home', component: HomeGuestComponent });
}

export default routes;
