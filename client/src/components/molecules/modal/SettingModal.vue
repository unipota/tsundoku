<template lang="pug">
  popup-modal-frame(name="setting" title="設定")
    .setting-modal
      .user-icon
        icon(name="user" :width="64" :height="64")
      .login-button(@click="pushLoginModal")
        | 新規登録/ログイン
      span.register-description
        | ユーザー登録をすると、他の端末やブラウザと同期することができます。
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import Icon from '@/components/assets/Icon.vue'
import PopupModalFrame from '@/components/atoms/PopupModalFrame.vue'

@Component({ components: { Icon, PopupModalFrame } })
export default class SettingModal extends Vue {
  $store!: ExStore

  pushLoginModal() {
    this.$store.commit('modal/push', { name: 'login' })
  }
}
</script>

<style lang="sass" scoped>
.setting-modal
  display: flex
  flex-flow: column
  align-items: center

.user-icon
  margin:
    top: 16px

.login-button
  margin:
    top: 32px
  background: var(--border-gray)
  color: var(--text-black)
  border-radius: 999vw
  min-width: 200px
  text-align: center
  font:
    weight: bold
  padding: 8px
  cursor: pointer

.register-description
  margin:
    top: 16px
    left: 8px
    right: 8px
  background: var(--bg-gray)
  border:
    radius: 8px
  padding: 16px
  color: var(--text-gray)
  font:
    weight: bold
    size: 14px
</style>