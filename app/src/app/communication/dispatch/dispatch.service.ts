import { Injectable } from '@angular/core';
import { Subscription } from 'rxjs';
import { DebugEntryType } from '../../storage/debug/debug-entry-type';
import { DebugService } from '../../storage/debug/debug.service';
import { MessageType } from '../message/message.type';
import { WebSocketService } from '../websocket/websocket.service';
import { DebugHandler } from './handler/debug-handler';
import { ServerStatusHandler } from './handler/server-status-handler';

@Injectable()
export class DispatchService {

    private subscription: Subscription;

    constructor(private debugService: DebugService,
                private debugHandler: DebugHandler,
                private webSocketService: WebSocketService,
                private serverStatusHandler: ServerStatusHandler) {

    }

    register(): void {
        this.webSocketService.onConnected(() => {
            this.subscribe();
        });
    }

    private subscribe(): void {
        const complete = () => {
            if (this.subscription) {
                this.subscription.unsubscribe();
            }
        };
        complete();

        this.subscription = this.webSocketService.subscribe((message: any) => {
            this.dispatch(message);
        }, (error: Event) => {
            this.error(error);
        }, () => {
            complete();
        });
    }

    private error(error: Event): void {
        console.error(error);
        this.debugService.notify(JSON.stringify(error), DebugEntryType.ERROR);
    }

    private dispatch(message: any): void {
        if (!message) {
            this.debugService.error(new Error(`dispatcher received nil or empty message`));
            return;
        }

        console.log(message);

        if (message.Type === MessageType.DEBUG) {
            return this.debugHandler.handle(message);
        } else if (message.Type === MessageType.SERVER_STATUS) {
            return this.serverStatusHandler.handle(message);
        }

    }

}
