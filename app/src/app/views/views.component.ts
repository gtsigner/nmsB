import { Component, OnInit } from '@angular/core';
import { MenuItem } from 'primeng/api';

@Component({
    selector: 'app-views',
    templateUrl: './views.component.html'
})
export class ViewsComponent implements OnInit {

    items: MenuItem[];

    constructor() {

    }

    ngOnInit(): void {

        this.items = [
            {
                label: 'Debug',
                routerLink: ['debug']
            },
            {
                label: 'Map'
            },
            {
                label: 'Bot'
            },
            {
                label: 'Control',
                routerLink: ['control']
            }
        ];
    }


}
