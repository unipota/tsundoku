declare module 'v-lazy-image' {
  import { VueConstructor } from 'vue'

  export const VLazyImage: VLazyImage

  export interface VLazyImageProps {
    src: string
    srcset: string | undefined
    intersectionOptions: {} | undefined
    usePicture: boolean
  }

  export interface VLazyImageData {
    observer: null
    intersected: false
    loaded: false
  }

  export interface VLazyImageMethods {
    load: () => {}
  }

  export interface VLazyImageComputed {
    srcImage: () => string
    srcsetImage: () => string | boolean
  }

  export interface VLazyImage extends VueConstructor {
    props: VLazyImageProps
    data: () => VLazyImageData
    methods: VLazyImageMethods
    computed: VLazyImageComputed
  }

  export default VLazyImage
}
