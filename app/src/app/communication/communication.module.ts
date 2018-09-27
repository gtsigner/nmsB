import { NgModule } from '@angular/core';
import { DispatchModule } from './dispatch/dispatch.module';
import { MessageModule } from './message/message.module';
import { WebSocketModule } from './websocket/websocket.module';
import { WebSocketService } from './websocket/websocket.service';

@NgModule({
    imports: [
        MessageModule,
        DispatchModule,
        WebSocketModule
    ]
})
export class CommunicationModule {

    constructor(private webSocketService: WebSocketService) {
    }

}
