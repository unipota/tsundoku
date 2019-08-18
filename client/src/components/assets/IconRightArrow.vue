<template functional>
  <svg
    :width="props.width"
    :height="props.height"
    viewBox="0 0 13 21"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M12.5 10.25C12.5 10.7908 12.281 11.3085 11.893 11.6852L3.39295 19.9352C2.60033 20.7045 1.33414 20.6856 0.564838 19.893C-0.204468 19.1003 -0.18557 17.8341 0.607048 17.0648L7.6284 10.25L0.607049 3.43516C-0.185569 2.66586 -0.204467 1.39967 0.564838 0.60705C1.33414 -0.185568 2.60033 -0.204466 3.39295 0.56484L11.893 8.81484C12.281 9.19148 12.5 9.70922 12.5 10.25Z"
      :fill="props.color"
    />
  </svg>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class IconRightArrow extends Vue {
  @Prop({ type: String, default: 'currentColor' })
  private color!: string

  @Prop({ type: Number, default: 21 })
  private height!: number

  @Prop({ type: Number, default: 13 })
  private width!: number
}
</script>

<style lang="sass"></style>
