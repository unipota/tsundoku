<template lang="pug">
  popup-modal-frame(name="setting" title="設定")
    .setting-modal
      .user-icon
        icon(name="user" :width="64" :height="64")
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
  justify-content: center
</style>
