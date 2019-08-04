<template lang="pug">
  transition-group.floating-add-tsundoku-button(
    :class="{'is-active': active}"
    name="transition-button"
    tag="div"
    v-click-outside="handleClickOutside")
    .button-search(@click="openSearchModal" key="search" v-show="active")
      icon(name="search" :width="24" :height="24")
      span.under-button-label 検索
    .button-scan(@click="openScanModal" key="scan" v-show="active")
      icon(name="scanner" :width="30" :height="30")
      span.under-button-label スキャン
    .button-scan(@click="openEditModal" key="edit" v-show="active")
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
import vClickOutside from 'v-click-outside'

import Icon from '@/components/assets/Icon.vue'

@Component({
  components: { Icon },
  directives: {
    clickOutside: vClickOutside.directive
  }
})
export default class FloatingAddTsundokuButton extends Vue {
  active = false

  handleClick() {
    this.active = !this.active
  }

  handleClickOutside() {
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
  right: 12px
  bottom: calc(63px + 12px)
  display: inline-flex
  flex-wrap: nowrap
  align-items: center

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
  transition: border-radius .3s, width .3s $easeInOutQuint, height .3s $easeInOutQuint, margin .3s

  .is-active &
    width: 42px
    height: 42px
    border:
      radius: 100vw 100vw 100vw 100vw
    padding: 0
    margin:
      right: 12px
    justify-content: center

  &-inner-wrapper
    display: inline-flex
    flex: 0
    align-items: center
    justify-content: center

.icon-plus
  transform: rotate(0)
  height: 36px
  display: inline-flex
  align-items: center
  justify-content: center

  .is-active &
    transform: rotate(135deg)

.under-button-label
  position: absolute
  top: calc(100% + 4px)
  text-align: center
  white-space: nowrap
  font:
    weight: bold
  color: var(--tsundoku-red)
  opacity: 0
  transform: translateY(-100%)

  .is-active &
    opacity: 1
    transform: translateY(0)
    transition: all .5s $easeInOutQuint .3s

%button
  display: flex
  align-items: center
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
    transform: translateX(60px) rotate(30deg)
    opacity: 0

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
    transition: transform .5s $easeInOutQuint, opacity .5s

  &-leave-active
    opacity: 0
    position: absolute

  &-move
    transition: transform .5s
</style>
