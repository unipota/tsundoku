<template lang="pug">
  .mobile-top-bar
    .price-display
      portal-target(name="priceDisplay")
    router-link.setting(v-tooltip="'設定'" to="/user")
      icon(name="setting" :width="32" :height="32")
    router-link.user(v-if="!userLogined" v-tooltip="'新規登録/ログイン'" to="/login")
      icon(name="user" :width="32" :height="32")
    .user-icon-wrap(v-else)
      user-icon(src="userIconUrl")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import UserIcon from '@/components/atoms/UserIcon.vue'
import Icon from '@/components/assets/Icon.vue'

@Component({
  components: { UserIcon, Icon }
})
export default class MobileTopBar extends Vue {
  public $store!: ExStore

  get userLogined(): boolean {
    return this.$store.state.userLogined
  }
  get userIconUrl(): string | undefined {
    return this.$store.state.userIconUrl
  }
}
</script>

<style lang="sass">
.mobile-top-bar
  display: flex
  justify-content: flex-end
  align-items: center
  padding:
    top: 8px
    left: 16px
    right: 16px
    bottom: 8px
  background: rgba(255,255,255,1)
  box-shadow: 0 0px 6px -4px var(--text-gray)

.price-display
  margin:
    right: auto

.setting, .user
  display: flex
  align-items: center
  margin: 0 2px

.user-icon-wrap
  width: 28px
  height: 28px
  margin: 0 2px
</style>
