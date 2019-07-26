<template lang="pug">
  div
    portal-target.modal-wrap(name="modalView")
    .top-bar-wrap
      mobile-top-bar
    .content-wrap(ref="scrollContainer")
      // keep-alive (see: DesktopTemplate)
      router-view
    .bottom-bar-wrap
      mobile-tab-bar
    floating-add-tsundoku-button(v-if="selectedPath === 'tsundoku'")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'

import MobileTabBar from '@/components/atoms/MobileTabBar.vue'
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
  get firstRouteName() {
    return this.$route.matched[0].path
  }
  get selectedPath() {
    return this.firstRouteName === ''
      ? 'tsundoku'
      : this.firstRouteName.slice(1)
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
  z-index: 2000
</style>
