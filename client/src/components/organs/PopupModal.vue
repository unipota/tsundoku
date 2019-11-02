<template lang="pug">
  transition(name="popup-modal-shown" mode="out-in")
    modal(:is="currentModalName" v-if="showModal")
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { VueConstructor } from 'vue'
import { ExStore } from 'vuex'

import LoginModal from '@/components/molecules/modal/LoginModal.vue'
import SettingModal from '@/components/molecules/modal/SettingModal.vue'
import PrivacyPolicyModal from '@/components/molecules/modal/PrivacyPolicyModal.vue'
import AboutModal from '../molecules/modal/AboutModal.vue'

interface PopupModalMap {
  [key: string]: VueConstructor<Vue>
}

@Component({ components: {} })
export default class PopupModal extends Vue {
  $store!: ExStore

  components: PopupModalMap = {
    login: LoginModal,
    setting: SettingModal,
    privacy: PrivacyPolicyModal,
    about: AboutModal
  }

  get showModal(): boolean {
    return this.$store.getters['modal/currentModalName'] !== ''
  }

  get currentModalName(): VueConstructor<Vue> {
    return this.components[this.$store.getters['modal/currentModalName']]
  }
}
</script>

<style lang="sass" scoped>
.popup-modal-shown
  &-enter, &-leave-to
    transform: translateY(24px)
    opacity: 0
  &-enter-active, &-leave-active
    transition: all .3s $easeInOutQuint
</style>
