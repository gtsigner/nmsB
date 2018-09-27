import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { DebugEntry } from './debug-entry';
import { DebugEntryType } from './debug-entry-type';

@Injectable()
export class DebugService {

    private _count: number;
    private subject: BehaviorSubject<DebugEntry[]>;

    constructor() {
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
        this.append({ date, message, type } as DebugEntry);
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
