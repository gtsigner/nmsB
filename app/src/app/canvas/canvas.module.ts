import { NgModule } from '@angular/core';
import { CanvasDebugService } from './canvas-debug.service';
import { CanvasComponent } from './canvas.component';

@NgModule({
    declarations: [
        CanvasComponent
    ],
    exports: [
        CanvasComponent
    ],
    providers: [
        CanvasDebugService
    ]
})
export class CanvasModule {

}