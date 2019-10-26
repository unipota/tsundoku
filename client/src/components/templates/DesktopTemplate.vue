<template lang="pug">
  .view-desktop(:class="{ 'modal-shown': modalShown }")
    .popup-modal-wrap
      popup-modal
    portal-target.modal-wrap(name="modalView")
    .modal-overlay(@click="closeModal")
    .nav-wrap
      desktop-nav(v-if="$store.state.showDesktopNav")
    .content-wrap
      // keep-alive だと複数存在する同名のポータルでハマるのでとりあえず無効化
      router-view
</template>

<script lang="ts">
import { Vue, Component, Watch } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import { ViewNames } from '../../router'
import DesktopNav from '@/components/molecules/DesktopNav.vue'
import PopupModal from '@/components/organs/PopupModal.vue'

@Component({
  components: { DesktopNav, PopupModal }
})
export default class DesktopTemplate extends Vue {
  public $store!: ExStore
  hideNavList: ViewNames[] = []

  @Watch('$route')
  private handleShowDesktopNav() {
    if (this.hideNavList.includes(this.$route.name as ViewNames)) {
      this.$store.commit('setShowDesktopNav', false)
    } else {
      this.$store.commit('setShowDesktopNav', true)
    }
  }

  get modalShown() {
    return (
      this.$route.meta.isModal || this.$store.getters['modal/currentModalName']
    )
  }

  mounted() {
    this.handleShowDesktopNav()
  }

  closeModal() {
    if (this.$route.meta.isModal) {
      //
    } else {
      this.$store.commit('modal/popAll')
    }
  }
}
</script>

<style lang="sass" scoped>
.view-desktop
  display: flex

.nav-wrap
  position: fixed
  z-index: 1000
  left: 0
  top: 0
  width: 320px
  height: 100%
  padding: 24px
  flex-shrink: 0

  // border:
  //   right: solid 1px var(--border-gray)

  transition: filter 0.5s $easeInOutQuint
  .modal-shown &
    // filter: blur(1px)

.content-wrap
  width: 100vw
  min-height: 100%
  margin:
    left: 320px
  overflow:
    x: hidden
    y: scroll

  .modal-shown &
    pointer-events: none // for iOS safari
    filter: blur(1px)

.popup-modal-wrap
  position: fixed
  z-index: 5000
  width: 100vw
  height: 100vh
  display: flex
  align-items: center
  justify-content: center
  pointer-events: none

.modal-wrap
  position: fixed
  z-index: 3000
  width: 100vw
  height: 100%
  pointer-events: none

.modal-overlay
  position: fixed
  top: 0
  left: 0
  z-index: 2999
  background-color: rgba(0, 0, 0, 0.5)
  width: 100vw
  height: 100vh
  // backdrop-filter: blur(4px)

  pointer-events: none

  transition: opacity .5s $easeInOutQuint
  opacity: 0

  .modal-shown &
    opacity: 1
    pointer-events: auto

.nav-wrap, .content-wrap
  transition: filter .5s $easeInOutQuint
  .modal-shown &
    // filter: blur(1px)
</style>
