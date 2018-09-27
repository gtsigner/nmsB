import { MessageDirection } from './message.direction';
import { MessageType } from './message.type';

export interface Message {
    Type: MessageType
    Direction: MessageDirection
    ClientId: string
    RequestId: string
}