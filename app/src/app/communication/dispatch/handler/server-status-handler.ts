import { Injectable } from '@angular/core';
import { ServerStatus } from '../../../storage/server-status/server-status';
import { ServerStatusService } from '../../../storage/server-status/server-status.service';
import { ServerStatusMessage } from '../../message/server-status-message';
import { Handler } from './handler';

@Injectable()
export class ServerStatusHandler implements Handler<ServerStatusMessage> {

    constructor(private serverStatusService: ServerStatusService) {
    }

    handle(message: ServerStatusMessage): void {
        const serverStatus: ServerStatus = this.toServerStatus(message);
        this.serverStatusService.update(serverStatus);
    }

    private toServerStatus(message: ServerStatusMessage): ServerStatus {
        return {
            clients: message.Clients,
            release: message.Release,
            version: message.Version,
            connected: message.Connected
        } as ServerStatus;
    }

}
