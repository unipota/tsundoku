<template functional>
  <svg
    :width="props.width"
    :height="props.height"
    viewBox="0 0 24 24"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M22.8221 4.10738C23.7617 3.16776 23.7617 1.64433 22.8221 0.704714C21.8824 -0.234906 20.359 -0.234904 19.4194 0.704716L11.7634 8.3607L4.5806 1.1779C3.64098 0.23828 2.11755 0.23828 1.17793 1.1779C0.238308 2.11752 0.238309 3.64095 1.17793 4.58057L8.36074 11.7634L0.704716 19.4194C-0.234905 20.359 -0.234905 21.8824 0.704716 22.822C1.64434 23.7617 3.16777 23.7617 4.10739 22.822L11.7634 15.166L19.8926 23.2952C20.8322 24.2348 22.3557 24.2348 23.2953 23.2952C24.2349 22.3556 24.2349 20.8322 23.2953 19.8926L15.1661 11.7634L22.8221 4.10738Z"
      :fill="props.color"
    />
  </svg>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class IconClose extends Vue {
  @Prop({ type: String, default: 'white' })
  private color!: string

  @Prop({ type: Number, default: 24 })
  private height!: number

  @Prop({ type: Number, default: 24 })
  private width!: number
}
</script>

<style lang="sass"></style>
