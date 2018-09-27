import { Message } from './message';

export type MessageCallback<T extends Message> = (message: T) => void;
