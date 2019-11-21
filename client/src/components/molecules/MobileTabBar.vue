<template lang="pug">
  .wrapper
    .tab-bar
      .tab-body(v-for="(tab, index) in tabs"
          :ref="tab.name"
          :key="tab.name"
          :style="{width: tab.width}")
        router-link(:to="{ name: tab.to, hash: $route.hash}")
          transition-group.tab(name="transition-tab" tag="div"
            :class="{'selected': selectedPath === tab.name, [tab.name]: true}")
            .icon(:class="[tab.name]" key="icon")
              icon(
                :name="tab.name" 
                :color="selectedPath === tab.name ? undefined: `var(--${tab.inactiveColor})`"
                :height="20"
                :width="30"
                key="icon")
            .label(v-if="selectedPath === tab.name" key="label")
              | {{ $t(tab.name) }}
</template>

<script lang="ts">
import { Vue, Component, Watch } from 'vue-property-decorator'

import Icon from '@/components/assets/Icon.vue'

interface TabProps {
  name: string
  to: string
  inactiveColor: string
  width: string
}

interface Tabs {
  tsundoku: TabProps
  kidoku: TabProps
  toukei: TabProps
}

type TabNames = keyof Tabs

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
  get tabStyle() {
    return {}
  }

  private tabs: Tabs = {
    tsundoku: {
      name: 'tsundoku',
      to: 'tsundoku',
      inactiveColor: 'tsundoku-red-bg',
      width: ''
    },
    kidoku: {
      name: 'kidoku',
      to: 'kidoku',
      inactiveColor: 'kidoku-blue-bg',
      width: ''
    },
    toukei: {
      name: 'toukei',
      to: 'toukei',
      inactiveColor: 'toukei-black-bg',
      width: ''
    }
  }

  @Watch('selectedPath')
  onChangeTab(toTab: TabNames, fromTab: TabNames) {
    this.tabs[toTab].width = ''
    this.tabs[fromTab].width = ''

    // DOMの更新を待つ必要がある
    this.$nextTick(() => {
      const elements = this.$refs[toTab] as Element[]
      const el = elements[0]
      this.tabs[toTab].width = el.clientWidth + 'px'
    })
  }
}
</script>

<style lang="sass" scoped>
.wrapper
  width: 100%
  background: rgba(255,255,255,0.8)
  backdrop-filter: blur(4px)
  border:
    radius: 32px 32px 0 0
  @supports (padding-bottom: env(safe-area-inset-bottom))
    padding-bottom: calc(env(safe-area-inset-bottom))

  .tab-bar
    padding: 8px 36px
    display: flex
    justify-content: space-around
    width: auto
    max-width: 375px - (18px * 2)
    margin: auto
    box-sizing: content-box

    .tab-body
      transition: width .2s linear

    .tab
      display: flex
      position: relative
      border-radius: 44px
      cursor: pointer
      padding:
        top: 12px
        left: 12px
        right: 12px
        bottom: 12px
      transition: background-color .5s, width .5s
      background:
        color: transparent

      &.selected
        padding:
          left: 18px
          right: 18px

      &.tsundoku
        color: var(--tsundoku-red)
        &.selected
          background-color: var(--tsundoku-red-bg)
      &.kidoku
        color: var(--kidoku-blue)
        &.selected
          background-color: var(--kidoku-blue-bg)
      &.toukei
        color: var(--toukei-black)
        &.selected
          background-color: var(--toukei-black-bg)

      .icon
        width: 30px
        height: 20px
        display: flex
        justify-content: center
        align-items: center

      .label
        height: 20px
        white-space: nowrap
        margin:
          top: auto
          right: 0
          bottom: auto
          left: 8px
        font:
          size: 16px
          weight: bold

.transition-tab
  &-enter, &-leave-to
    opacity: 0

  &-enter-active, &-leave-active
    transition: opacity .5s
    transform: translateX(20px)

  &-leave-active
    display: none
    opacity: 0
    position: absolute

  &-move
    transition: transform .8s $easeInOutQuint
</style>
