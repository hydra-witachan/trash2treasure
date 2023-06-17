import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { DecimalPipe } from '@angular/common';
import { RouterModule } from '@angular/router';
import { AppComponent } from './app.component';
import { AuthComponent } from './auth/auth.component';
import { TopUpComponent } from './top-up/top-up.component';
import routes from './routes';
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
    SignUpComponent
  ],
  imports: [
    BrowserModule,
    RouterModule.forRoot(routes),
  ],
  providers: [DecimalPipe],
  bootstrap: [AppComponent]
})
export class AppModule { }
