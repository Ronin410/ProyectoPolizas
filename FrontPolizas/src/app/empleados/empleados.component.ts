import { AfterViewInit, Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { PolizasApiService } from '../polizas-api.service';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'template-empleados',
  templateUrl: './empleados.component.html',
  styleUrls: ['./empleados.component.css'],
  template: ` <app-mostrar-polizas [idEmpleado]="idEmpleadoBuscar"></app-mostrar-polizas> `,

})
export class EmpleadosComponent implements OnInit,AfterViewInit {
  empleadosData: any;
  token : any;
  polizasData  : any;
  @ViewChild('myDialogPolizas') dialogPolizas!: ElementRef;
  @ViewChild('myDialog') dialog!: ElementRef;
  @ViewChild('myDialogModificar') dialogModificar!: ElementRef;
  mostrar = false;
  modificando ="1";
  idpoliza = 0;
  empleadoGenero = 0;
  sku = 0;
  cantidad = 0;
  nombreCliente = "";

  constructor(
    private polizasApi: PolizasApiService,
  ) {}
  ngAfterViewInit(): void {
  }

  ngOnInit(): void {

    this.Obtenertoken();
    this.consultarEmpleado();
  }

  consultarEmpleado(){
    this.polizasApi.getEmpleados().subscribe(
      (response) => {
        this.empleadosData = response.Data;
      },
      (error) => {
        console.error('Error al obtener datos:', error);
      }
    );
  }


  AbrirDialog(){
    let empleados = document.getElementById('empleados') as HTMLElement;
    empleados?.style.setProperty('pointer-events', 'none');
    let buttonAgregar = document.getElementById('buttonAgregar') as HTMLElement;
    buttonAgregar?.style.setProperty('pointer-events', 'none'); 
    this.dialog.nativeElement.show();
    this.mostrar = false;
  } 

  CerrarDialog(){
    let empleados = document.getElementById('empleados') as HTMLElement;
    empleados?.style.setProperty('pointer-events', 'auto');
    let buttonAgregar = document.getElementById('buttonAgregar') as HTMLElement;
    buttonAgregar?.style.setProperty('pointer-events', 'auto'); 
    this.dialog.nativeElement.close();

  }

  CerrarDialog2(){
    let empleados = document.getElementById('empleados') as HTMLElement;
    empleados?.style.setProperty('pointer-events', 'auto');
    
    let buttonAgregar = document.getElementById('buttonAgregar') as HTMLElement;
    buttonAgregar?.style.setProperty('pointer-events', 'auto'); 
    this.dialogPolizas.nativeElement.close();
  }

  CerrarDialogModificar(){
    let empleados = document.getElementById('empleados') as HTMLElement;
    empleados?.style.setProperty('pointer-events', 'auto');
    
    let buttonAgregar = document.getElementById('buttonAgregar') as HTMLElement;
    buttonAgregar?.style.setProperty('pointer-events', 'auto'); 
    this.dialogModificar.nativeElement.close();
  }


  Obtenertoken(){
    if(localStorage.getItem('token') == null){
      this.polizasApi.postToken().subscribe(
        (response) => {
          this.token = response.data;
          localStorage.setItem('token', response.data.token);
          this.consultarEmpleado();
        },
        (error) => {
          console.error('Error al obtener datos:', error);
        }
      );
    }
  }

  consultarPolizas(idEmpleado: string) {
    //this.polizasData = null;
    let empleados = document.getElementById('empleados') as HTMLElement;
    empleados?.style.setProperty('pointer-events', 'none');
    
    let buttonAgregar = document.getElementById('buttonAgregar') as HTMLElement;
    buttonAgregar?.style.setProperty('pointer-events', 'none'); 
    this.polizasApi.getPolizas(idEmpleado).subscribe(
      (data) => {
        this.polizasData = data.Data;
      },
      (error) => {
        // Manejo de errores, si es necesario
        console.error('Error al obtener datos del empleado:', error);
      }
    );
  }

  mostrarPolizas(idEmpleado: any){
    this.consultarPolizas(idEmpleado);
    this.dialogPolizas.nativeElement.show();
    this.mostrar = false;
  }

  modificarPoliza(polizaModificar : any){
    console.log("Modificar "+ polizaModificar);
    let empleados = document.getElementById('empleados') as HTMLElement;
    empleados?.style.setProperty('pointer-events', 'none');
    
    let buttonAgregar = document.getElementById('buttonAgregar') as HTMLElement;
    buttonAgregar?.style.setProperty('pointer-events', 'none'); 
    this.idpoliza = polizaModificar.IdPoliza;
    this.empleadoGenero = polizaModificar.EmpleadoGenero;
    this.sku = polizaModificar.Sku;
    this.cantidad = polizaModificar.Cantidad;
    this.nombreCliente = polizaModificar.NombreCliente;

    this.dialogModificar.nativeElement.show();

  }
}