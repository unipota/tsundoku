<template lang="pug">
  .wrapper
    .tab-bar
      router-link.tab(
        v-for="tab in tabs"
        :key="tab.name"
        :to="{ name: tab.to, hash: $route.hash}"
        :class="{'selected': selectedPath === tab.name, [tab.name]: true}"
      )
        span.icon(:class="[tab.name]" key="icon")
          icon(
            :name="tab.name" 
            :color="selectedPath === tab.name ? undefined: `var(--${tab.inactiveColor})`"
            :height="20"
            :width="30")
        span.label(v-if="selectedPath === tab.name" key="label")
          | {{ $t(tab.name) }}
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

import Icon from '@/components/assets/Icon.vue'

@Component({
  components: { Icon }
})
export default class MobileTabBar extends Vue {
  get firstRouteName() {
    return this.$route.matched[0].path
  }
  get selectedPath() {
    return this.firstRouteName === ''
      ? 'tsundoku'
      : this.firstRouteName.slice(1)
  }

  private tabs = {
    tsundoku: {
      name: 'tsundoku',
      to: 'tsundoku',
      inactiveColor: 'tsundoku-red-bg'
    },
    kidoku: {
      name: 'kidoku',
      to: 'kidoku',
      inactiveColor: 'kidoku-blue-bg'
    },
    toukei: {
      name: 'toukei',
      to: 'toukei',
      inactiveColor: 'toukei-black-bg'
    }
  }
}
</script>

<style lang="sass" scoped>
.wrapper
  width: 100%
  background: rgba(255,255,255,0.6)
  backdrop-filter: blur(4px)
  border:
    radius: 32px 32px 0 0

  .tab-bar
    padding: 8px 36px
    display: flex
    width: auto
    max-width: 375px - (18px * 2)
    margin: auto
    box-sizing: content-box

    .tab
      display: flex
      position: relative
      // width: max-content
      // height: max-content
      border-radius: 44px
      cursor: pointer
      padding:
        top: 12px
        bottom: 12px
      transition: background-color .5s, width .5s
      background:
        color: transparent

      &:not(:last-child)
        margin-right: auto

      &.selected
        padding:
          left: 18px
          right: 18px

      &.tsundoku
        color: $tsundoku-red
        &.selected
          background-color: $tsundoku-red-bg
      &.kidoku
        color: $kidoku-blue
        &.selected
          background-color: $kidoku-blue-bg
      &.toukei
        color: $toukei-black
        &.selected
          background-color: $toukei-black-bg

      .icon
        display: flex
        margin:
          top: auto
          right: auto
          bottom: auto
          left: 0

      .label
        white-space: nowrap
        margin:
          top: auto
          right: 0
          bottom: auto
          left: 8px
        font:
          size: 16px
          weight: bold
</style>
