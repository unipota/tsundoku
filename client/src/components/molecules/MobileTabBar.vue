<template lang="pug">
  div.wrapper
    div.tab-bar
      router-link.tab.tsundoku(
        :class="{'selected': selectedPath === 'tsundoku'}"
        :to="{ name: tabs.tsundoku.to, hash: $route.hash}"
      )
        span.icon.tsundoku
          icon(name="tsundoku" :color="selectedPath === 'tsundoku' ? undefined: 'var(--tsundoku-red-bg)'" :height="20")
        span.label(v-if="selectedPath === 'tsundoku'")
          | {{ $t('tsundoku') }}

      router-link.tab.kidoku(
        :class="{'selected': selectedPath === 'kidoku'}"
        :to="{ name: tabs.kidoku.to, hash: $route.hash}"
      )
        span.icon.tsundoku
          icon(name="kidoku" :color="selectedPath==='kidoku' ? undefined: 'var(--kidoku-blue-bg)'" :height="20").icon.kidoku
        span.label(v-if="selectedPath === 'kidoku'")
          | {{ $t('kidoku') }}

      router-link.tab.toukei(
        :class="{'selected': selectedPath === 'toukei'}"
        :to="{ name: tabs.toukei.to }"
      )
        span.icon.toukei
          icon(name="toukei" :color="selectedPath==='toukei' ? undefined: 'var(--toukei-black-bg)'" :height="20").icon.toukei
        span.label(v-if="selectedPath === 'toukei'")
          | {{ $t('toukei') }}
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
      to: 'tsundoku'
    },
    kidoku: {
      name: 'kidoku',
      to: 'kidoku'
    },
    toukei: {
      name: 'toukei',
      to: 'toukei'
    }
  }
}
</script>

<style lang="sass" scoped>
.wrapper
  width: 100%
  background: rgba(255,255,255,0.6)
  backdrop-filter: blur(2px)
  border:
    radius: 32px 32px 0 0

  .tab-bar
    padding: 20px 30px
    display: flex
    width: auto
    max-width: 375px - (18px * 2)
    margin: auto
    box-sizing: content-box

    .tab
      display: flex
      width: max-content
      height: max-content
      border-radius: 44px
      cursor: pointer
      padding:
        top: 12px
        bottom: 12px

      &:not(:last-child)
        margin-right: auto

      &.selected
        padding:
          left: 19px
          right: 19px

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
        margin:
          top: auto
          right: 0
          bottom: auto
          left: 8px
        font:
          size: 16px
          weight: bold
</style>