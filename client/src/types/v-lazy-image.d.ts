declare module 'v-lazy-image' {
  var VLazyImage: {
    props: {
      src: string,
      srcset: string | undefined,
      intersectionOptions: {} | undefined,
      usePicture: boolean,
    },
    inheritAttrs: boolean,
    data: () => {
      observer: null,
      intersected: false,
      loaded: false
    },
    computed: {
      srcImage: () => string,
      srcsetImage: () => string | boolean,
    },
    methods: {
      load: VoidFunction,
    },
    render: (h: object) => any,
    mounted: VoidFunction,
    destroyed: VoidFunction,
  }
  export default VLazyImage
}