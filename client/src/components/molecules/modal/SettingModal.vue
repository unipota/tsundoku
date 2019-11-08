<template lang="pug">
  popup-modal-frame(name="setting" title="設定")
    .setting-modal
      template(v-if="!userLogined")
        .user-icon-wrap
          icon(name="user" :width="64" :height="64")
        rounded-button.login-button(@click="pushLoginModal")
          | 新規登録/ログイン
        span.register-description
          | ユーザー登録をすると、他の端末やブラウザと同期することができます。
      template(v-else)
        .user-icon-container
          user-icon(:src="userIconUrl")
        .user-info-container
          .user-info-item     
            .user-info-item-name
              | ユーザー名
            .user-info-item-value
              | {{userName}}
          .user-info-item     
            .user-info-item-name
              | 登録日
            .user-info-item-value
              | {{createdAt}}
        rounded-button.logout-button(@click="logout")
          | ログアウト
      .link-wrap
        p.link(@click="openAboutModal")
          | ツンドクについて
        p.link(@click="openPrivacyPolicyModal")
          | {{ $t('privacy_policy') }}
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'
import dayjs from 'dayjs'

import Icon from '@/components/assets/Icon.vue'
import PopupModalFrame from '@/components/atoms/PopupModalFrame.vue'
import RoundedButton from '@/components/atoms/RoundedButton.vue'
import UserIcon from '@/components/atoms/UserIcon.vue'

@Component({ components: { Icon, PopupModalFrame, RoundedButton, UserIcon } })
export default class SettingModal extends Vue {
  $store!: ExStore

  pushLoginModal() {
    this.$store.commit('modal/push', { name: 'login' })
  }

  get userLogined() {
    return this.$store.state.userLogined
  }

  get userIconUrl() {
    return this.$store.state.userIconUrl
  }

  get userName() {
    return this.$store.state.userScreenName
  }

  get createdAt() {
    return dayjs(this.$store.state.userCreatedAt).format('YYYY/MM/DD')
  }

  logout() {
    this.$store.dispatch('userLogout')
  }

  openAboutModal() {
    this.$store.commit('modal/push', { name: 'about' })
  }

  openPrivacyPolicyModal() {
    this.$store.commit('modal/push', { name: 'privacy' })
  }
}
</script>

<style lang="sass" scoped>
.setting-modal
  display: flex
  flex-flow: column
  align-items: center
  min-height: 100%

.user-icon-wrap
  margin:
    top: 16px

.login-button
  margin:
    top: 32px

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

.user-icon-container
  margin:
    top: 16px
  width: 54px
  height: 54px

.user-info-container
  width: 100%
  padding:
    top: 32px
    left: 5%
    right: 5%

.user-info-item
  display: flex
  justify-content: space-between
  margin:
    top: 8px

.user-info-item-name
  color: var(--text-gray)
  font:
    weight: bold

.user-info-item-value
  color: var(--text-black)
  font:
    weight: bold

.logout-button
  margin:
    top: 48px

.link-wrap
  margin:
    top: auto
    bottom: 24px
  text-align: center

  .link
    color: var(--text-gray)
    font:
      weight: bold
    cursor: pointer
    margin: 8px auto
</style>
