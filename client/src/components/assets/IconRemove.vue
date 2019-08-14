<template functional>
  <svg
    :width="props.width"
    :height="props.height"
    viewBox="0 0 25 29"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M8.98496 0C8.01428 0 7.22738 0.786894 7.22738 1.75758V3.51515H2.39405C1.42337 3.51515 0.636475 4.30204 0.636475 5.27273C0.636475 6.09199 1.19702 6.78034 1.95551 6.97515C1.95494 6.99347 1.95466 7.01185 1.95466 7.0303V27.2424C1.95466 28.2131 2.74155 29 3.71223 29H21.288C22.2587 29 23.0456 28.2131 23.0456 27.2424V7.0303C23.0456 7.01185 23.0453 6.99347 23.0447 6.97515C23.8032 6.78034 24.3637 6.09199 24.3637 5.27273C24.3637 4.30204 23.5769 3.51515 22.6062 3.51515H17.7728V1.75758C17.7728 0.786894 16.9859 0 16.0153 0H8.98496ZM5.46981 7.0303H19.5304V25.4848H5.46981V7.0303ZM8.98496 12.303C9.95564 12.303 10.7425 13.0899 10.7425 14.0606V21.9697C10.7425 22.9404 9.95564 23.7273 8.98496 23.7273C8.01428 23.7273 7.22738 22.9404 7.22738 21.9697V14.0606C7.22738 13.0899 8.01428 12.303 8.98496 12.303ZM17.7728 14.0606C17.7728 13.0899 16.9859 12.303 16.0153 12.303C15.0446 12.303 14.2577 13.0899 14.2577 14.0606V21.9697C14.2577 22.9404 15.0446 23.7273 16.0153 23.7273C16.9859 23.7273 17.7728 22.9404 17.7728 21.9697V14.0606Z"
      :fill="props.color"
    />
  </svg>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class IconRemove extends Vue {
  @Prop({ type: String, default: 'var(--text-gray)' })
  private color!: string

  @Prop({ type: Number, default: 25 })
  private height!: number

  @Prop({ type: Number, default: 29 })
  private width!: number
}
</script>

<style lang="sass"></style>
