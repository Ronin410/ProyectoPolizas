import { Component, ElementRef, Input,OnChanges,OnInit, SimpleChanges, ViewChild } from '@angular/core';
import { PolizasApiService} from '../polizas-api.service';
import { EmpleadosComponent } from '../empleados/empleados.component';
import { ToastrService } from 'ngx-toastr';


@Component({
  selector: 'app-modificar-poliza',
  templateUrl: './modificar-poliza.component.html',
  styleUrls: ['./modificar-poliza.component.css'],
  template: '<input #cantidadModificar type="text" />'
})
export class ModificarPolizaComponent implements OnInit, OnChanges {
    @ViewChild('DialogModificar') dialogModificar!: ElementRef;
    @ViewChild('cantidadModificar') cantidadModificarHtml!: ElementRef;

    @Input() polizaModificar : any;
    @Input() empleadosData: any;
    @Input() idpoliza : any;
    @Input() empleadoGenero : any;
    @Input() sku : any;
    @Input() cantidad : any;
    @Input() nombreCliente : any;
    @Input() modificando : any;

    empleados : any;
    inventario : any;
    polizaNueva:any;
    skuModificar:any;
    cantidadMin:any;
    skuModificacion:any;
    vacios = false;
    cantidadModificar:any;
    nombreClienteModificar:any;
    constructor(
      private polizasApi: PolizasApiService,
      private Empleados: EmpleadosComponent,
      private toastr: ToastrService,
    
    ) {}
  ngOnChanges(changes: SimpleChanges): void {
      this.skuModificar = this.sku;
      this.cantidadModificar = this.cantidad;
  }

    ngOnInit(): void {
      this.consultarInventario();      
    }
    
    consultarInventario() {
      this.polizasApi.getInventario().subscribe(
        (data) => {
          this.inventario = data.Data;
        },
        (error) => {
          // Manejo de errores, si es necesario
          console.error('Error al obtener datos del empleado:', error);
        }
      );
    }

    ModificarPoliza(){
    this.skuModificacion = document.getElementById("skuModificar") as HTMLInputElement;
    this.cantidadModificar = document.getElementById("cantidadModificar") as HTMLInputElement;
    this.nombreClienteModificar = document.getElementById("nombreClienteModificar") as HTMLInputElement;

    if( this.cantidadModificar.value != "" && this.nombreClienteModificar.value != "" )
    {
      if(this.cantidadModificar.value >0){
        this.cantidadMin = false;
      
        var poliza = {
        "idpoliza" : +this.idpoliza,
          "empleadogenero":+this.empleadoGenero,
          "sku":+this.skuModificacion.value,
          "cantidad":+this.cantidadModificar.value,
          "nombrecliente": this.nombreClienteModificar.value
      }

        this.polizasApi.postModificarPoliza(poliza).subscribe(
          (data) => {
            if(data.Meta.Status == "FAILURE"){
              this.toastr.error(data.Data.Message, 'Error');
            }else{
              this.toastr.success(data.Data.Message, 'Exito');
              this.Empleados.mostrarPolizas(this.empleadoGenero);
              this.Empleados.CerrarDialogModificar();
              
            }
          },
          (error) => {
            // Manejo de errores, si es necesario
            console.error('Error al obtener datos del empleado:', error);
          }
        );
      }else{
        this.cantidadMin = true;
      }

    }else{
        this.vacios = true;
      }
    
  }
    CerrarDialog(){
      const cantidadElement = document.getElementById('cantidadModificar') as HTMLInputElement;
      cantidadElement.value = this.cantidad;
      const nombreClienteElement = document.getElementById('nombreClienteModificar') as HTMLInputElement;
      nombreClienteElement.value = this.nombreCliente;
      this.skuModificar = this.sku;

      this.Empleados.CerrarDialogModificar();
      this.empleados= this.empleadoGenero;
      this.vacios = false;
      this.Empleados.mostrarPolizas(this.empleados);      
    }

    onKeypressEvent(event : Event){
      this.vacios = false;
      this.cantidadMin = false;

  }
}
