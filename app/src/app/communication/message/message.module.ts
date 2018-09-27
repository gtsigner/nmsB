import { APP_INITIALIZER, NgModule } from '@angular/core';
import { WebSocketModule } from '../websocket/websocket.module';
import { MessageFactoryService } from './message-factory.service';
import { MessageService } from './message.service';

@NgModule({
    imports: [
        WebSocketModule
    ],
    providers: [
        MessageService,
        MessageFactoryService,
        {
            provide: APP_INITIALIZER,
            deps: [MessageService],
            useFactory: initialize,
            multi: true
        }
    ]
})
export class MessageModule {
}

export function initialize(messageService: MessageService) {
    return () => {
        return messageService.register();
    };
}
