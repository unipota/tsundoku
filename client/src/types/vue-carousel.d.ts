declare module 'vue-carousel' {
  import { VueConstructor } from 'vue'

  export const Carousel: Carousel
  export const Slide: Slide

  interface CarouselProps {
    /**
     *  Adjust the height of the carousel for the current slide
     */
    adjustableHeight: boolean,
    /**
     * Slide transition easing for adjustableHeight
     * Any valid CSS transition easing accepted
     */
    adjustableHeightEasing: string,
    /**
     *  Center images when the size is less than the container width
     */
    centerMode: boolean,
    /**
     * Slide transition easing
     * Any valid CSS transition easing accepted
     */
    easing: string,
    /**
     * Flag to make the carousel loop around when it reaches the end
     */
    loop: boolean,
    /**
     * Minimum distance for the swipe to trigger
     * a slide advance
     */
    minSwipeDistance: number,
    /**
     * Flag to toggle mouse dragging
     */
    mouseDrag: boolean,
    /**
     * Flag to toggle touch dragging
     */
    touchDrag: boolean,
    /**
     * Listen for an external navigation request using this prop.
     */
    navigateTo: number,
    /**
     * Amount of padding to apply around the label in pixels
     */
    navigationClickTargetSize: number,
    /**
     * Flag to render the navigation component
     * (next/prev buttons)
     */
    navigationEnabled: boolean,
    /**
     * Text content of the navigation next button
     */
    navigationNextLabel: string,
    /**
     * Text content of the navigation prev button
     */
    navigationPrevLabel: string,
    /**
     * The fill color of the active pagination dot
     * Any valid CSS color is accepted
     */
    paginationActiveColor: string,
    /**
     * The fill color of pagination dots
     * Any valid CSS color is accepted
     */
    paginationColor: string,
    /**
     * Flag to render pagination component
     */
    paginationEnabled: boolean,
    /**
     * The padding inside each pagination dot
     * Pixel values are accepted
     */
    paginationPadding: number,
    /**
     * Configure the position for the pagination component.
     * The possible values are: 'bottom', 'top', 'bottom-overlay' and 'top-overlay'
     */
    paginationPosition: string,
    /**
     * The size of each pagination dot
     * Pixel values are accepted
     */
    paginationSize: number,
    /**
     * Maximum number of slides displayed on each page
     */
    perPage: number,
    /**
     * Configure the number of visible slides with a particular browser width.
     * This will be an array of arrays, ex. [[320, 2], [1199, 4]]
     * Formatted as [x, y] where x=browser width, and y=number of slides displayed.
     * ex. [1199, 4] means if (window <= 1199) then show 4 slides per page
     */
    perPageCustom: [number, number][],
    /**
     * Resistance coefficient to dragging on the edge of the carousel
     * This dictates the effect of the pull as you move towards the boundaries
     */
    resistanceCoef: number,
    /**
     * Scroll per page, not per item
     */
    scrollPerPage: boolean,
    /**
     *  Space padding option adds left and right padding style (in pixels) onto VueCarousel-inner.
     */
    spacePadding: number,
    /**
     *  Specify by how much should the space padding value be multiplied of, to re-arange the final slide padding.
     */
    spacePaddingMaxOffsetFactor: number,
    /**
     * Slide transition speed
     * Number of milliseconds accepted
     */
    speed: number,
    /**
     * Name (tag) of slide component
     * Overwrite when extending slide component
     */
    tagName: string,
    /**
     * Support for v-model functionality
     */
    value: number,
    /**
     * Support Max pagination dot amount
     */
    maxPaginationDotCount: number,
    /**
     * Support right to left
     */
    rtl: boolean
  }

  export interface Carousel extends VueConstructor {
    props: CarouselProps
  }

  export interface Slide extends VueConstructor {
    props: { title: string }
  }
}
