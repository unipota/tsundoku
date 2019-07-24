<template functional>
  <svg
    :width="props.width"
    :height="props.height"
    viewBox="0 0 29 32"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M19.9294 11.7765C19.9294 16.2792 16.2792 19.9294 11.7765 19.9294C7.27372 19.9294 3.62352 16.2792 3.62352 11.7765C3.62352 7.27372 7.27372 3.62352 11.7765 3.62352C16.2792 3.62352 19.9294 7.27372 19.9294 11.7765ZM16.8961 22.3848C15.348 23.1333 13.6112 23.5529 11.7765 23.5529C5.2725 23.5529 0 18.2804 0 11.7765C0 5.2725 5.2725 0 11.7765 0C18.2804 0 23.5529 5.2725 23.5529 11.7765C23.5529 14.7161 22.4758 17.4043 20.6947 19.4677C21.0764 19.6392 21.4241 19.9017 21.703 20.2513L27.5557 27.5874C28.4917 28.7607 28.2994 30.4706 27.1261 31.4067C25.9528 32.3427 24.2429 32.1504 23.3068 30.9771L17.4542 23.641C17.1529 23.2633 16.9685 22.83 16.8961 22.3848Z"
      :fill="props.color"
    />
  </svg>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class IconLogo extends Vue {
  @Prop({ type: String, default: 'white' })
  private color!: string

  @Prop({ type: Number, default: 20 })
  private height!: number

  @Prop({ type: Number, default: 20 })
  private width!: number
}
</script>

<style lang="sass"></style>
