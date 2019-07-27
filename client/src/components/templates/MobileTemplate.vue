<template lang="pug">
  div
    portal-target.modal-wrap(name="modalView")
    .top-bar-wrap
      mobile-top-bar(v-if="$store.getters.getShowMobileTopBar")
    .content-wrap(ref="scrollContainer")
      // keep-alive (see: DesktopTemplate)
      router-view
    .bottom-bar-wrap
      mobile-tab-bar(v-if="$store.getters.getShowMobileTabBar")
    transition(name="transition-floating-button")
      floating-add-tsundoku-button(v-show="selectedPath === 'tsundoku'")
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

<style lang="sass">
.top-bar-wrap
  position: fixed
  z-index: 1000
  top: 0
  width: 100%

.content-wrap
  position: relative
  overflow:
    x: hidden
    y: scroll
  width: 100vw
  height: 100vh
  padding:
    top: 80px
    bottom: 90px
  -webkit-overflow-scrolling: touch

.bottom-bar-wrap
  position: fixed
  z-index: 1000
  bottom: 0
  width: 100%

.modal-wrap
  position: fixed
  z-index: 3000

.transition-floating-button
  &-enter, &-leave-to
    transform: translateX(100%)
    opacity: 0

  &-enter-active, &-leave-active
    transition: transform .5s $easeInOutQuint, opacity .5s
</style>
