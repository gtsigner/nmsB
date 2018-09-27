import { Injectable } from '@angular/core';
import { WebSocketService } from '../websocket/websocket.service';
import { Message } from './message';
import { MessageFactoryService } from './message-factory.service';

@Injectable()
export class MessageService {

    constructor(private webSocketService: WebSocketService, private messageFactory: MessageFactoryService) {

    }

    private init(): void {
        this.webSocketService.onOpen(() => {
            this.handshake();
        });

    }

    handshake(): void {
        const message: Message = this.messageFactory.clientHandshakeMessage();
        this.webSocketService.send(message);
    }

}
