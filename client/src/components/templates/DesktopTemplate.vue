<template lang="pug">
  .view-desktop(:class="{ 'modal-shown': $route.meta.isModal }")
    portal-target.modal-wrap(name="modalView")
    .nav-wrap
      desktop-nav(v-if="$store.state.showDesktopNav")
    .content-wrap
      // keep-alive だと複数存在する同名のポータルでハマるのでとりあえず無効化
      routerView
</template>

<script lang="ts">
import { Vue, Component, Watch } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import { ViewNames } from '../../router'
import DesktopNav from '@/components/molecules/DesktopNav.vue'

@Component({
  components: { DesktopNav }
})
export default class DesktopTemplate extends Vue {
  public $store!: ExStore
  hideNavList: ViewNames[] = ['login']

  @Watch('$route')
  private handleShowDesktopNav() {
    if (this.hideNavList.includes(this.$route.name as ViewNames)) {
      this.$store.commit('setShowDesktopNav', false)
    } else {
      this.$store.commit('setShowDesktopNav', true)
    }
  }

  mounted() {
    this.handleShowDesktopNav()
  }
}
</script>

<style lang="sass" scoped>
.view-desktop
  display: flex

.nav-wrap
  width: 320px
  padding: 24px
  flex-shrink: 0

  transition: filter 0.5s $easeInOutQuint
  .modal-shown &
    filter: blur(8px)

.content-wrap
  width: 100%
  padding:
    top: 24px
    bottom: 24px

.nav-wrap, .content-wrap
  transition: filter 0.5s $easeInOutQuint
  .modal-shown &
    filter: blur(8px)
</style>
