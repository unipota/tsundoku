<template functional>
  <svg
    :width="props.width"
    :height="props.height"
    viewBox="0 0 101 144"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <rect width="101" height="144" rx="10" fill="#F5F5F5" />
    <g opacity="0.4">
      <rect
        x="27"
        y="77.186"
        width="33.8642"
        height="8.35193"
        rx="1"
        fill="#4B4B4B"
      />
      <rect
        width="38.1899"
        height="8.50219"
        rx="1"
        transform="matrix(0.274045 0.961717 -0.963631 0.267238 64.5342 47)"
        fill="#4B4B4B"
      />
      <rect
        x="25.7905"
        y="62.8684"
        width="30.2359"
        height="8.35193"
        rx="1"
        fill="#4B4B4B"
      />
    </g>
  </svg>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class DummyBookCover extends Vue {
  @Prop({ type: Number, default: 144 })
  private height!: number

  @Prop({ type: Number, default: 101 })
  private width!: number
}
</script>

<style lang="sass"></style>
