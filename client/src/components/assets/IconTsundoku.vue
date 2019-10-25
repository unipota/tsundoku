<template functional>
  <svg
    :width="props.width"
    :height="props.height"
    viewBox="0 0 46 40"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <rect
      :fill="props.color"
      x="2.83081"
      y="30.1384"
      width="35.3846"
      height="8.49231"
      rx="1"
    />
    <rect :fill="props.color" y="1" width="35.3846" height="8.49231" rx="1" />
    <rect
      :fill="props.color"
      x="14.1538"
      y="15.5693"
      width="31.8462"
      height="8.49231"
      rx="1"
    />
  </svg>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class IconTsundoku extends Vue {
  @Prop({ type: String, default: 'var(--tsundoku-red)' })
  private color!: string

  @Prop({ type: Number, default: 39.63 })
  private height!: number

  @Prop({ type: Number, default: 46 })
  private width!: number
}
</script>

<style lang="sass"></style>
