import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { ActivatedRoute } from '@angular/router';

const HEADERS = {
  headers: new HttpHeaders({ "Content-Type": "application/json" }),
};

@Injectable({
  providedIn: 'root',
})



export class PolizasApiService {
  private apiUrl = 'http://localhost:3000';
  private jsonToken ={appId:"58ccba34-6382-481d-b87f-3fe7d95d430e", appKey: "53d16da74313af15c29d5a486390a572e6255d4855fd7405d7b017a4de06bf76"  };
  HEADER:any;
  
  //localStorage.setItem('token', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwOi8vY29wcGVsLmNvbSIsImF1ZCI6Imh0dHA6Ly9jb3BwZWwuY29tIiwiaWF0IjoxNjkyMDM3NTgxLCJuYmYiOjE2OTIwMzc1ODEsImV4cCI6MTY5MjA2NjY4MSwiYXBwSWQiOiI1OGNjYmEzNC02MzgyLTQ4MWQtYjg3Zi0zZmU3ZDk1ZDQzMGUiLCJlbWFpbCI6bnVsbCwidXNlciI6IjU4Y2NiYTM0LTYzODItNDgxZC1iODdmLTNmZTdkOTVkNDMwZSIsInVzZXJUeXBlIjoiMyIsImxvZ2luVHlwZSI6IjMiLCJkZXZpY2VJZCI6bnVsbCwicGFpcyI6Im1leCIsImV4dGVybm8iOmZhbHNlLCJpcCI6IjAuMC4wLjAiLCJkYXRhIjpudWxsfQ.KWWajtRWVWGsrXGcFDdIgBFTrl3YXOuDX6a0H6Rsi7k');
  private token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwOi8vY29wcGVsLmNvbSIsImF1ZCI6Imh0dHA6Ly9jb3BwZWwuY29tIiwiaWF0IjoxNjkyMDM3NTgxLCJuYmYiOjE2OTIwMzc1ODEsImV4cCI6MTY5MjA2NjY4MSwiYXBwSWQiOiI1OGNjYmEzNC02MzgyLTQ4MWQtYjg3Zi0zZmU3ZDk1ZDQzMGUiLCJlbWFpbCI6bnVsbCwidXNlciI6IjU4Y2NiYTM0LTYzODItNDgxZC1iODdmLTNmZTdkOTVkNDMwZSIsInVzZXJUeXBlIjoiMyIsImxvZ2luVHlwZSI6IjMiLCJkZXZpY2VJZCI6bnVsbCwicGFpcyI6Im1leCIsImV4dGVybm8iOmZhbHNlLCJpcCI6IjAuMC4wLjAiLCJkYXRhIjpudWxsfQ.KWWajtRWVWGsrXGcFDdIgBFTrl3YXOuDX6a0H6Rsi7k';
  constructor(private http: HttpClient) {
    

  }

  getEmpleados(): Observable<any> {
    return this.http.get(`${this.apiUrl}/api.polizas/Empleados`);
  }
  
  getPolizas(parametros: string): Observable<any> {
    return this.http.get(
      `${this.apiUrl}/api.polizas/ConsultarPolizasEmpleado?idempleado=${parametros}`
    );
  }

  postPolizas(idpoliza:string):  Observable<any>{
    return this.http.post(
      `${this.apiUrl}/api.polizas/EliminarPolizas?idpoliza=${idpoliza}`,""
    );
  }

  getInventario(): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}/api.polizas/ConsultarInventario`);
  }

  postToken():  Observable<any>{
    this.HEADER = {
      headers: new HttpHeaders({
         'Content-Type': 'application/json',
         //'Authorization': this.sessionStorageItem.token
      })
    };
    return this.http.post( 'https://apigateway.coppel.com:58443/sso-dev/api/v1/app/authenticate', this.jsonToken, this.HEADER);


  }

  postAgregarPoliza(poliza:any):  Observable<any>{
    return this.http.post(
      `${this.apiUrl}/api.polizas/AgregarPoliza`, poliza, this.HEADER
    );
  }

  postModificarPoliza(poliza:any):  Observable<any>{
    return this.http.post(
      `${this.apiUrl}/api.polizas/ActualizarPoliza2`, poliza, this.HEADER
    );
  }

}
