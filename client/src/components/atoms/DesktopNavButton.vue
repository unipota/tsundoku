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
  // background-color: $tsundoku-red-bg
  background-color: var(--tsundoku-red)
  border-radius: 100vw
  user-select: none
  cursor: pointer
  transition: box-shadow .5s
  box-shadow: 0px 0px 5px -1px #e3402a6e
  &:hover
    box-shadow: 0px 3px 5px -1px #e3402a6e
.icon
  display: flex
  align-items: center
  justify-content: center
  width: 48px
  height: 48px
  margin-right: 16px
  border-radius: 100vw
  background-color: var(--tsundoku-red)
.label
  // color: var(--tsundoku-red)
  color: white
  font-weight: bold
</style>
