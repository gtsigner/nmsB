import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { CommunicationModule } from './communication/communication.module';

@NgModule({
    imports: [
        BrowserModule,
        CommunicationModule.forRoot()
    ],
    declarations: [
        AppComponent

    ],
    providers: [],
    bootstrap: [AppComponent]
})
export class AppModule {

}
