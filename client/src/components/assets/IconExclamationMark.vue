<template functional>
  <svg
    :width="props.width"
    :height="props.height"
    viewBox="0 0 7 32"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path
      :fill="props.color"
      d="M0 0H7L6.16168 22.3562H0.838323L0 0ZM0.419162 32V25.863H6.58084V32H0.419162Z"
    />
  </svg>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class IconExclamationMark extends Vue {
  @Prop({ type: String, default: 'var(--danger-red)' })
  private color!: string

  @Prop({ type: Number, default: 32 })
  private height!: number

  @Prop({ type: Number, default: 7 })
  private width!: number
}
</script>

<style lang="sass"></style>
