import { APP_INITIALIZER, ModuleWithProviders, NgModule } from '@angular/core';
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

    static forRoot(): ModuleWithProviders {
        return {
            ngModule: CommunicationModule,
            providers: [
                {
                    provide: APP_INITIALIZER,
                    deps: [],
                    useFactory: initialize,
                    multi: true
                }
            ]
        };
    }
}

export function initialize() {
    return () => {
        console.log('CommunicationModuleP');
    };

}
