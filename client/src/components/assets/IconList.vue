<template functional>
  <svg
    :width="props.width"
    :height="props.height"
    viewBox="0 0 24 24"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <rect :fill="props.color" x="3" y="4" width="4" height="4" rx="1" />
    <rect :fill="props.color" x="9" y="4" width="13" height="4" rx="1" />
    <rect :fill="props.color" x="9" y="10" width="13" height="4" rx="1" />
    <rect :fill="props.color" x="9" y="16" width="13" height="4" rx="1" />
    <rect :fill="props.color" x="3" y="10" width="4" height="4" rx="1" />
    <rect :fill="props.color" x="3" y="16" width="4" height="4" rx="1" />
  </svg>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class IconList extends Vue {
  @Prop({ type: String, default: 'currentColor' })
  private color!: string

  @Prop({ type: Number, default: 24 })
  private height!: number

  @Prop({ type: Number, default: 24 })
  private width!: number
}
</script>

<style lang="sass"></style>
