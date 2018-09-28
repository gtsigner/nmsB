import { NgModule } from '@angular/core';
import { DebugModule } from '../../../storage/debug/debug.module';
import { ServerStatusModule } from '../../../storage/server-status/server-status.module';
import { DebugHandler } from './debug-handler';
import { ServerStatusHandler } from './server-status-handler';

@NgModule({
    imports: [
        DebugModule,
        ServerStatusModule
    ],
    providers: [
        DebugHandler,
        ServerStatusHandler
    ]
})
export class HandlerModule {

}