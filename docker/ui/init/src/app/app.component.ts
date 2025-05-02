import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ApiService } from './api.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent implements OnInit{
  myForm: FormGroup = new FormGroup({});
  message: string = ''; // Assign an initial value to the 'message' property

  constructor(private formBuilder: FormBuilder, private apiService: ApiService) {}
  
  ngOnInit(): void {
    this.myForm = this.formBuilder.group({
      name: ['', Validators.required]
    });
  }
  
  onSubmit(): void {
    if (this.myForm.valid) {
      const name = this.myForm.value.name;
      this.apiService.sayHello(name).subscribe({
        next: (response) => {
          this.message = response.message;
        },
        error: (error) => {
          console.error('Error:', error);
          this.message = 'Error occurred. Please try again.';
        }
      });
    }
  }
}