import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { ButtonModule } from 'primeng/button';
import { MessageModule } from '../../communication/message/message.module';
import { ServerStatusModule } from '../../storage/server-status/server-status.module';
import { ControlViewComponent } from './control-view.component';

@NgModule({
    imports: [
        // angular
        BrowserModule,
        // primeng
        ButtonModule,
        // custom
        MessageModule,
        ServerStatusModule
    ],
    declarations: [
        ControlViewComponent
    ]
})
export class ControlViewModule {

}
