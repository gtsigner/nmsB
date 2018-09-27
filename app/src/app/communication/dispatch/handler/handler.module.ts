import { NgModule } from '@angular/core';
import { DebugModule } from '../../../storage/debug/debug.module';
import { ServerStatusModule } from '../../../storage/server-status/server-status.module';
import { ErrorHandler } from './error-handler';
import { ServerStatusHandler } from './server-status-handler';

@NgModule({
    imports: [
        DebugModule,
        ServerStatusModule
    ],
    providers: [
        ErrorHandler,
        ServerStatusHandler
    ]
})
export class HandlerModule {

}