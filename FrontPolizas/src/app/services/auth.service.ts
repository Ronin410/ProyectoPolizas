import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http"; 
  
 @Injectable({ 
     providedIn: 'root', 
 }) 
 export class AuthService { 
     userToken: string | null = ''; 
     refToken: string | null = ''; 
  
     constructor(private http: HttpClient) { }


    readToken() { 
        if (localStorage.getItem('token')) { 
            this.userToken = localStorage.getItem('token') ; 
            this.refToken = localStorage.getItem('refresh_token'); 
        } else { 
            this.userToken = ''; 
            this.refToken = ''; 
        } 
        return String(this.userToken); 
    }
}