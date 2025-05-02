import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable()
export class ApiService {
  private apiUrl = 'http://localhost:8080/hello';

  constructor(private http: HttpClient) { }

  sayHello(name: string) {
    return this.http.post<any>(this.apiUrl, { name });
  }
}