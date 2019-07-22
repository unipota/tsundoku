<template lang="pug">
  .desktop-nav-button(role="button")
    .icon
      slot
    .label
      | {{ text }}
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class DesktopNavBarButton extends Vue {
  @Prop({ type: String, default: '' }) text!: string
}
</script>

<style lang="sass" scoped>
.desktop-nav-button
  display: flex
  align-items: center
  background-color: $tsundoku-red-bg
  border-radius: 100vw
  user-select: none
  cursor: pointer
.icon
  display: flex
  align-items: center
  justify-content: center
  width: 48px
  height: 48px
  margin-right: 16px
  border-radius: 100vw
  background-color: $tsundoku-red
.label
  color: $tsundoku-red
  font-weight: bold
</style>
