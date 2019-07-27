<template functional>
  <svg
    :width="props.width"
    :height="props.height"
    viewBox="0 0 38 38"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path
      d="M23.1281 4.58247C23.3138 4.26078 23.7251 4.15056 24.0468 4.33629L29.8715 7.69919C30.1932 7.88492 30.3035 8.29626 30.1177 8.61796L19.3564 27.2571C19.1707 27.5787 18.7594 27.689 18.4377 27.5032L12.613 24.1403C12.2913 23.9546 12.181 23.5433 12.3668 23.2216L23.1281 4.58247Z"
      :fill="props.color"
    />
    <path
      d="M11.7781 33.4573C11.3373 33.6954 10.8003 33.3854 10.7861 32.8847L10.5991 26.2973C10.5841 25.7712 11.1519 25.4326 11.6077 25.6957L17.5931 29.1514C18.0489 29.4145 18.0395 30.0755 17.5764 30.3256L11.7781 33.4573Z"
      :fill="props.color"
    />
  </svg>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class IconPen extends Vue {
  @Prop({ type: String, default: 'var(--text-black)' })
  private color!: string

  @Prop({ type: Number, default: 30 })
  private height!: number

  @Prop({ type: Number, default: 30 })
  private width!: number
}
</script>

<style lang="sass"></style>
