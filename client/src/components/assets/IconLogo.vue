<template functional>
  <svg
    width="41"
    :height="props.height"
    viewBox="0 0 41 33"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <rect x="1" y="25.2998" width="28" height="7" rx="1" :fill="props.color" />
    <rect
      x="32.0345"
      width="31.9759"
      height="7.03679"
      rx="1"
      transform="rotate(74.2986 32.0345 0)"
      :fill="props.color"
    />
    <rect y="13.2998" width="25" height="7" rx="1" :fill="props.color" />
  </svg>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component({
  components: {},
})
export default class IconLogo extends Vue {
  @Prop({ type: String, default: 'var(--toukei-black)' })
  private color!: string

  @Prop({ type: Number, default: 33 })
  private height!: number
}
</script>

<style lang="sass">
</style>
