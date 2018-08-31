import { CanvasComponent } from './canvas.component';

export interface CanvasRenderEvent {
    time: number
    delta: number
    component: CanvasComponent
    context2D: CanvasRenderingContext2D
}