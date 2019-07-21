<template lang="pug">
  section.main-nav
    .brand
      icon-logo
    .tab-wrap
      desktop-tab(:selected-tab="$route.name")
    router-link.button-wrap(:to="`/${$route.matched[0].name}/add-books-search`")
      desktop-nav-button(:text="$t('add_by_searching')")
        icon-search(color="white" height="24")
    router-link.button-wrap(:to="`/${$route.matched[0].name}/add-books-scan`")
      desktop-nav-button(:text="$t('add_by_scanning')")
        icon-scanner(color="white" height="24")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import DesktopTab from '@/components/atoms/DesktopTab.vue'
import DesktopNavButton from '@/components/atoms/DesktopNavButton.vue'
import IconLogo from '@/components/assets/IconLogo.vue'
import IconSearch from '@/components/assets/IconSearch.vue'
import IconScanner from '@/components/assets/IconScanner.vue'

@Component({
  components: {
    DesktopTab,
    DesktopNavButton,
    IconLogo,
    IconSearch,
    IconScanner
  }
})
export default class DesktopNav extends Vue {}
</script>

<style lang="sass" scoped>
.main-nav
  width: 320px
  height: 100%
  padding: 24px
.brand
  display: flex
  justify-content: flex-start
  align-items: center
  height: 40px
  width: 100%
  margin-bottom: 24px
  padding: 0 32px
.tab-wrap
  margin:
    top: 32px
    bottom: 48px
.button-wrap
  padding: 0 24px
  margin: 16px 0
</style>
