import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule } from '@angular/router';
import { MenubarModule } from 'primeng/menubar';
import { TabMenuModule } from 'primeng/primeng';
import {ToastModule} from 'primeng/toast';
import { ControlViewModule } from './control/control-view.module';
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
        TabMenuModule,
        ToastModule,
        // custom
        ControlViewModule,
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