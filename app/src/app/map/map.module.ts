import { NgModule } from '@angular/core';
import { CanvasModule } from '../canvas/canvas.module';
import { MapComponent } from './map.component';

@NgModule({
    imports: [
        CanvasModule
    ],
    declarations: [
        MapComponent
    ]
})
export class MapModule {

}