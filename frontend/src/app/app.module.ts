import routes from './routes';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { DecimalPipe } from '@angular/common';
import { RouterModule } from '@angular/router';
import { AppComponent } from './app.component';
import { AuthComponent } from './auth/auth.component';
import { TopUpComponent } from './top-up/top-up.component';
import { FooterComponent } from './footer/footer.component';
import { ConfirmationTopUpComponent } from './confirmation-top-up/confirmation-top-up.component';
import { TransactionSuccessComponent } from './transaction-success/transaction-success.component';
import { HomeCollectorComponent } from './home-collector/home-collector.component';
import { HomeDonatorComponent } from './home-donator/home-donator.component';
import { HomeGuestComponent } from './home-guest/home-guest.component';
import { UploadItemComponent } from './upload-item/upload-item.component';
import { DonateComponent } from './donate/donate.component';
import { ListItemComponent } from './list-item/list-item.component';
import { DetailItemComponent } from './detail-item/detail-item.component';
import { SignInComponent } from './sign-in/sign-in.component';
import { SignUpComponent } from './sign-up/sign-up.component';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { RedeemPointsComponent } from './redeem-points/redeem-points.component';
import { NotFoundComponent } from './not-found/not-found.component';
import { SuccessRedeemComponent } from './success-redeem/success-redeem.component';
import { SuccessDonateComponent } from './success-donate/success-donate.component';
import { DetailItemCollectorComponent } from './detail-item-collector/detail-item-collector.component';
import { ProfileComponent } from './profile/profile.component';

@NgModule({
  declarations: [
    AppComponent,
    AuthComponent,
    TopUpComponent,
    FooterComponent,
    ConfirmationTopUpComponent,
    TransactionSuccessComponent,
    HomeCollectorComponent,
    HomeDonatorComponent,
    HomeGuestComponent,
    UploadItemComponent,
    DonateComponent,
    ListItemComponent,
    DetailItemComponent,
    SignInComponent,
    SignUpComponent,
    RedeemPointsComponent,
    NotFoundComponent,
    SuccessRedeemComponent,
    SuccessDonateComponent,
    DetailItemCollectorComponent,
    ProfileComponent,
  ],
  imports: [
    BrowserModule,
    RouterModule.forRoot(routes),
    FormsModule,
    HttpClientModule,
  ],
  providers: [DecimalPipe],
  bootstrap: [AppComponent]
})
export class AppModule { }
