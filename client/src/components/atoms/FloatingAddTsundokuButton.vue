<template lang="pug">
  .floating-add-tsundoku-button
    transition(name="transition-overlay")
      .overlay(v-if="active" @click="handleClose")
    transition-group.button-container(
      :class="{'is-active': active}"
      name="transition-button"
      tag="div")
      .button-search(@click="openSearchModal" key="search" v-if="active")
        icon(name="search" :width="24" :height="24")
        span.under-button-label 検索
      .button-scan(@click="openScanModal" key="scan" v-if="active")
        icon(name="scanner" :width="30" :height="30")
        span.under-button-label スキャン
      .button-scan(@click="openEditModal" key="edit" v-if="active")
        icon(name="pen" :width="30" :height="30" color="white")
        span.under-button-label 入力
      .button-open(key="open" @click="handleClick")
        transition-group.button-open-inner-wrapper(name="transition-label" tag="div")
          .icon-plus(key="icon")
            icon(name="plus" color="var(--text-white)")
          span.button-label(key="label" v-if="!active")
            | 積む
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import Icon from '@/components/assets/Icon.vue'

@Component({
  components: { Icon }
})
export default class FloatingAddTsundokuButton extends Vue {
  $store!: ExStore
  active = false

  handleClick() {
    this.active = !this.active
  }

  handleClose() {
    this.active = false
  }

  openSearchModal() {
    this.active = false
    this.$router.push(`${this.firstRouteName}/add-books-search`)
  }

  openScanModal() {
    this.active = false
    this.$router.push(`${this.firstRouteName}/add-books-scan`)
  }

  openEditModal() {
    this.active = false
    this.$router.push(`${this.firstRouteName}/add-books-edit`)
  }

  get firstRouteName() {
    return this.$route.matched[0].path
  }
}
</script>

<style lang="sass" scoped>
.floating-add-tsundoku-button
  position: fixed
  z-index: 2000
  top: 0
  right: 0
  width: 100%
  height: 100%
  pointer-events: none

  & > *
    pointer-events: auto

.overlay
  position: absolute
  z-index: 0
  top: 0
  left: 0
  width: 100%
  height: 100%
  background: rgba(0,0,0,0.5)

.button-container
  position: absolute
  z-index: 10
  right: 12px
  bottom: calc(63px + 12px)
  @supports (bottom: env(safe-area-inset-bottom))
    bottom: calc(env(safe-area-inset-bottom) + 63px + 12px)
  display: flex
  flex:
    flow: column
  flex-wrap: nowrap
  align-items: flex-end

.button-open
  cursor: pointer
  user-select: none
  z-index: 100
  overflow: hidden
  display: flex
  align-items: center
  background-color: var(--tsundoku-red)
  box-shadow: 0px 3px 5px -1px #e3402a6e
  width: 84px
  height: 48px
  padding:
    left: 14px
  border:
    radius: 100vw
  color: var(--text-white)
  font:
    size: 18px
    weight: bold
  transition: border-radius .3s, width .3s $easeInOutQuint, height .3s $easeInOutQuint, margin .3s, transform .3s

  .is-active &
    width: 42px
    height: 42px
    border:
      radius: 100vw 100vw 100vw 100vw
    padding: 0
    margin: 0 auto
    justify-content: center

  &-inner-wrapper
    display: inline-flex
    flex: 0
    align-items: center
    justify-content: center

.icon-plus
  height: 14px
  display: flex
  align-items: center
  justify-content: center

  & svg
    transform: rotate(0)
    will-change: transform
    transition: transform .5s

    .is-active &
      transform: rotate(135deg)

.under-button-label
  position: absolute
  top: 50%
  right: 100%
  text-align: center
  white-space: nowrap
  font:
    weight: bold
  color: var(--bg-white)
  opacity: 0
  animation: button-label .5s $easeInOutQuint .5s forwards
  transform: translateY(-50%)

@keyframes button-label
  0%
    opacity: 0
  100%
    opacity: 1
    transform: translate(-12px, -50%)

%button
  display: flex
  align-items: center
  position: relative
  background-color: var(--tsundoku-red)
  box-shadow: 0px 3px 5px -1px #e3402a6e
  justify-content: center
  width: 48px
  height: 48px
  border:
    radius: 100%

.button-search
  @extend %button
  margin:
    bottom: 12px

.button-scan
  @extend %button
  margin:
    bottom: 12px

.button-label
  padding:
    left: 4px
  white-space: nowrap

.transition-button
  &-enter, &-leave-to
    transform: translateY(60px)
    opacity: 0

  &-enter-active
    transition: transform .5s $easeOutBack, opacity .5s
    pointer-events: none

  &-leave-active
    transition: transform .5s, opacity .5s

.transition-label
  &-enter, &-leave-to
    transform: translateX(10px)
    opacity: 0

  &-enter-active
    transition: transform .5s $easeInOutQuint, opacity .5s

  &-leave-active
    opacity: 0
    position: absolute

  &-move
    transition: transform .5s

.transition-overlay
  &-enter, &-leave-to
    opacity: 0

  &-enter-active, &-leave-active
    transition: opacity .5s
</style>
