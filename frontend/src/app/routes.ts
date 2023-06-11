import { Routes } from "@angular/router";
import { AuthComponent } from "./auth/auth.component";
import { TopUpComponent } from "./top-up/top-up.component";

const routes : Routes = [
    {path: 'auth', component: AuthComponent},
    {path: 'top-up', component: TopUpComponent},
];

export default routes;