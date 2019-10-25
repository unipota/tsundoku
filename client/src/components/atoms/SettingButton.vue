<template lang="pug">
  .setting-button(v-tooltip="'設定'" @click="openModal")
    icon(name="setting" :width="32" :height="32")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { ExStore } from 'vuex'

import Icon from '@/components/assets/Icon.vue'

@Component({ components: { Icon } })
export default class SettingButton extends Vue {
  $store!: ExStore

  openModal() {
    this.$store.commit('modal/push', { name: 'setting' })
  }
}
</script>

<style lang="sass">
.setting-button
  cursor: pointer
  opacity: 0.8
  transition: opacity .3s

  &:hover
    opacity: 1
</style>
