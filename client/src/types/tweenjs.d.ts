declare namespace TWEEN {
  export function getAll(): Tween[]
  export function removeAll(): void
  export function add(tween: Tween): void
  export function remove(tween: Tween): void
  export function update(time?: number, preserve?: boolean): boolean
  export function now(): number

  export class Tween {
    public constructor(object?: any, group?: Group)
    public getId(): number
    public isPlaying(): boolean
    public to(properties: any, duration: number): Tween
    public start(time?: number): Tween
    public stop(): Tween
    public end(): Tween
    public stopChainedTweens(): Tween
    public group(group: Group): Tween
    public delay(amount: number): Tween
    public repeat(times: number): Tween
    public repeatDelay(times: number): Tween
    public yoyo(enable: boolean): Tween
    public easing(easing: (k: number) => number): Tween
    public interpolation(
      interpolation: (v: number[], k: number) => number
    ): Tween
    public chain(...tweens: Tween[]): Tween
    public onStart(callback: (object?: any) => void): Tween
    public onStop(callback: (object?: any) => void): Tween
    public onUpdate(callback: (object?: any) => void): Tween
    public onComplete(callback: (object?: any) => void): Tween
    public update(time: number): boolean
  }

  export class Group {
    public constructor()
    public getAll(): Tween[]
    public removeAll(): void
    public add(tween: Tween): void
    public remove(tween: Tween): void
    public update(time?: number, preserve?: boolean): boolean
  }

  export var Easing: Easing
  export var Interpolation: Interpolation
}

interface Easing {
  Linear: {
    None(k: number): number
  }
  Quadratic: {
    In(k: number): number
    Out(k: number): number
    InOut(k: number): number
  }
  Cubic: {
    In(k: number): number
    Out(k: number): number
    InOut(k: number): number
  }
  Quartic: {
    In(k: number): number
    Out(k: number): number
    InOut(k: number): number
  }
  Quintic: {
    In(k: number): number
    Out(k: number): number
    InOut(k: number): number
  }
  Sinusoidal: {
    In(k: number): number
    Out(k: number): number
    InOut(k: number): number
  }
  Exponential: {
    In(k: number): number
    Out(k: number): number
    InOut(k: number): number
  }
  Circular: {
    In(k: number): number
    Out(k: number): number
    InOut(k: number): number
  }
  Elastic: {
    In(k: number): number
    Out(k: number): number
    InOut(k: number): number
  }
  Back: {
    In(k: number): number
    Out(k: number): number
    InOut(k: number): number
  }
  Bounce: {
    In(k: number): number
    Out(k: number): number
    InOut(k: number): number
  }
}

interface Interpolation {
  Linear(v: number[], k: number): number
  Bezier(v: number[], k: number): number
  CatmullRom(v: number[], k: number): number

  Utils: {
    Linear(p0: number, p1: number, t: number): number
    Bernstein(n: number, i: number): number
    Factorial(n: number): number
    CatmullRom(
      p0: number,
      p1: number,
      p2: number,
      p3: number,
      t: number
    ): number
  }
}

declare module '@tweenjs/tween.js' {
  export = TWEEN
}
