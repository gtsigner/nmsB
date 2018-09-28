import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AppRoutingModule } from './app-routing.module';

import { AppComponent } from './app.component';
import { CommunicationModule } from './communication/communication.module';
import { ViewsModule } from './views/views.module';

@NgModule({
    imports: [
        BrowserModule,
        BrowserAnimationsModule,

        // Custom
        ViewsModule,
        AppRoutingModule,
        CommunicationModule
    ],
    declarations: [
        AppComponent

    ],
    providers: [],
    bootstrap: [AppComponent]
})
export class AppModule {

}
