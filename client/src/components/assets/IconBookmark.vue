<template functional>
  <svg
    :width="props.width"
    :height="props.height"
    viewBox="0 0 32 32"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path
      :fill="props.color"
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M7 4C7 2.89543 7.89543 2 9 2H23C24.1046 2 25 2.89543 25 4V14V17V28.4612C25 29.1849 24.2551 29.669 23.5939 29.375L16 26L8.40614 29.375C7.74485 29.669 7 29.1849 7 28.4612V17V14V4ZM16 12C17.6569 12 19 10.6569 19 9C19 7.34315 17.6569 6 16 6C14.3431 6 13 7.34315 13 9C13 10.6569 14.3431 12 16 12Z"
    />
  </svg>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class IconBookmark extends Vue {
  @Prop({ type: String, default: 'currentColor' })
  private color!: string

  @Prop({ type: Number, default: 32 })
  private height!: number

  @Prop({ type: Number, default: 32 })
  private width!: number
}
</script>

<style lang="sass"></style>
