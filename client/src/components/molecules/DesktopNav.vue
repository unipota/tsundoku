<template lang="pug">
  section.main-nav
    .tab-header
      .brand
        icon(name="logo")
      router-link.setting(v-tooltip="'設定'" to="/user")
        icon(name="setting" :width="32" :height="32")
      router-link.user(v-if="!userLogined" v-tooltip="'新規登録/ログイン'" to="/login")
        icon(name="user" :width="32" :height="32")
      .user-icon-wrap(v-else)
        user-icon(:src="userIconUrl")
    .tab-wrap
      desktop-tab(:selected-tab="selectedPath" :price="price")
    .button-wrap
      router-link(
        :to="`${firstRouteName}/add-books-search`"
        v-tooltip.right="'キーワードで本を検索'")
        desktop-nav-button(:text="$t('add_by_searching')")
          icon(name="search" color="white" :height="24" :width="24")
    .button-wrap
      router-link(
        :to="`${firstRouteName}/add-books-scan`"
        v-tooltip.right="'カメラでバーコードをスキャン'")
        desktop-nav-button(:text="$t('add_by_scanning')")
          icon(name="scanner" color="white" :height="24" :width="24")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import DesktopTab from '@/components/atoms/DesktopTab.vue'
import DesktopNavButton from '@/components/atoms/DesktopNavButton.vue'
import Icon from '@/components/assets/Icon.vue'
import UserIcon from '@/components/atoms/UserIcon.vue'

@Component({
  components: {
    Icon,
    UserIcon,
    DesktopTab,
    DesktopNavButton
  }
})
export default class DesktopNav extends Vue {
  public $store!: ExStore

  get firstRouteName() {
    return this.$route.matched[0].path
  }
  get selectedPath() {
    return this.firstRouteName === ''
      ? 'tsundoku'
      : this.firstRouteName.slice(1)
  }
  get price(): number {
    return this.selectedPath === 'tsundoku'
      ? this.$store.getters.tsundokuPrice
      : this.selectedPath === 'kidoku'
      ? this.$store.getters.kidokuPrice
      : 0
  }
  get userLogined(): boolean {
    return this.$store.state.userLogined
  }
  get userIconUrl(): string | undefined {
    return this.$store.state.userIconUrl
  }
}
</script>

<style lang="sass" scoped>
.main-nav
  width: 100%
  height: 100%
.tab-header
  display: flex
  justify-content: flex-start
  align-items: center
  height: 40px
  width: 100%
  margin:
    bottom: 24px
  padding:
    left: 32px
.brand
  margin:
    right: auto
.setting, .user
  margin: 0 4px
  cursor: pointer
  opacity: 0.8
  transition: opacity .3s
  &:hover
    opacity: 1
.user-icon-wrap
  width: 28px
  height: 28px
  margin: 0 4px
.tab-wrap
  margin:
    top: 32px
    bottom: 48px
.button-wrap
  padding: 0 24px
  margin: 16px 0
</style>
