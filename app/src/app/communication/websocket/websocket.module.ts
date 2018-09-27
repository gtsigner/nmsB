import { APP_INITIALIZER, NgModule } from '@angular/core';
import { WebSocketService } from './websocket.service';

@NgModule({
    providers: [
        WebSocketService,
        {
            provide: APP_INITIALIZER,
            deps: [WebSocketService],
            useFactory: initialize,
            multi: true
        }
    ]
})
export class WebSocketModule {

}

export function initialize(webSocketService: WebSocketService) {
    return () => {
        webSocketService.connect();
    };
}
