<template lang="pug">
  dialog-frame(@close="closeDialog")
    template(v-slot:body)
      .description
        | ログアウトしますか？
      .button-wrap
        .text-button.delete(@click="onLogout")
          | ログアウト
        .rounded-button.cancel(@click="onCancel")
          | キャンセル
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import DialogFrame from '@/components/atoms/DialogFrame.vue'

@Component({ components: { DialogFrame } })
export default class LogoutDialog extends Vue {
  $store!: ExStore

  closeDialog() {
    this.$emit('close')
  }

  onLogout() {
    this.$emit('logout')
  }

  onCancel() {
    this.$emit('cancel')
  }
}
</script>

<style lang="sass" scoped>
.description
  margin:
    bottom: 18px
  text-align: center
  font:
    weight: bold

  .sub-description
    color: var(--text-gray)

.button-wrap
  display: flex
  justify-content: space-between
  align-items: center
  padding: 0
  font:
    weight: bold

.delete
  color: var(--danger-red)
  padding: 0 12px
  user-select: none
  transition: color .3s

  &:hover
    color: var(--danger-red-hovered)

.text-button
  cursor: pointer

.rounded-button
  background: var(--border-gray)
  padding: 8px 12px
  border:
    radius: 100vw
  cursor: pointer
  user-select: none
  transition: background .3s

  &:hover
    background: var(--border-gray-hovered)
</style>
