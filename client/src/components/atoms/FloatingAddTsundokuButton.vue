<template lang="pug">
  transition-group.floating-add-tsundoku-button(:class="{'is-active': active}" name="transition-button" tag="div")
    .button-search(key="search" v-if="active")
      icon-search
    .button-scan(key="scan" v-if="active")
      icon-scanner
    .button-open(key="open" @click="handleClick")
      transition-group(name="transition-label" tag="div")
        .icon-plus(key="icon")
          icon-plus(color="var(--text-white)")
        span.button-label(key="label" v-if="!active")
          | 積む
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import IconPlus from '@/components/assets/IconPlus.vue'
import IconSearch from '@/components/assets/IconSearch.vue'
import IconScanner from '@/components/assets/IconScanner.vue'

@Component({
  components: { IconPlus, IconSearch, IconScanner }
})
export default class FloatingAddTsundokuButton extends Vue {
  active = false

  handleClick() {
    this.active = !this.active
  }
}
</script>

<style lang="sass" scoped>
.floating-add-tsundoku-button
  display: inline-flex
  flex-wrap: nowrap
  align-items: center
  position: fixed
  right: 0
  bottom: 94px + 18px

.button-open
  z-index: 100
  display: flex
  align-items: center
  background-color: var(--tsundoku-red)
  width: 94px
  height: 60px
  padding:
    left: 16px
  border:
    radius: 100vw 0 0 100vw
  color: var(--text-white)
  font:
    size: 24px
    weight: bold
  transition: border-radius .3s, width .3s, height .3s, margin .3s

  div
    display: flex
    align-items: center

  .is-active &
    width: 48px
    height: 48px
    border:
      radius: 100vw 100vw 100vw 100vw
    padding: 0
    margin:
      right: 12px
    justify-content: center

.icon-plus

  .is-active &
    transform: rotate(45deg)

%button
  display: flex
  align-items: center
  background-color: var(--tsundoku-red)
  justify-content: center
  width: 60px
  height: 60px
  border:
    radius: 100%

.button-search
  @extend %button
  margin:
    right: 12px

.button-scan
  @extend %button
  margin:
    right: 12px

.button-label
  padding:
    left: 4px
  white-space: nowrap

.transition-button
  &-enter, &-leave-to
    transform: translateX(120px)
    opacity: 0

  &-enter-active, &-leave-active
    transition: transform .5s, opacity .5s

.transition-label
  &-enter, &-leave-to
    transform: translateY(10px)
    opacity: 0

  &-enter-active
    transition: transform .5s, opacity .5s

  &-leave-active
    opacity: 0
    position: absolute

  &-move
    transition: transform .5s
</style>
