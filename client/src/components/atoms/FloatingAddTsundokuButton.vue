<template lang="pug">
  .floating-add-tsundoku-button.pointer
   | +積む
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

@Component({
  components: {}
})
export default class FloatingAddTsundokuButton extends Vue {}
</script>

<style lang="sass" scoped>
.floating-add-tsundoku-button
  position: fixed
  display: flex
  align-items: center
  right: 0
  bottom: 94px + 18px
  width: 94px
  height: 60px
  padding:
    left: 16px
  background-color: var(--tsundoku-red)
  border:
    radius: 100vw 0 0 100vw
  color: var(--text-white)
  font:
    size: 24px
    weight: bold
  box-sizing: border-box
</style>
