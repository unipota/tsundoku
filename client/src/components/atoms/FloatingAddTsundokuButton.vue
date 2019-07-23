<template lang="pug">
  transition-group.floating-add-tsundoku-button(:class="{'is-active': active}" name="transition-button" tag="div")
    router-link(:to="`${firstRouteName}/add-books-search`").button-search(key="search" v-show="active")
      icon-search(width="30" height="30")
    router-link(:to="`${firstRouteName}/add-books-scan`").button-scan(key="scan" v-show="active")
      icon-scanner(width="38" height="38")
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
  get firstRouteName() {
    return this.$route.matched[0].path
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
  bottom: calc(94px + 1vh)

.button-open
  z-index: 100
  overflow: hidden
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
  transition: border-radius .3s, width .3s $easeOutBack, height .3s $easeOutBack, margin .3s

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
  transform: rotate(0)

  .is-active &
    transform: rotate(135deg)

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
    transform: translateX(160px) rotate(30deg)
    // opacity: 0

  &-enter-active
    transition: transform .5s $easeOutBack, opacity .5s
    pointer-events: none

  &-leave-active
    transition: transform .5s, opacity .5s

.transition-label
  &-enter, &-leave-to
    transform: translateY(10px)
    opacity: 0

  &-enter-active
    transition: transform .5s $easeOutBack, opacity .5s

  &-leave-active
    opacity: 0
    position: absolute

  &-move
    transition: transform .5s
</style>
