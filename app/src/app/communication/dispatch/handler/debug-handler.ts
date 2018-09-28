import { Injectable } from '@angular/core';
import { DebugEntryType } from '../../../storage/debug/debug-entry-type';
import { DebugService } from '../../../storage/debug/debug.service';
import { DebugMessage } from '../../message/debug-message';
import { Handler } from './handler';

@Injectable()
export class DebugHandler implements Handler<DebugMessage> {

    constructor(private debugService: DebugService) {
    }

    handle(message: DebugMessage): void {
        const type: DebugEntryType = message.DebugType as DebugEntryType;
        this.debugService.notify(message.Text, type);
    }

}
