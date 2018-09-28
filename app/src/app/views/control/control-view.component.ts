import { Component, OnDestroy, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { Message } from '../../communication/message/message';
import { MessageFactoryService } from '../../communication/message/message-factory.service';
import { MessageService } from '../../communication/message/message.service';
import { ServerStatus } from '../../storage/server-status/server-status';
import { ServerStatusService } from '../../storage/server-status/server-status.service';

@Component({
    templateUrl: './control-view.component.html'
})
export class ControlViewComponent implements OnInit, OnDestroy {

    serverStatus: ServerStatus;

    private subscription: Subscription;

    constructor(private serverStatusService: ServerStatusService,
                private messageService: MessageService,
                private messageFactoryService: MessageFactoryService) {
    }


    ngOnInit(): void {
        this.subscription = this.serverStatusService.on(((serverStatus: ServerStatus) => {
            this.serverStatus = serverStatus;
        }));
    }

    ngOnDestroy(): void {
        if (this.subscription) {
            this.subscription.unsubscribe();
        }
    }

    inject(): void {
        const injectMessage: Message = this.messageFactoryService.injectMessage();
        this.messageService.send(injectMessage);
    }

}
