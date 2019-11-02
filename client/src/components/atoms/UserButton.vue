<template lang="pug">
  .user-button(v-tooltip="'新規登録/ログイン'" @click="openModal")
    icon(name="user" :width="32" :height="32")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import Icon from '@/components/assets/Icon.vue'

@Component({ components: { Icon } })
export default class UserButton extends Vue {
  $store!: ExStore

  openModal() {
    this.$store.commit('modal/push', { name: 'login' })
  }
}
</script>

<style lang="sass">
.user-button
  display: flex
  align-items: center
  cursor: pointer
  opacity: 0.8
  transition: opacity .3s

  &:hover
    opacity: 1
</style>
