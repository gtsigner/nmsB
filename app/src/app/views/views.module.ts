import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule } from '@angular/router';
import { MenubarModule } from 'primeng/menubar';
import { DebugViewModule } from './debug/debug-view.module';
import { StatusBarModule } from './status-bar/status-bar.module';
import { ViewsComponent } from './views.component';

@NgModule({
    imports: [
        // angular
        RouterModule,
        BrowserModule,
        // primeng
        MenubarModule,
        // custom
        DebugViewModule,
        StatusBarModule
    ],
    declarations: [
        ViewsComponent
    ],
    exports: [
        ViewsComponent
    ]
})
export class ViewsModule {

}