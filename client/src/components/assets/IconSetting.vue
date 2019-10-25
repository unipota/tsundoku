<template functional>
  <svg
    :width="props.width"
    :height="props.height"
    :fill="props.color"
    viewBox="0 0 36 36"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path
      :fill="props.color"
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M32 26.2487L32 10.0829L18 2L4.00007 10.0829L4.00007 26.2487L18 34.3315L32 26.2487ZM11.414 18.1659C11.414 21.8032 14.3626 24.7519 18 24.7519C21.6374 24.7519 24.5861 21.8032 24.5861 18.1659C24.5861 14.5285 21.6374 11.5798 18 11.5798C14.3626 11.5798 11.414 14.5285 11.414 18.1659Z"
    />
  </svg>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class IconSetting extends Vue {
  @Prop({ type: String, default: 'var(--text-gray)' })
  private color!: string

  @Prop({ type: Number, default: 36 })
  private height!: number

  @Prop({ type: Number, default: 36 })
  private width!: number
}
</script>

<style lang="sass"></style>
