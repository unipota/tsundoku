<template lang="pug">
  #app
    component(:is="templateComponent + '-template'")
      template(v-if="templateComponent='mobile'" #topBar)
        component(:is="templateComponent + '-top-bar'")
      template(#tabBar)
        component(:is="templateComponent + '-tab-bar'")
      template(#routerView)
        router-view
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import MobileTemplate from '@/components/templates/MobileTemplate.vue'
import MobileTabBar from '@/components/atoms/MobileTabBar.vue'
import MobileTopBar from '@/components/molecules/MobileTopBar.vue'
import DesktopTemplate from '@/components/templates/DesktopTemplate.vue'
import DesktopTopBar from '@/components/molecules/DesktopTopBar.vue'
import DesktopTabBar from '@/components/molecules/DesktopTabBar.vue'

@Component({
  components: {
    MobileTemplate,
    DesktopTemplate,
    MobileTabBar,
    DesktopTabBar,
    MobileTopBar,
    DesktopTopBar
  }
})
export default class App extends Vue {
  private windowWidth = 0

  private handleResizeWindow() {
    this.windowWidth = window.innerWidth
  }

  private setLocale() {
    // ブラウザの言語設定を取得
    this.$store.commit('setLocale', window.navigator.language.split('-')[0])
  }

  get templateComponent() {
    if (this.windowWidth < 750) {
      return 'mobile'
    } else {
      return 'desktop'
    }
  }

  mounted() {
    this.$nextTick(() => {
      this.handleResizeWindow()
    })
    this.setLocale()
    window.addEventListener('resize', this.handleResizeWindow)
    window.addEventListener('orientationchange', this.handleResizeWindow)
  }
}
</script>

<style lang="sass">
@import '@/style/global.sass'
@import '@/style/reset.sass'
</style>
