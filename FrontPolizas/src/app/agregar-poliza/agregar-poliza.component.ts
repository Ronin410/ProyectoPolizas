import { Component, ElementRef, Input,OnChanges,OnInit, SimpleChanges, ViewChild } from '@angular/core';
import { PolizasApiService} from '../polizas-api.service';
import { EmpleadosComponent } from '../empleados/empleados.component';
//import {PolizasComponent} from '../polizas/polizas.component';
import { take } from 'rxjs';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-agregar-poliza',
  templateUrl: './agregar-poliza.component.html',
  styleUrls: ['./agregar-poliza.component.css']
})

export class AgregarPolizaComponent implements OnInit {
  @Input()
  empleadosData: any;
  inventario : any;
  cantidad : any;
  nombreCliente : any;
  nombreEmpleado : any;
  sku : any;
  vacios = false;
  error = false;
  errorMessage : any;
  cantidadMin = false;


  @ViewChild('myDialog') dialog!: ElementRef;


  constructor(
    private polizasApi: PolizasApiService,
    private Empleados: EmpleadosComponent,
    private toastr: ToastrService
    //private polisasComponent: PolizasComponent
 ) {}

  ngOnInit(): void {
    this.consultarInventario(); 
    this.clear()
  }
  
  consultarInventario() {
    this.polizasApi.getInventario().subscribe(
      (data) => {
        this.inventario = data.Data;
        console.log(data);
      },
      (error) => {
        console.error('Error al obtener datos del inventario:', error);
      }
    );
  }

  clear(){

    let inputCantidad = document.getElementById("cantidad") as HTMLInputElement;
    inputCantidad.value = ""; 
    let inputCliente = document.getElementById("nombre-cliente") as HTMLInputElement;
    inputCliente.value = ""; 
    
  }

  guardarPoliza(){
    this.nombreEmpleado = document.getElementById("idEmpleado") as HTMLInputElement;
    this.sku = document.getElementById("sku") as HTMLInputElement;
    this.cantidad = document.getElementById("cantidad") as HTMLInputElement;
    this.nombreCliente = document.getElementById("nombre-cliente") as HTMLInputElement;

    if( this.cantidad.value != "" && this.nombreCliente.value != "")
    {
      if(+this.cantidad.value >0){
        this.cantidadMin = false;
        var poliza = {
          "empleadogenero":+this.nombreEmpleado.value,
          "sku":+this.sku.value,
          "cantidad":+this.cantidad.value,
          "nombrecliente": this.nombreCliente.value
        }
            
        var polizaGuardar = JSON.stringify(poliza);

        this.polizasApi.postAgregarPoliza(polizaGuardar).subscribe(
          (data) => {
            if(data.Meta.Status == "FAILURE"){
              this.error = true;
              this.toastr.error(data.Data.Message, 'Error');
            }else{
              this.toastr.success(data.Data.Message, 'Exito');
              this.Empleados.mostrarPolizas(this.nombreEmpleado.value);
              this.CerrarDialog();
            }
          },
          (error) => {
            console.error('Error al obtener datos del empleado:', error);
          }
        );
      }
    }else{
      this.vacios = true;
    }
    console.log(this.nombreEmpleado.value);
    console.log(this.sku.value);
    console.log(this.cantidad.value);
    console.log(this.nombreCliente.value);

    console.log("Guardar");
  }
    CerrarDialog(){
    this.clear();
    this.vacios = false;
    this.Empleados.CerrarDialog();
    console.log("Close");

  }

  onKeypressEvent(event : Event){
      this.vacios = false;
  }

}