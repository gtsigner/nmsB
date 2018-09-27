import { Injectable } from '@angular/core';
import { Subscription } from 'rxjs';
import { WebSocketService } from '../websocket/websocket.service';
import { Message } from './message';
import { MessageCallback } from './message-callback';
import { MessageFactoryService } from './message-factory.service';

@Injectable()
export class MessageService {

    private subscription: Subscription;
    private pending: Map<string, MessageCallback<any>>;

    constructor(private webSocketService: WebSocketService,
                private messageFactory: MessageFactoryService) {
        this.pending = new Map<string, MessageCallback<any>>();
    }

    async register(): Promise<void> {
        return new Promise<void>((resolve, reject) => {
            // wait for the websocket to connected
            this.webSocketService.onConnected(() => {
                // register the callback handle for the messages
                this.registerCallbackHandle();

                setTimeout(() => {
                    // execute the handshake
                    this.handshake().then(resolve, reject);
                }, 1000);
            });
        });
    }

    private registerCallbackHandle(): void {
        const complete = () => {
            if (this.subscription) {
                this.subscription.unsubscribe();
            }
        };
        complete();

        // subscribe to the websocket
        this.subscription = this.webSocketService.subscribe(
            (message: any) => {
                // wait for incoming messages and handle for pending
                this.handleMessage(message);
            },
            undefined,
            complete
        );
    }

    private handleMessage<T extends Message>(message: T): void {
        // get the request id
        const requestId: string = message.RequestId;

        // verify if request id given
        if (!requestId) {
            // TODO error
            return;
        }

        // check if the request is pending
        if (!this.pending.has(requestId)) {
            return;
        }
        // get the callback
        const callback: MessageCallback<T> = this.pending.get(requestId);
        // executed the callback
        callback(message);
        // remove the pending callback
        this.pending.delete(requestId);
    }

    private async handshake(): Promise<void> {
        const message: Message = this.messageFactory.clientHandshakeMessage();
        const response: Message = await this.request(message);
        this.messageFactory.clientId = response.ClientId;
    }

    send<T extends Message>(message: T): void {
        this.webSocketService.send(message);
    }

    async request<T extends Message, R extends Message>(message: T): Promise<R> {
        // check if a request id given
        const requestId: string = message.RequestId;
        if (!requestId) {
            throw new Error(`unable to request message [ ${message} ], because message is missing request id`);
        }

        // create the promise for the request
        const promise: Promise<R> = new Promise(resolve => {
            // the callback for the response message
            const callback: MessageCallback<R> = (response: R) => {
                // delete the callback from pending
                this.pending.delete(requestId);
                // resolve the promise
                resolve(response);
            };
            // set the callback as pending
            this.pending.set(requestId, callback);
        });

        // send the message
        this.send(message);

        return promise;
    }

}
