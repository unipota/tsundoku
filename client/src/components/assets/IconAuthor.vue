<template functional>
  <svg
    :width="props.width"
    :height="props.height"
    viewBox="0 0 10 12"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <circle :fill="props.color" cx="5" cy="3" r="3" />
    <path
      :fill="props.color"
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M10 12C10 9.23858 7.76142 7 5 7C2.23858 7 0 9.23858 0 12H10Z"
    />
  </svg>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class IconLogo extends Vue {
  @Prop({ type: String, default: 'var(--text-gray)' })
  private color!: string

  @Prop({ type: Number, default: 12 })
  private height!: number

  @Prop({ type: Number, default: 10 })
  private width!: number
}
</script>

<style lang="sass"></style>
