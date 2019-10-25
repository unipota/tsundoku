<template lang="pug">
  #app
    component(:is="$store.state.viewType + '-template'")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'
import { createDecorator } from 'vue-class-component'
import { throttle } from 'lodash'

const Meta = createDecorator((options, key) => {
  if (!options.methods) {
    return
  }
  options['metaInfo'] = options.methods[key]
})

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

  @Meta
  metaInfo() {
    return {
      title: 'ツンドク',
      htmlAttrs: {},
      bodyAttrs: {
        class: [this.modalShown ? 'scroll-fix' : '']
      }
    }
  }

  get modalShown() {
    return (
      this.$route.meta.isModal || this.$store.getters['modal/currentModalName']
    )
  }

  toggleThemeType(mql: MediaQueryListEvent | MediaQueryList) {
    if (mql.matches) {
      this.$store.commit('setCurrentTheme', 'dark')
    } else {
      this.$store.commit('setCurrentTheme', 'light')
    }
  }

  async mounted() {
    this.handleResizeWindow()
    this.setLocale()
    window.addEventListener('resize', throttle(this.handleResizeWindow, 100))
    window.addEventListener('orientationchange', this.handleResizeWindow)
    const isDark: MediaQueryList = window.matchMedia(
      '(prefers-color-scheme: dark)'
    )
    isDark.addListener(this.toggleThemeType)
    this.toggleThemeType(isDark)
    await this.$store.dispatch('whoAmI')
    await this.$store.dispatch('getMyBooks')
  }
}
</script>

<style lang="sass">
@import '@/style/global.sass'
@import '@/style/reset.sass'
@import '@/style/tooltip.sass'
</style>
