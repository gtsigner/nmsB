import { Injectable } from '@angular/core';
import { BehaviorSubject, Subscription } from 'rxjs';
import { ServerStatus } from './server-status';

@Injectable()
export class ServerStatusService {

    private subject: BehaviorSubject<ServerStatus>;

    constructor() {
        this.subject = new BehaviorSubject(undefined);
    }

    update(serverStatus: ServerStatus): void {
        console.log(serverStatus);
        this.subject.next(serverStatus);
    }

    on(listener: (serverStatus: ServerStatus) => void): Subscription {
        return this.subject.subscribe(listener);
    }

}
