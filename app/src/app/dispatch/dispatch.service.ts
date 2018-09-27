import { Subscription } from 'rxjs';
import { Message } from '../message/message';
import { MessageType } from '../message/message.type';
import { WebSocketService } from '../websocket/websocket.service';

export class DispatchService {

    private subscription: Subscription;

    constructor(private webSocketService: WebSocketService) {

    }

    private init(): void {
        this.webSocketService.onOpen(() => {
            this.subscribe();
        });
    }

    private subscribe(): void {
        this.subscription = this.webSocketService.subscribe((message: MessageEvent) => {
            this.dispatch(message.data);
        }, (error: Event) => {
            this.error(error);
        }, () => {
            if (this.subscription) {
                this.subscription.unsubscribe();
            }
        });
    }

    private error(error: Event): void {
// TODO Handle error
    }

    private dispatch(data: string): void {
        if (!data) {
            // TODO Handle error
            return;
        }
        const message: Message = JSON.parse(data);

        if (message.Type === MessageType.ERROR) {

        } else if (message.Type === MessageType.SERVER_STATUS) {

        }

        // TODO Handle error

    }

}

