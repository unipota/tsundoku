<template functional>
  <svg width="40" height="38" viewBox="0 0 40 38" fill="none" xmlns="http://www.w3.org/2000/svg">
    <rect
      y="38"
      width="38"
      height="9.3086"
      rx="1"
      transform="rotate(-90 0 38)"
      :fill="props.color"
    />
    <rect
      x="15.3457"
      y="38"
      width="38"
      height="9.3086"
      rx="1"
      transform="rotate(-90 15.3457 38)"
      :fill="props.color"
    />
    <rect
      x="30.6914"
      y="37.4216"
      width="34.2"
      height="9.3086"
      rx="1"
      transform="rotate(-90 30.6914 37.4216)"
      :fill="props.color"
    />
  </svg>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component({
  components: {},
})
export default class IconKidoku extends Vue {
  @Prop({ type: String, default: 'var(--kidoku-blue)' })
  public color!: string
}
</script>

<style lang="sass">
</style>
