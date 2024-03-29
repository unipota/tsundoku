<template lang="pug">
  nav.tab-bar
    router-link.tab(
      v-for="tab in Object.keys(tabs)"
      :key="tab"
      :class="{ [tab]: true, 'active': isActive(tab) }"
      :to="{ name: tabs[tab].to, hash: $route.hash }"
      v-tooltip.right="tabs[tab].tooltip"
    )
      span.icon(:class="tab")
        icon(
          :name="tab"
          :color="isActive(tab) ? undefined : tabs[tab].inactiveColor"
          :height="48"
        )
      transition-group.label.label-with-price(name="transition-label" mode="out-in" tag="div")
        span.item-label(key="item-label" :class="{'with-price': tab !== 'toukei' && isActive(tab)}")
          | {{ $t(tab) }}
        span.price-label(v-if="tab !== 'toukei' && isActive(tab)" key="item-price")
          tweened-number(:num="price" formatLocal)
</template>

<script lang="ts">
import { Vue, Prop, Component } from 'vue-property-decorator'

import Icon from '@/components/assets/Icon.vue'
import TweenedNumber from '@/components/atoms/TweenedNumber.vue'

@Component({
  components: {
    Icon,
    TweenedNumber
  }
})
export default class DesktopTab extends Vue {
  @Prop({ type: String, default: '' }) selectedTab!: string
  @Prop({ type: Number, default: 0 }) price!: number

  private tabs = {
    tsundoku: {
      to: 'tsundoku',
      inactiveColor: 'var(--tsundoku-red-bg)',
      tooltip: '積んでいる本のリストとツンドク総額'
    },
    kidoku: {
      to: 'kidoku',
      inactiveColor: 'var(--kidoku-blue-bg)',
      tooltip: '読み終わった本のリストとキドク総額'
    },
    toukei: {
      to: 'toukei',
      inactiveColor: 'var(--toukei-black-bg)',
      tooltip: 'ツンドク状況の推移'
    }
  }

  private isActive(tab: string) {
    return this.selectedTab === tab
  }
}
</script>

<style lang="sass" scoped>
.tab-bar
  display: flex
  flex-direction: column
  width: 100%
.tab
  display: flex
  align-items: center
  height: 72px
  width: 100%
  padding: 0 32px
  border-radius: 100vw
  transition: background-color .3s
  &:not(:last-child)
    margin-bottom: 24px
  &:not(.active):hover
    background-color: var(--bg-gray)

.icon
  display: flex
  align-items: center
  justify-content: center
  width: 48px
  margin-right: 24px

.tsundoku
  color: var(--tsundoku-red-bg)
  &.active
    color: var(--tsundoku-red)
    background-color: var(--tsundoku-red-bg)
.kidoku
  color: var(--kidoku-blue-bg)
  &.active
    color: var(--kidoku-blue)
    background-color: var(--kidoku-blue-bg)
.toukei
  color: var(--toukei-black-bg)
  &.active
    color: var(--toukei-black)
    background-color: var(--toukei-black-bg)

.label
  font-weight: bold

.label-with-price
  display: flex
  flex-direction: column
  align-items: flex-start
  .item-label
    transition: transform .3s $easeInOutQuint, opacity .3s, font-size .2s
    &.with-price
      font-size: 0.8rem
  .price-label
    font-size: 1.2rem
  .price-label::before
    content: '¥'
    display: inline-block
    margin-right: 0.25rem

.transition-label
  &-enter
    transform: translateY(24px)
    opacity: 0

  &-leave-to
    transform: translateY(24px)
    opacity: 0

  &-enter-active, &-leave-active
    transition: transform .3s $easeInOutQuint, opacity .3s

  &-leave-active
    position: absolute

  &-move
    transition: transform .3s
</style>
