import { Component, ElementRef, EventEmitter, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { Subscription } from 'rxjs';
import { CanvasDebugService } from './canvas-debug.service';
import { CanvasRenderEvent } from './canvas-render-event';

@Component({
    selector: 'nms-canvas',
    templateUrl: './canvas.component.html'
})
export class CanvasComponent implements OnInit, OnDestroy {

    private static SECOND_IN_MS: number = 1000.0;

    @ViewChild('canvas')
    canvasRef: ElementRef;

    subscription: Subscription;

    onRender: EventEmitter<CanvasRenderEvent>;

    private debug: boolean;

    private fps: number;
    private fpsTimer: number;
    private fpsCounter: number;
    private lastUpdate: number;

    private context2D: CanvasRenderingContext2D;

    constructor(private canvasDebugService: CanvasDebugService) {
        this.onRender = new EventEmitter(false);
    }

    ngOnInit(): void {
        // initial timers and fps
        this.fps = 0;
        this.fpsTimer = 0;
        this.lastUpdate = Date.now();

        // subscribe to debug flag
        this.subscription = this.canvasDebugService.subscribeDebug({
            next: (debug: boolean) => {
                this.debug = debug;
            }
        });

        // get the canvas context
        const element: HTMLCanvasElement = this.canvasRef.nativeElement;
        this.context2D = element.getContext(`2d`);
    }

    private render(now: number): void {
        // calculate last frame delta
        const delta: number = now - this.lastUpdate;
        this.lastUpdate = now;

        // render internal stuff
        this.internalRender(delta);

        // create render event
        const event: CanvasRenderEvent = {
            delta,
            time: now,
            component: this,
            context2D: this.context2D
        };

        // emit event
        this.onRender.emit(event);
        // request the next frame
        this.nextFrame();
    }

    private nextFrame(): void {
        requestAnimationFrame((time: number) => {
            this.render(time);
        });
    }

    private internalRender(delta: number): void {
        // increment temp fps counter
        this.fpsCounter++;
        // add delta to fps timer
        this.fpsTimer = this.fpsTimer + delta;
        // check if timer greater then 1 second
        if (this.fpsTimer > CanvasComponent.SECOND_IN_MS) {
            // store fps counted
            this.fps = this.fpsCounter;
            // reset fps counter to zero
            this.fpsCounter = 0;
            // decrement fps timer
            this.fpsTimer = this.fpsTimer - CanvasComponent.SECOND_IN_MS;
        }

        // clean the canvas
        const width: number = this.context2D.canvas.width;
        const height: number = this.context2D.canvas.height;
        this.context2D.clearRect(0, 0, width, height);

        // check for debug
        if (this.debug) {
            // render debug stats
            this.renderDebug(delta);
        }
    }

    private renderDebug(delta: number): void {
        const offset: any = {
            x: 10,
            y: 10
        };

        this.context2D.fillStyle = 'rgb(255,255,255)';
        const width: number = this.context2D.canvas.width;

        // render fps

        const fpsString: string = `${this.fps}`;
        const fpsTextMetrics: TextMetrics = this.context2D.measureText(fpsString);

        const fpsY: number = offset.y;
        const fpsX: number = width - offset.x - fpsTextMetrics.width;

        this.context2D.fillText(fpsString, fpsX, fpsY);

        // render delta

        const deltaString: string = `${delta} ms`;
        const deltaTextMetrics: TextMetrics = this.context2D.measureText(deltaString);

        const deltaY: number = offset.y;
        const deltaX: number = fpsX - deltaTextMetrics.width;
        this.context2D.fillText(fpsString, deltaX, deltaY);
    }

    ngOnDestroy(): void {
        if (this.subscription) {
            this.subscription.unsubscribe();
        }
    }

}
