import { Routes } from "@angular/router";
import { AuthComponent } from "./auth/auth.component";
import { TopUpComponent } from "./top-up/top-up.component";
import { ConfirmationTopUpComponent } from "./confirmation-top-up/confirmation-top-up.component";

const routes : Routes = [
    {path: 'auth', component: AuthComponent},
    {path: 'top-up', component: TopUpComponent},
    {path: 'top-up/confirmation', component: ConfirmationTopUpComponent},
];

export default routes;