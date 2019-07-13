<template functional>
  <svg
    :width="props.width"
    :height="props.height"
    viewBox="0 0 46 34"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M30.9068 0.000281139C32.1306 0.0169802 33.2216 0.775517 33.6635 1.91696L37.4158 11.6098H43.0007C44.6575 11.6098 46.0007 12.9529 46.0007 14.6098C46.0007 16.2666 44.6575 17.6098 43.0007 17.6098H35.3602C34.1213 17.6098 33.0098 16.8482 32.5625 15.6928L30.7563 11.027L21.9054 32.159C21.454 33.2367 20.4179 33.9543 19.2502 33.9979C18.0826 34.0415 16.9959 33.4032 16.4653 32.3621L9.785 19.2524L5.69082 27.5473C4.9575 29.0331 3.1586 29.643 1.67286 28.9097C0.187132 28.1763 -0.422812 26.3774 0.310514 24.8917L7.05209 11.2332C7.55443 10.2154 8.58812 9.56829 9.72307 9.56104C10.858 9.55379 11.8999 10.1877 12.4152 11.1989L18.8721 23.87L28.0987 1.84103C28.5716 0.712071 29.6829 -0.0164177 30.9068 0.000281139Z"
      :fill="props.color"
    />
  </svg>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class IconToukei extends Vue {
  @Prop({ type: String, default: 'var(--toukei-black)' })
  private color!: string

  @Prop({ type: Number, default: 34 })
  private height!: number

  @Prop({ type: Number, default: 46 })
  private width!: number
}
</script>

<style lang="sass"></style>
