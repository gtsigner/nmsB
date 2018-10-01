import { Injectable } from '@angular/core';
import { BehaviorSubject, Subscription } from 'rxjs';
import { DebugEntry } from './debug-entry';
import { DebugEntryType } from './debug-entry-type';
import { MessageService } from 'primeng/api';

@Injectable()
export class DebugService {

    private _count: number;
    private subject: BehaviorSubject<DebugEntry[]>;

    constructor(private messageService:MessageService) {
        this._count = 100;
        this.subject = new BehaviorSubject([]);
    }

    error(e: Error): void {
        this.notify(JSON.stringify(e), DebugEntryType.ERROR);
    }

    notify(message: string, type: DebugEntryType, date?: Date): void {
        if (!date) {
            date = new Date();
        }
        const entry:DebugEntry = { date, message, type } as DebugEntry
        this.append(entry);
        this.pushMessage(entry)
    }

    private pushMessage(entry:DebugEntry): void {
        if(entry.type === DebugEntryType.ERROR){
            this.messageService.add({
                severity: 'error',
                summary:'Error',
                detail: entry.message
            })
        }else if(entry.type === DebugEntryType.INFO){
            this.messageService.add({
                severity: 'info',
                summary:'Info',
                detail: entry.message
            })
        }
    }

    on(listener: (entities: DebugEntry[]) => void): Subscription {
        return this.subject.subscribe(listener);
    }

    private append(entry: DebugEntry): void {
        const list: DebugEntry[] = this.subject.getValue();
        list.push(entry);
        while (list.length > this._count) {
            list.shift();
        }
        this.subject.next(list);
    }

    set count(n: number) {
        this._count = n;
        const list: DebugEntry[] = this.subject.getValue();
        while (list.length > this._count) {
            list.shift();
        }
        this.subject.next(list);
    }

    get count(): number {
        return this._count;
    }

}
