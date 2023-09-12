import { Injectable } from '@angular/core'; 
 import { 
   HttpRequest, 
   HttpHandler, 
   HttpEvent, 
   HttpInterceptor, 
   HttpErrorResponse, 
 } from '@angular/common/http'; 
 import { AuthService } from './services/auth.service'; 
 import { Observable, throwError, BehaviorSubject } from 'rxjs'; 
 import { catchError, filter, take, switchMap } from 'rxjs/operators'; 
  

 @Injectable() 
 export class TokenInterceptor implements HttpInterceptor { 
   private isRefreshing = false; 
   private refreshTokenSubject: BehaviorSubject<any> = new BehaviorSubject<any>( 
     null 
   ); 
  
   constructor() {}

   intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
   
    
     var token =localStorage.getItem("token");
 
     let request = req;
 
     if (token) {
       request = req.clone({
         setHeaders: {
           authorization: token
         }
       });
     }
 
     return next.handle(request);
   }
 }
