import { Message } from './message';

export interface ServerStatusMessage extends Message {
    Version: string
    Release: string
    Connected: boolean
    Clients: number
}