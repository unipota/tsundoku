<template lang="pug">
  .user-button-wrap
    .user-button(v-if="!userLogined" v-tooltip="'新規登録/ログイン'" @click="openLoginModal")
      icon(name="user" :width="32" :height="32")
    .user-button(v-else v-tooltip="'ユーザー設定'" @click="openSettingModal")
      user-icon(:src="userIconUrl")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import Icon from '@/components/assets/Icon.vue'
import UserIcon from '@/components/atoms/UserIcon.vue'

@Component({ components: { Icon, UserIcon } })
export default class UserButton extends Vue {
  $store!: ExStore

  openLoginModal() {
    this.$store.commit('modal/push', { name: 'login' })
  }

  openSettingModal() {
    this.$store.commit('modal/push', { name: 'setting' })
  }

  get userLogined(): boolean {
    return this.$store.state.userLogined
  }

  get userIconUrl(): string | undefined {
    return this.$store.state.userIconUrl
  }
}
</script>

<style lang="sass">
.user-button
  position: relative
  width: 32px
  height: 32px
  display: flex
  align-items: center
  cursor: pointer
  transition: opacity .3s

  &:hover
    opacity: 0.8
</style>
