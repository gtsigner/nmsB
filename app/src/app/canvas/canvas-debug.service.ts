import { Injectable } from '@angular/core';
import { BehaviorSubject, PartialObserver, Subscription } from 'rxjs';

@Injectable()
export class CanvasDebugService {

    private debugSubject: BehaviorSubject<boolean>;

    constructor() {
        this.debugSubject = new BehaviorSubject<boolean>(false);
    }

    setDebug(b: boolean) {
        this.debugSubject.next(b);
    }

    subscribeDebug(observer?: PartialObserver<boolean>): Subscription {
        return this.debugSubject.subscribe(observer);
    }

}
