<template lang="pug">
  section.main-nav
    .tab-header
      .brand(@click="openAboutModal")
        icon(name="logo")
      setting-button.setting
      user-button.user(v-if="!userLogined")
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
    .button-wrap
      router-link(
        :to="`${firstRouteName}/add-books-edit`"
        v-tooltip.right="'手動で入力'")
        desktop-nav-button(:text="$t('add_by_edit')")
          icon(name="pen" color="white" :height="24" :width="24")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import DesktopTab from '@/components/atoms/DesktopTab.vue'
import DesktopNavButton from '@/components/atoms/DesktopNavButton.vue'
import Icon from '@/components/assets/Icon.vue'
import UserIcon from '@/components/atoms/UserIcon.vue'
import SettingButton from '@/components/atoms/SettingButton.vue'
import UserButton from '@/components/atoms/UserButton.vue'

@Component({
  components: {
    Icon,
    UserIcon,
    DesktopTab,
    DesktopNavButton,
    SettingButton,
    UserButton
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
  openAboutModal() {
    this.$store.commit('modal/push', { name: 'about' })
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
  cursor: pointer
.setting, .user
  margin: 0 4px
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
