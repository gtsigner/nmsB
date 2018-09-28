import { Message } from './message';

export interface DebugMessage extends Message {
    Text: string;
    DebugType: string;
}
