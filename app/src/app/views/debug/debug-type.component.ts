import { Component, Input } from '@angular/core';
import { DebugEntryType } from '../../storage/debug/debug-entry-type';

@Component({
    selector: 'app-debug-type',
    templateUrl: './debug-type.component.html'
})
export class DebugTypeComponent {

    @Input()
    type: DebugEntryType;

    DebugEntryType = DebugEntryType;

    constructor() {
    }

}
