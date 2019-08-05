<template lang="pug">
  .view-mobile(:class="{ 'modal-shown': $route.meta.isModal }")
    portal-target.popover-wrap(name="popoverView")
    portal-target.modal-wrap(name="modalView")
    .modal-overlay
    .top-bar-wrap
      mobile-top-bar(v-if="$store.state.showMobileTopBar")
    .content-wrap(ref="scrollContainer")
      // keep-alive (see: DesktopTemplate)
      router-view
    .bottom-bar-wrap
      mobile-tab-bar(v-if="$store.state.showMobileTabBar")
    transition(name="transition-floating-button")
      floating-add-tsundoku-button.tsundoku-button(v-show="selectedPath === 'tsundoku'")
</template>

<script lang="ts">
import { Vue, Component, Watch } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import { ViewNames } from '../../router'
import MobileTabBar from '@/components/molecules/MobileTabBar.vue'
import MobileTopBar from '@/components/molecules/MobileTopBar.vue'
import FloatingAddTsundokuButton from '@/components/atoms/FloatingAddTsundokuButton.vue'

@Component({
  components: {
    MobileTabBar,
    MobileTopBar,
    FloatingAddTsundokuButton
  }
})
export default class MobileTemplate extends Vue {
  public $store!: ExStore
  get firstRouteName() {
    return this.$route.matched[0].path
  }
  get selectedPath() {
    return this.firstRouteName === ''
      ? 'tsundoku'
      : this.firstRouteName.slice(1)
  }

  hideTopBarList: ViewNames[] = ['login', 'user']
  hideTabBarList: ViewNames[] = ['login', 'user']

  @Watch('$route')
  private handleShowMobileBars() {
    this.$store.commit(
      'setShowMobileTopBar',
      !this.hideTopBarList.includes(this.$route.name as ViewNames)
    )
    this.$store.commit(
      'setShowMobileTabBar',
      !this.hideTabBarList.includes(this.$route.name as ViewNames)
    )
  }

  mounted() {
    this.handleShowMobileBars()
  }
}
</script>

<style lang="sass" scoped>
.view-mobile
  width: 100vw

.top-bar-wrap
  position: fixed
  z-index: 1000
  top: 0
  width: 100%

.content-wrap
  position: relative
  padding:
    top: 64px
    bottom: 90px

.bottom-bar-wrap
  position: fixed
  z-index: 1000
  bottom: 0
  bottom: env(safe-area-inset-bottom)
  width: 100%

.top-bar-wrap, .content-wrap, .tsundoku-button
  transition: filter 0.5s $easeInOutQuint
  .modal-shown &
    filter: blur(4px)

.popover-wrap
  position: fixed
  z-index: 4000

.modal-wrap
  position: fixed
  z-index: 3000
  width: 100%
  height: 100%
  pointer-events: none

.modal-overlay
  position: fixed
  top: 0
  bottom: 0
  z-index: 2999
  background-color: rgba(0, 0, 0, 0.5)

  width: 100vw
  height: 100vh

  pointer-events: none

  transition: opacity .5s $easeInOutQuint
  opacity: 0

  .modal-shown &
    opacity: 1
    pointer-events: auto

.transition-floating-button
  &-enter, &-leave-to
    transform: translateX(100%)
    opacity: 0

  &-enter-active, &-leave-active
    transition: transform .5s $easeInOutQuint, opacity .5s
</style>
