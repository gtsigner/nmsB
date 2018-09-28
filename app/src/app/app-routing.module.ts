import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ControlViewComponent } from './views/control/control-view.component';
import { DebugViewComponent } from './views/debug/debug-view.component';

const routes: Routes = [
    { path: 'debug', component: DebugViewComponent },
    { path: 'control', component: ControlViewComponent }
    // { path: '', redirectTo: 'debug', pathMatch: 'full' }
];

@NgModule({
    imports: [
        RouterModule.forRoot(routes, { useHash: true })
    ],
    exports: [
        RouterModule
    ]
})

export class AppRoutingModule {

}
