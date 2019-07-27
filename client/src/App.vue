<template lang="pug">
  #app
    component(:is="$store.state.viewType + '-template'")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'

const MobileTemplate = () => import('@/components/templates/MobileTemplate.vue')
const DesktopTemplate = () =>
  import('@/components/templates/DesktopTemplate.vue')

// TODO: move this to better position
const mobileThreshould = 750

@Component({
  components: {
    MobileTemplate,
    DesktopTemplate
  }
})
export default class App extends Vue {
  private windowWidth = 0
  public $store!: ExStore

  private handleResizeWindow() {
    const width = window.innerWidth
    if (this.$store.state.viewType !== 'mobile' && width < mobileThreshould) {
      this.$store.commit('setViewType', 'mobile')
    } else if (
      this.$store.state.viewType !== 'desktop' &&
      width >= mobileThreshould
    ) {
      this.$store.commit('setViewType', 'desktop')
    }
    this.windowWidth = width
  }

  private setLocale() {
    // ブラウザの言語設定を取得
    // this.$store.commit('setLocale', window.navigator.language.split('-')[0])
    this.$store.commit('setLocale', 'ja') // 今のところは日本語のみにしておく
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

#app
  height: 100vh
</style>
