<template functional>
  <svg
    :width="props.width"
    :height="props.height"
    viewBox="0 0 24 24"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <rect :fill="props.color" x="4" y="4" width="6.85714" height="7" rx="1" />
    <rect :fill="props.color" x="4" y="13" width="6.85714" height="7" rx="1" />
    <rect
      :fill="props.color"
      x="13.1428"
      y="4"
      width="6.85714"
      height="7"
      rx="1"
    />
    <rect
      :fill="props.color"
      x="13.1428"
      y="13"
      width="6.85714"
      height="7"
      rx="1"
    />
  </svg>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class IconGrid extends Vue {
  @Prop({ type: String, default: 'currentColor' })
  private color!: string

  @Prop({ type: Number, default: 24 })
  private height!: number

  @Prop({ type: Number, default: 24 })
  private width!: number
}
</script>

<style lang="sass"></style>
