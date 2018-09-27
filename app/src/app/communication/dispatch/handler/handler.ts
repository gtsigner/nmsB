import { Message } from '../../message/message';

export interface Handler<T extends Message> {
    handle(message: T): void;
}
