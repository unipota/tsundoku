<template lang="pug">
  .view-desktop
    portal-target.modal-wrap(name="modalView")
    .nav-wrap
      desktop-nav
    .view-wrap
      // keep-alive だと複数存在する同名のポータルでハマるのでとりあえず無効化
      routerView
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import DesktopNav from '@/components/molecules/DesktopNav.vue'

@Component({
  components: { DesktopNav }
})
export default class DesktopTemplate extends Vue {}
</script>

<style lang="sass" scoped>
.view-desktop
  display: flex
.nav-wrap
  width: 320px
  padding: 24px
  flex-shrink: 0
.view-wrap
  width: 100%
  padding:
    top: 24px
    bottom: 24px
</style>
