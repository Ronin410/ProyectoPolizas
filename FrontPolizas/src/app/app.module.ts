import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { EmpleadosComponent } from './empleados/empleados.component';
import { PolizasComponent } from './polizas/polizas.component';
import { HTTP_INTERCEPTORS, HttpClientModule } from '@angular/common/http';
import { AgregarPolizaComponent } from './agregar-poliza/agregar-poliza.component';
import { AuthService } from './services/auth.service';
import { TokenInterceptor } from './token.interceptor';
import { MostrarPolizasComponent } from './mostrar-polizas/mostrar-polizas.component';
import { ModificarPolizaComponent } from './modificar-poliza/modificar-poliza.component';
import { FormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ToastrModule } from 'ngx-toastr';
@NgModule({
  declarations: [AppComponent, EmpleadosComponent, PolizasComponent, AgregarPolizaComponent, MostrarPolizasComponent, ModificarPolizaComponent],
  imports: [BrowserModule, AppRoutingModule, HttpClientModule,FormsModule, BrowserAnimationsModule,    ToastrModule.forRoot()],
  providers: [
    AuthService,
    {
      provide: HTTP_INTERCEPTORS,    
      useClass: TokenInterceptor,
      multi: true
    }
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}
