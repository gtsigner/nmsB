import { DebugEntryType } from './debug-entry-type';

export interface DebugEntry {
    type: DebugEntryType;
    date: Date;
    message: string;
}