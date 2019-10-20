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
      d="M7.05723 2.00952C7.05723 1.45198 6.60525 1 6.0477 1C5.49016 1 5.03818 1.45198 5.03818 2.00952V8.5715L2.6153 6.75433C2.16926 6.4198 1.53649 6.5102 1.20196 6.95623C0.867438 7.40227 0.957833 8.03504 1.40387 8.36957L5.44196 11.3981C5.8009 11.6673 6.29444 11.6673 6.65338 11.3981L10.6915 8.36957C11.1375 8.03504 11.2279 7.40227 10.8934 6.95623C10.5589 6.5102 9.92608 6.4198 9.48005 6.75433L7.05723 8.57145V2.00952Z"
    />
  </svg>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class IconDownArrow2 extends Vue {
  @Prop({ type: String, default: 'var(--text-black)' })
  private color!: string

  @Prop({ type: Number, default: 12 })
  private height!: number

  @Prop({ type: Number, default: 12 })
  private width!: number
}
</script>

<style lang="sass"></style>
