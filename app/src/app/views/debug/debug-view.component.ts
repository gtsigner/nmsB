import { Component, OnDestroy, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { DebugEntry } from '../../storage/debug/debug-entry';
import { DebugService } from '../../storage/debug/debug.service';

@Component({
    templateUrl: './debug-view.component.html'
})
export class DebugViewComponent implements OnInit, OnDestroy {

    entities: DebugEntry[];

    private subscription: Subscription;

    constructor(private debugService: DebugService) {
    }

    ngOnInit(): void {
        this.subscription = this.debugService.on(((entities: DebugEntry[]) => {
            this.entities = entities;
        }));
    }


    ngOnDestroy(): void {
        if (this.subscription) {
            this.subscription.unsubscribe();
        }
    }

}
