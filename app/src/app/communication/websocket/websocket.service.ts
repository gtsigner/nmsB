import { EventEmitter, Injectable } from '@angular/core';
import { Subscription } from 'rxjs';
import { Message } from '../message/message';

@Injectable()
export class WebSocketService {

    private socket: WebSocket;
    private connected: EventEmitter<void>;
    private eventEmitter: EventEmitter<any>;

    constructor() {
        this.connected = new EventEmitter();
    }

    private socketUrl(): string {
        const wsUrl: URL = new URL('ws', window.location.href);
        wsUrl.protocol = wsUrl.protocol.replace(/^http/, 'ws');
        return wsUrl.href;
    }

    connect(): void {
        this.eventEmitter = new EventEmitter();

        const url: string = this.socketUrl();
        this.socket = new WebSocket(url);


        this.socket.onopen = (e: Event) => {
            this.connected.emit();
        };

        this.socket.onmessage = (e: MessageEvent) => {
            if (e.data) {
                const message: any = JSON.parse(e.data);
                this.eventEmitter.next(message);
            }
        };

        this.socket.onerror = (e: Event) => {
            this.eventEmitter.error(e);
        };

        this.socket.onclose = (e: CloseEvent) => {
            this.eventEmitter.complete();
            this.socket = undefined;
            this.eventEmitter = undefined;
        };
    }

    onConnected(listener: () => void): Subscription {
        const subscription: Subscription = this.connected.subscribe(listener);
        return subscription;
    }

    subscribe(generatorOrNext?: any, error?: any, complete?: any): Subscription {
        const subscription: Subscription = this.eventEmitter.subscribe(generatorOrNext, error, complete);
        return subscription;
    }

    send<T extends Message>(message: T): void {
        if (this.socket && this.socket.readyState === WebSocket.OPEN) {
            const data: string = JSON.stringify(message);
            this.socket.send(data);
        }
    }

    close(): void {
        if (this.socket && this.socket.readyState === WebSocket.OPEN) {
            this.socket.close();
        }
    }


}