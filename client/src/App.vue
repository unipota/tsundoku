<template lang="pug">
  #app
    component(:is="templateComponent + '-template'")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
const MobileTemplate = () => import('@/components/templates/MobileTemplate.vue')
const DesktopTemplate = () =>
  import('@/components/templates/DesktopTemplate.vue')

@Component({
  components: {
    MobileTemplate,
    DesktopTemplate
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
