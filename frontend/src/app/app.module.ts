import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule } from '@angular/router';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { AuthComponent } from './auth/auth.component';
import { TopUpComponent } from './top-up/top-up.component';

@NgModule({
  declarations: [
    AppComponent,
    AuthComponent,
    TopUpComponent
  ],
  imports: [
    BrowserModule,
    RouterModule.forRoot([
      {path: 'auth', component: AuthComponent},
      {path: 'top-up', component: TopUpComponent},
    ]),
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
