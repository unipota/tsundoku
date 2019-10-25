<template functional>
  <svg
    :width="props.width"
    :height="props.height"
    viewBox="0 0 26 21"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path
      :fill="props.color"
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M25.0931 0.72895C26.229 1.76658 26.3087 3.52859 25.2711 4.66452L11.6996 19.5217C11.1953 20.0737 10.4903 20.3999 9.74302 20.4268C8.99577 20.4537 8.26909 20.179 7.72642 19.6646L0.869276 13.1646C-0.247301 12.1062 -0.294445 10.343 0.763977 9.22643C1.8224 8.10985 3.58559 8.06271 4.70216 9.12113L9.49988 13.669L21.1575 0.906941C22.1951 -0.228988 23.9571 -0.308677 25.0931 0.72895Z"
    />
  </svg>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class IconKidoku extends Vue {
  @Prop({ type: String, default: 'var(--kidoku-blue)' })
  private color!: string

  @Prop({ type: Number, default: 21 })
  private height!: number

  @Prop({ type: Number, default: 26 })
  private width!: number
}
</script>

<style lang="sass"></style>
