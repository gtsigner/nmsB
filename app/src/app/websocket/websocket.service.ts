import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

@Injectable()
export class WebSocketService {

    private socket: WebSocket;
    private subject: Subject<MessageEvent>;

    private socketUrl(): string {
        return '';
    }

    private open(): void {
        const url: string = this.socketUrl();
        this.socket = new WebSocket(url);

        this.socket.onopen = (e: Event) => {

        };

        this.socket.onerror = (e: Event) => {
            this.subject.error(e);
        };

        this.socket.onclose = (e: CloseEvent) => {
            this.subject.complete();
            this.socket = undefined;
        };
    }

    send<T>(message: T): void {
        if (this.socket && this.socket.readyState == WebSocket.OPEN) {
            const data: string = JSON.stringify(message);
            this.socket.send(data);
        }
    }


}