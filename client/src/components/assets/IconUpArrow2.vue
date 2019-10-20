<template functional>
  <svg
    :width="props.width"
    :height="props.height"
    viewBox="0 0 12 12"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path
      :fill="props.color"
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M5.03823 10.5906C5.03823 11.1481 5.49021 11.6001 6.04775 11.6001C6.6053 11.6001 7.05728 11.1481 7.05728 10.5906L7.05728 4.0286L9.48016 5.84577C9.9262 6.18029 10.559 6.0899 10.8935 5.64386C11.228 5.19783 11.1376 4.56506 10.6916 4.23053L6.6535 1.20196C6.29456 0.932759 5.80102 0.932759 5.44208 1.20196L1.40398 4.23053C0.957949 4.56506 0.867554 5.19783 1.20208 5.64386C1.53661 6.0899 2.16938 6.18029 2.61541 5.84577L5.03823 4.02865L5.03823 10.5906Z"
    />
  </svg>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class IconUpArrow2 extends Vue {
  @Prop({ type: String, default: 'var(--text-black)' })
  private color!: string

  @Prop({ type: Number, default: 12 })
  private height!: number

  @Prop({ type: Number, default: 12 })
  private width!: number
}
</script>

<style lang="sass"></style>
