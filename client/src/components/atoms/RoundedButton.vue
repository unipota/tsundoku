<template lang="pug">
  .button(@click="emitClick")
    slot
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

@Component({})
export default class RoundedButton extends Vue {
  emitClick() {
    this.$emit('click')
  }
}
</script>

<style lang="sass" scoped>
.button
  background: var(--border-gray)
  color: var(--text-black)
  border-radius: 999vw
  min-width: 200px
  text-align: center
  font:
    weight: bold
  padding: 8px
  cursor: pointer
  user-select: none
  transition: background .3s

  &:hover
   background: var(--border-gray-hovered)
</style>
