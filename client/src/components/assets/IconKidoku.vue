<template functional>
  <svg
    :width="props.width"
    :height="props.height"
    viewBox="0 0 40 38"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <rect
      :fill="props.color"
      y="38"
      width="38"
      height="9.3086"
      rx="1"
      transform="rotate(-90 0 38)"
    />
    <rect
      :fill="props.color"
      x="15.3457"
      y="38"
      width="38"
      height="9.3086"
      rx="1"
      transform="rotate(-90 15.3457 38)"
    />
    <rect
      :fill="props.color"
      x="30.6914"
      y="37.4216"
      width="34.2"
      height="9.3086"
      rx="1"
      transform="rotate(-90 30.6914 37.4216)"
    />
  </svg>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class IconKidoku extends Vue {
  @Prop({ type: String, default: 'var(--kidoku-blue)' })
  private color!: string

  @Prop({ type: Number, default: 40 })
  private height!: number

  @Prop({ type: Number, default: 38 })
  private width!: number
}
</script>

<style lang="sass"></style>
