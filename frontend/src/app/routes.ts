import { Routes } from "@angular/router";
import { AuthComponent } from "./auth/auth.component";
import { TopUpComponent } from "./top-up/top-up.component";
import { ConfirmationTopUpComponent } from "./confirmation-top-up/confirmation-top-up.component";
import { TransactionSuccessComponent } from "./transaction-success/transaction-success.component";
import { HomeComponent } from "./home/home.component";

const routes : Routes = [
    {path: 'home', component: HomeComponent},
    {path: '', component: HomeComponent},
    {path: 'auth', component: AuthComponent},
    {path: 'top-up', component: TopUpComponent},
    {path: 'top-up/confirmation', component: ConfirmationTopUpComponent},
    {path: 'top-up/confirmation/success', component: TransactionSuccessComponent},
];

export default routes;