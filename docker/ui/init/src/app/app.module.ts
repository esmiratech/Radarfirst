import { BrowserModule } from '@angular/platform-browser';
import { ApplicationRef, DoBootstrap, NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { ApiService } from './api.service';
import { AppComponent } from './app.component';
import { ComponentType } from '@angular/cdk/portal';

@NgModule({
  declarations: [
    AppComponent,
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    CommonModule,
    ReactiveFormsModule
  ],
  providers: [ApiService],
  bootstrap: [AppComponent]
})
export class AppModule implements DoBootstrap {
  private bootstrappableComponents: { [key: string]: ComponentType<BootstrappableComponents> } = {
    'app-component': AppComponent,
  };

  ngDoBootstrap(app: ApplicationRef): void {
    const appComponent = 'app-component';
    const rootElements: HTMLCollectionOf<Element> = document.getElementsByTagName('app-root');

    if (rootElements.length === 0) {
      throw new Error('At least one "app-root" element is required to bootstrap radar-app');
    }

    Array.from(rootElements).forEach((element: Element) => {
      // @ts-ignore fix strict null checks when you touch this file, please
      const feature: string = element.getAttribute('feature');
      const component: string = element.getAttribute('component') || appComponent;

      if (feature !== null && component !== appComponent) {
        throw new Error('Bootstrap radar-app with "feature" OR "component" attribute, not both.');
      }

      if (feature === null && component === null) {
        throw new Error('Use either "feature" OR "component" attribute to bootstrap radar-app.');
      }

      if (!Object.keys(this.bootstrappableComponents).includes(component)) {
        throw new Error(`"${component}" is not a bootstrappable component`);
      }

      // append a <my-component> element to the root and bootstrap the app.
      const componentElement = document.createElement(component);
      element.appendChild(componentElement);
      app.bootstrap(this.bootstrappableComponents[component]);
    });
  }
 }
 type BootstrappableComponents =
 | AppComponent;