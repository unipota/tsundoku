<template functional>
  <svg
    :width="props.width"
    :height="props.height"
    viewBox="0 0 13 21"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path
      d="M9.10705 19.9352C9.89967 20.7045 11.1659 20.6856 11.9352 19.893C12.7045 19.1003 12.6856 17.8341 11.893 17.0648L9.10705 19.9352ZM11.893 3.43516C12.6856 2.66586 12.7045 1.39967 11.9352 0.60705C11.1659 -0.185568 9.89967 -0.204468 9.10705 0.564838L11.893 3.43516ZM2 10.25L0.607049 8.81484C0.218993 9.19148 1.04829e-07 9.70922 9.83803e-08 10.25C9.19315e-08 10.7908 0.218993 11.3085 0.607049 11.6852L2 10.25ZM11.893 17.0648L3.39295 8.81484L0.607049 11.6852L9.10705 19.9352L11.893 17.0648ZM3.39295 11.6852L11.893 3.43516L9.10705 0.564838L0.607049 8.81484L3.39295 11.6852Z"
      :fill="props.color"
    />
  </svg>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class IconBack extends Vue {
  @Prop({ type: String, default: 'var(--border-gray)' })
  private color!: string

  @Prop({ type: Number, default: 21 })
  private height!: number

  @Prop({ type: Number, default: 13 })
  private width!: number
}
</script>

<style lang="sass"></style>
