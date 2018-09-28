import { Injectable } from '@angular/core';
import { ClientHandshakeMessage } from './client-handshake.message';
import { Message } from './message';
import { MessageDirection } from './message.direction';
import { MessageType } from './message.type';

@Injectable()
export class MessageFactoryService {

    private _clientId: string;
    private requestId: number;

    constructor() {
        this.requestId = 0;

    }

    private nextRequestId(): number {
        return this.requestId++;
    }

    clientHandshakeMessage(): ClientHandshakeMessage {
        const message: ClientHandshakeMessage = this.invoke({
            Direction: MessageDirection.CLIENT_2_SERVER,
            Type: MessageType.CLIENT_HANDSHAKE
        } as ClientHandshakeMessage);

        return message;
    }

    injectMessage(): Message {
        const message: ClientHandshakeMessage = this.invoke({
            Direction: MessageDirection.CLIENT_2_SERVER,
            Type: MessageType.INJECT
        } as ClientHandshakeMessage);
        return message;
    }

    set clientId(id: string) {
        this._clientId = id;
    }

    private invoke<T extends Message>(message: T): T {
        const requestId: number = this.nextRequestId();

        message.RequestId = `${requestId}`;
        message.ClientId = this._clientId;

        return message;
    }

}