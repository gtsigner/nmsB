import { Message } from './message';

export interface ErrorMessage extends Message {
    Error: string;
}
