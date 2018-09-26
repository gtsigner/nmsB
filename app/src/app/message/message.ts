import {MessageType} from "./message.type";
import {MessageDirection} from "./message.direction";

export interface Message {
    Type: MessageType
    Direction: MessageDirection
    ClientId: string
    RequestId: string
}