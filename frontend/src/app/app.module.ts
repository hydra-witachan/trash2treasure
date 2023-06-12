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
import { HomeComponent } from './home/home.component';

@NgModule({
  declarations: [
    AppComponent,
    AuthComponent,
    TopUpComponent,
    FooterComponent,
    ConfirmationTopUpComponent,
    TransactionSuccessComponent,
    HomeComponent
  ],
  imports: [
    BrowserModule,
    RouterModule.forRoot(routes),
  ],
  providers: [DecimalPipe],
  bootstrap: [AppComponent]
})
export class AppModule { }
