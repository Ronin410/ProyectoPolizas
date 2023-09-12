import { Component, OnInit } from '@angular/core';
import { PolizasApiService } from '../polizas-api.service';
import { switchMap } from 'rxjs/operators';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'template-polizas',
  templateUrl: './polizas.component.html',
  styleUrls: ['./polizas.component.css'],
})

export class PolizasComponent implements OnInit {
  polizasData: any[] = [];
  idEmpleado: string | null = null;
  mensaje: string |null = null;
  
  constructor(
    private polizasApi: PolizasApiService,
    private _route: ActivatedRoute
  ) {}

  ngOnInit(): void {
    this._route.paramMap.subscribe((params) => {
      const idEmpleado = params.get('idempleado');
      if (idEmpleado) {
        this.ejecutarServicioPolizas(idEmpleado);
      }
    });
  }

  ejecutarServicioPolizas(idEmpleado: string) {
    this.polizasApi.getPolizas(idEmpleado).subscribe(
      (data) => {
        this.polizasData = data.Data;
        console.log(data);
      },
      (error) => {
        // Manejo de errores, si es necesario
        console.error('Error al obtener datos del empleado:', error);
      }
    );
  }

  eliminarPoliza(idpoliza:string, idempleado:string){
    this.polizasApi.postPolizas(idpoliza).subscribe(
      (result) => {
        this.mensaje = result.Data.Message;
        console.log(this.mensaje);
        this.ejecutarServicioPolizas(idempleado)
      },
      (error) => {
        // Manejo de errores, si es necesario
        console.error('Error al obtener datos del empleado:', error);
      }
    );

  }

  modificarPoliza(idpoliza:string){
    console.log("Modificar "+ idpoliza);

  }

}
