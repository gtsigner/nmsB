import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { MessageModule } from 'primeng/message';
import { ServerStatusModule } from '../../storage/server-status/server-status.module';
import { StatusBarComponent } from './status-bar.component';

@NgModule({
    imports: [
        // angular
        BrowserModule,
        // primeng
        MessageModule,
        // custom
        ServerStatusModule
    ],
    declarations: [
        StatusBarComponent
    ],
    exports: [
        StatusBarComponent
    ]
})
export class StatusBarModule {

}