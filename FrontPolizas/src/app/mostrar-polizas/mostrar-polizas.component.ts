import { AfterViewInit, Component, ElementRef, Input, OnInit, ViewChild } from '@angular/core';
import { PolizasApiService } from '../polizas-api.service';
import { ActivatedRoute } from '@angular/router';
import { EmpleadosComponent } from '../empleados/empleados.component';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-mostrar-polizas',
  templateUrl: './mostrar-polizas.component.html',
  styleUrls: ['./mostrar-polizas.component.css']
})


export class MostrarPolizasComponent implements OnInit, AfterViewInit {
  @Input() polizasData: any[] = [];
  mensaje: string |null = null;
  idEmpleado: any | undefined;
  error = false;
  constructor(
    private polizasApi: PolizasApiService,
    private _route: ActivatedRoute,
    private empleadoComponent : EmpleadosComponent,
    private toastr: ToastrService

  ) {}
  ngAfterViewInit(): void {
    this._route.paramMap.subscribe((params) => {
      if (this.idEmpleado) {
        this.ejecutarServicioPolizas(this.idEmpleado);
      }
    }); 
  }

  ngOnInit(): void {
 
  }


  ejecutarServicioPolizas(idEmpleado: string) {
    this.polizasApi.getPolizas(idEmpleado).subscribe(
      (data) => {
        this.polizasData = data.Data;
      },
      (error) => {
        console.error('Error al obtener datos del empleado:', error);
      }
    );
  }

  eliminarPoliza(idpoliza:string, idempleado:string){
    this.polizasApi.postPolizas(idpoliza).subscribe(
      (result) => {
        if(result.Meta.Status == "FAILURE"){
          this.error = true;
          this.toastr.error(result.Data.Message, 'Error');
        }else{
          this.toastr.success(result.Data.Message, 'Exito');
          this.ejecutarServicioPolizas(idempleado)
        }

      },
      (error) => {
        // Manejo de errores, si es necesario
        console.error('Error al obtener datos del empleado:', error);
      }
    );

  }

  modificarPoliza(polizaModificar : any){
    this.polizasData = [];
    this.empleadoComponent.CerrarDialog2();
    this.empleadoComponent.modificarPoliza(polizaModificar);
  }

  CerrarDialog(){
    this.empleadoComponent.CerrarDialog2();
  }
  
}
