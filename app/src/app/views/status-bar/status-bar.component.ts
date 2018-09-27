import { Component, OnDestroy, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { ServerStatus } from '../../storage/server-status/server-status';
import { ServerStatusService } from '../../storage/server-status/server-status.service';

@Component({
    selector: 'app-status-bar',
    templateUrl: './status-bar.component.html'
})
export class StatusBarComponent implements OnInit, OnDestroy {

    serverStatus: ServerStatus;

    private subscription: Subscription;

    constructor(private serverStatusService: ServerStatusService) {
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

}
