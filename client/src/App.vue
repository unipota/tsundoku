<template lang="pug">
  #app
    component(:is="templateComponent + '-template'")
      template(v-slot:topBar)
        component(:is="templateComponent + '-top-bar'")
      template(v-slot:tabBar)
        component(:is="templateComponent + '-tab-bar'")
      template(v-slot:routerView)
        router-view
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import MobileTemplate from '@/components/templates/MobileTemplate.vue'
import DesktopTemplate from '@/components/templates/DesktopTemplate.vue'
import MobileTabBar from '@/components/atoms/MobileTabBar.vue'
import DesktopTabBar from '@/components/atoms/DesktopTabBar.vue'
import MobileTopBar from '@/components/atoms/MobileTopBar.vue'
import DesktopTopBar from '@/components/atoms/DesktopTopBar.vue'

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

  get templateComponent() {
    if (this.windowWidth < 750) {
      return 'mobile'
    } else {
      return 'desktop'
    }
  }

  mounted(): void {
    this.$nextTick(() => {
      this.handleResizeWindow()
    })
    window.addEventListener('resize', this.handleResizeWindow)
    window.addEventListener('orientationchange', this.handleResizeWindow)
  }
}
</script>

<style lang="sass">
@import '@/style/global.sass'
@import '@/style/reset.sass'
</style>
