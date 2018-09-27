import { Injectable } from '@angular/core';
import { DebugEntryType } from '../../../storage/debug/debug-entry-type';
import { DebugService } from '../../../storage/debug/debug.service';
import { ErrorMessage } from '../../message/error-message';
import { Handler } from './handler';

@Injectable()
export class ErrorHandler implements Handler<ErrorMessage> {

    constructor(private debugService: DebugService) {
    }

    handle(message: ErrorMessage): void {
        this.debugService.notify(message.Error, DebugEntryType.ERROR);
    }

}
