import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { MessageModule } from 'primeng/message';
import { TableModule } from 'primeng/table';
import { DebugModule } from '../../storage/debug/debug.module';
import { DebugTypeComponent } from './debug-type.component';
import { DebugViewComponent } from './debug-view.component';

@NgModule({
    imports: [
        // angular
        BrowserModule,
        // primeng
        TableModule,
        MessageModule,
        // custom
        DebugModule
    ],
    declarations: [
        DebugViewComponent,
        DebugTypeComponent
    ]
})

export class DebugViewModule {

}
