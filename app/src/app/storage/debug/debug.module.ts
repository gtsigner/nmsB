import { NgModule } from '@angular/core';
import { DebugService } from './debug.service';
import { MessagesModule } from 'primeng/messages';
import { MessageService } from 'primeng/api';

@NgModule({
    providers: [
        DebugService,
        MessageService
    ]
})
export class DebugModule {

}