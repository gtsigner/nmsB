import { APP_INITIALIZER, NgModule } from '@angular/core';
import { DebugModule } from '../../storage/debug/debug.module';
import { MessageService } from '../message/message.service';
import { WebSocketModule } from '../websocket/websocket.module';
import { DispatchService } from './dispatch.service';
import { HandlerModule } from './handler/handler.module';

@NgModule({
    imports: [
        DebugModule,
        HandlerModule,
        WebSocketModule
    ],
    providers: [
        DispatchService,
        {
            provide: APP_INITIALIZER,
            deps: [DispatchService],
            useFactory: initialize,
            multi: true
        }
    ]
})
export class DispatchModule {

}

export function initialize(dispatchService: DispatchService) {
    return () => {
        return dispatchService.register();
    };
}
