import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { PolizasComponent } from './polizas/polizas.component';

const routes: Routes = [
  { path: 'polizas/:idempleado', component: PolizasComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
